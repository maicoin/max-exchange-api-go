// Copyright 2018 MaiCoin Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package max

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/maicoin/max-exchange-api-go/models"

	event "github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
)

// wsClient allow to connect and receive stream data
// from max.com ws service.
type wsClient struct {
	conn   *websocket.Conn
	connMu sync.RWMutex
	stopCh chan struct{}
	evBus  event.Bus

	accessKey string
	secretKey string
	URL       string
	logger    *log.Logger
}

// NewWSClient returns a websocket client.
func NewWSClient(opts ...WebsocketClientOption) (*wsClient, error) {
	client := &wsClient{
		stopCh: make(chan struct{}),
		evBus:  event.New(),
		URL:    "wss://max-ws.maicoin.com",
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	for _, opt := range opts {
		opt(client)
	}

	d := &websocket.Dialer{
		Subprotocols:    []string{"p1", "p2"},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Proxy:           http.ProxyFromEnvironment,
	}

	var err error
	client.conn, _, err = d.Dial(client.URL, nil)
	if err != nil {
		return nil, err
	}

	go client.handleMsg()

	return client, nil
}

// Close web socket connection
func (w *wsClient) Close() {
	w.stopCh <- struct{}{}
	w.conn.Close()
}

// SubscribeTicker subscribes the realtime price information
func (w *wsClient) SubscribeTicker(market string, ch chan *models.TickerEvent) (TickerSubscription, error) {
	handler := func(ev *models.TickerEvent) {
		ch <- ev
	}

	unsubscriber, err := w.subscribeChannel("ticker", map[string]interface{}{
		"market": market,
	}, handler)
	if err != nil {
		return nil, err
	}

	return &tickerSubscription{
		ch:      ch,
		onEvent: handler,
		unsubscribe: func() {
			unsubscriber()
			close(ch)
		},
	}, nil
}

// SubscribeOrderBook subscribes the realtime changes on order books
func (w *wsClient) SubscribeOrderBook(market string, ch chan *models.OrderBookEvent) (OrderBookSubscription, error) {
	handler := func(ev *models.OrderBookEvent) {
		ch <- ev
	}

	unsubscriber, err := w.subscribeChannel("orderbook", map[string]interface{}{
		"market": market,
	}, handler)
	if err != nil {
		return nil, err
	}

	return &orderBookSubscription{
		ch:      ch,
		onEvent: handler,
		unsubscribe: func() {
			unsubscriber()
			close(ch)
		},
	}, nil
}

// SubscribeTrade subscribes the realtime trades information
func (w *wsClient) SubscribeTrade(market string, ch chan *models.TradeEvent) (TradeSubscription, error) {
	handler := func(ev *models.TradeEvent) {
		ch <- ev
	}

	unsubscriber, err := w.subscribeChannel("trade", map[string]interface{}{
		"market": market,
	}, handler)
	if err != nil {
		return nil, err
	}

	return &tradeSubscription{
		ch:      ch,
		onEvent: handler,
		unsubscribe: func() {
			unsubscriber()
			close(ch)
		},
	}, nil
}

// SubscribeAccount subscribes the accounts changes for an user
//
// Note:
//     Use WSAuthToken() to pass your auth tokens.
func (w *wsClient) SubscribeAccount(ch chan models.AccountEvent) (AccountSubscription, error) {
	handler := func(ev models.AccountEvent) {
		ch <- ev
	}

	topic := "account"
	if err := w.evBus.SubscribeAsync(topic, handler, true); err != nil {
		return nil, err
	}

	unsubscriber := func() {
		w.evBus.Unsubscribe(topic, handler)
	}

	return &accountSubscription{
		ch:      ch,
		onEvent: handler,
		unsubscribe: func() {
			unsubscriber()
			close(ch)
		},
	}, nil
}

func (w *wsClient) subscribeChannel(channel string, params interface{}, handler interface{}) (func(), error) {
	req := &subscriptionSignature{
		Cmd:     "subscribe",
		Channel: channel,
		Params:  params,
	}

	topic := toTopic(channel, params)
	if err := w.evBus.SubscribeAsync(topic, handler, true); err != nil {
		return nil, err
	}

	unsubscriber := func() {
		w.evBus.Unsubscribe(topic, handler)
	}

	return unsubscriber, w.sendMsg(req)
}

func (w *wsClient) sendMsg(msg interface{}) error {
	w.connMu.Lock()
	defer w.connMu.Unlock()

	return w.conn.WriteJSON(msg)
}

func (w *wsClient) readMsg(msg interface{}) error {
	w.connMu.RLock()
	defer w.connMu.RUnlock()

	return w.conn.ReadJSON(msg)
}

func (w *wsClient) handleMsg() {
	errCh := make(chan error, 1)
	for {
		resp := subscriptionResponse{}

		select {
		case errCh <- w.readMsg(&resp):
			if err := <-errCh; err != nil {
				w.logger.Printf("Failed to read JSON, %v\n", err)
				continue
			}

			w.handleResponse(resp)
		case <-w.stopCh:
			return
		}
	}
}

func (w *wsClient) handleResponse(resp subscriptionResponse) {
	switch resp["info"] {
	case "challenge":
		if w.accessKey == "" || w.secretKey == "" {
			w.logger.Println("Authentication disabled")
			return
		}

		msg, ok := resp["msg"]
		if ok {
			m, _ := msg.(string)
			authResp := struct {
				Cmd       string `json:"cmd,omitempty"`
				AccessKey string `json:"access_key,omitempty"`
				Answer    string `json:"answer,omitempty"`
			}{
				"auth",
				w.accessKey,
				signPayload([]byte(w.secretKey), []byte(w.accessKey+m)),
			}

			if err := w.sendMsg(authResp); err != nil {
				w.logger.Println("Authentication failed", err)
			}
		} else {
			m, _ := json.Marshal(resp)
			w.logger.Println("Invalid challenge message", string(m))
		}
	case "authenticated":
		w.logger.Println("Authenticated")
	case "account":
		go w.evBus.Publish("account", resp)
	case "subscribed":
		topic := toTopic(resp["channel"], map[string]interface{}{
			"market": resp["market"],
		})
		w.logger.Println(topic, "subscribed")
	case "ticker":
		ev := &tickerEventJSON{}

		if err := mapStruct(resp, &ev); err != nil {
			w.logger.Println("Failed to decode ticker response", err)
			return
		}

		topic := toTopic(resp["info"], map[string]interface{}{
			"market": ev.Market,
		})

		e, err := ev.Ticker()
		if err != nil {
			w.logger.Println("Failed to parse ticker event", err)
			return
		}

		go w.evBus.Publish(topic, e)
	case "orderbook":
		ev := &models.OrderBookEvent{}

		if err := mapStruct(resp, &ev); err != nil {
			w.logger.Println("Failed to decode orderbook response", err)
			return
		}

		topic := toTopic(resp["info"], map[string]interface{}{
			"market": ev.Market,
		})

		go w.evBus.Publish(topic, ev)
	case "trade":
		ev := &tradeEventJSON{}

		if err := mapStruct(resp, &ev); err != nil {
			w.logger.Println("Failed to decode trade response", err)
			return
		}

		topic := toTopic(resp["info"], map[string]interface{}{
			"market": ev.Market,
		})

		e, err := ev.Trade()
		if err != nil {
			w.logger.Println("Failed to parse trade event", err)
			return
		}

		go w.evBus.Publish(topic, e)
	default:
		b, _ := json.Marshal(resp)
		w.logger.Println("Unhandled message", b)
	}
}
