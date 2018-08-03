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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/maicoin/max-exchange-api-go"
	"github.com/maicoin/max-exchange-api-go/models"
)

func main() {
	client, err := max.NewWSClient(
		max.WSAuthToken(os.Getenv("MAX_ACCESS_KEY"), os.Getenv("MAX_SECRET_KEY")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	tickerSub, err := client.SubscribeTicker("mithtwd", make(chan *models.TickerEvent, 10))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tickerSub.Close()

	orderBookSub, err := client.SubscribeOrderBook("mithtwd", make(chan *models.OrderBookEvent, 10))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer orderBookSub.Close()

	tradeSub, err := client.SubscribeTrade("mithtwd", make(chan *models.TradeEvent, 10))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tradeSub.Close()

	accountSub, err := client.SubscribeAccount(make(chan models.AccountEvent, 10))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer accountSub.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(sigs)

	for {
		select {
		case ticker := <-tickerSub.Chan():
			b, _ := json.Marshal(ticker)
			fmt.Printf("Ticker: %v\n", string(b))
		case orderBook := <-orderBookSub.Chan():
			b, _ := json.Marshal(orderBook)
			fmt.Printf("OrderBook: %v\n", string(b))
		case trade := <-tradeSub.Chan():
			b, _ := json.Marshal(trade)
			fmt.Printf("Trade: %v\n", string(b))
		case account := <-accountSub.Chan():
			b, _ := json.Marshal(account)
			fmt.Printf("Account: %v\n", string(b))
		case sig := <-sigs:
			fmt.Println("Shutting down", "signal", sig)
			os.Exit(0)
		}
	}
}
