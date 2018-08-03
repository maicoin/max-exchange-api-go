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

	"github.com/maicoin/max-exchange-api-go"
)

func main() {
	client := max.NewClient()
	defer client.Close()

	getMarkets(client)
	getTickers(client)
	getTrades(client)
}

func getTrades(client max.API) {
	results, err := client.Trades(context.Background(), "mithtwd")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Trades")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}

func getMarkets(client max.API) {
	results, err := client.Markets(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Markets")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}

func getTickers(client max.API) {
	results, err := client.Tickers(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Tickers")
	resultBytes, _ := json.MarshalIndent(results, "", "\t")
	fmt.Println(string(resultBytes))
}
