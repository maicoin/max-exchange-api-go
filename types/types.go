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

package types

import (
	"strconv"
)

type Price = float64

type Volume = float64

type Timestamp = int32

func ParsePrice(s string) (Price, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseVolume(s string) (Volume, error) {
	return strconv.ParseFloat(s, 64)
}

// DepositState represents the deposit state
type DepositState = string

// OrderType represtns the order type, e.g., limit, market, etc.
type OrderType = string

// WithdrawalState represents the withdrawal state
type WithdrawalState = string

// OrderSide indicates the order is a sell order or buy order
type OrderSide = string

// DepositState represents the order state
type OrderState = string
