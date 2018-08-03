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

	"github.com/maicoin/max-exchange-api-go/models"
)

type subscriptionSignature struct {
	Cmd     string      `json:"cmd,omitempty"`
	Channel string      `json:"channel,omitempty"`
	Params  interface{} `json:"params,omitempty"`
}

type subscriptionResponse map[string]interface{}

func toTopic(channel interface{}, params interface{}) string {
	b, _ := json.Marshal([]interface{}{
		channel,
		params,
	})

	return string(b)
}

type TickerSubscription interface {
	Chan() <-chan *models.TickerEvent
	Close()
}

type tickerSubscription struct {
	ch          <-chan *models.TickerEvent
	onEvent     func(ticker *models.TickerEvent)
	unsubscribe func()
}

func (s *tickerSubscription) Chan() <-chan *models.TickerEvent {
	return s.ch
}

func (s *tickerSubscription) Close() {
	s.unsubscribe()
}

type OrderBookSubscription interface {
	Chan() <-chan *models.OrderBookEvent
	Close()
}

type orderBookSubscription struct {
	ch          <-chan *models.OrderBookEvent
	onEvent     func(ob *models.OrderBookEvent)
	unsubscribe func()
}

func (s *orderBookSubscription) Chan() <-chan *models.OrderBookEvent {
	return s.ch
}

func (s *orderBookSubscription) Close() {
	s.unsubscribe()
}

type TradeSubscription interface {
	Chan() <-chan *models.TradeEvent
	Close()
}

type tradeSubscription struct {
	ch          <-chan *models.TradeEvent
	onEvent     func(t *models.TradeEvent)
	unsubscribe func()
}

func (s *tradeSubscription) Chan() <-chan *models.TradeEvent {
	return s.ch
}

func (s *tradeSubscription) Close() {
	s.unsubscribe()
}

type AccountSubscription interface {
	Chan() <-chan models.AccountEvent
	Close()
}

type accountSubscription struct {
	ch          <-chan models.AccountEvent
	onEvent     func(t models.AccountEvent)
	unsubscribe func()
}

func (s *accountSubscription) Chan() <-chan models.AccountEvent {
	return s.ch
}

func (s *accountSubscription) Close() {
	s.unsubscribe()
}
