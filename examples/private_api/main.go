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
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/maicoin/max-exchange-api-go"
	"github.com/maicoin/max-exchange-api-go/models"
)

func main() {
	// logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := max.NewClient(
		// max.Logging(logger),
		max.AuthToken(os.Getenv("MAX_ACCESS_KEY"), os.Getenv("MAX_SECRET_KEY")),
	)
	defer client.Close()

	createStopMarketSellOrder(client)
	cancelAllSellOrder(client)
	getMyTrades(client)
}

func getMyTrades(client max.API) {
	results, err := client.MyTrades(context.Background(), "mithtwd")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("My trades")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}

func cancelAllSellOrder(client max.API) {
	results, err := client.CancelOrders(context.Background(), max.OrderSide(max.OrderSideSell))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("cancelAllSellOrder")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}

func createStopMarketSellOrder(client max.API) {
	results, err := client.CreateOrders(context.Background(), "mithtwd", []*models.OrderRequest{
		&models.OrderRequest{
			Side:      max.OrderSideSell,
			Volume:    10.0,
			StopPrice: 15.0,
			OrderType: max.OrderTypeStopMarket,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("createStopMarketSellOrder")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}
