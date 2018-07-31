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

package models

import (
	"time"

	"github.com/maicoin/max-exchange-api-go/types"
)

type TickerEvent struct {
	Market string `json:"market,omitempty"`
	Ticker
}

type OrderBookEvent struct {
	Action    string          `json:"action,omitempty"`
	Market    string          `json:"market,omitempty"`
	ID        int             `json:"id,omitempty"`
	Side      types.OrderSide `json:"side,omitempty"`
	Volume    types.Volume    `json:"volume,omitempty"`
	Price     types.Price     `json:"price,omitempty"`
	OrderType types.OrderType `json:"ord_type,omitempty"`
}

type TradeEvent struct {
	At     time.Time    `json:"at,omitempty"`
	Market string       `json:"market,omitempty"`
	Volume types.Volume `json:"volume,omitempty"`
	Price  types.Price  `json:"price,omitempty"`
}

type AccountEvent map[string]interface{}
