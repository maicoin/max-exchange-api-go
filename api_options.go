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
	"time"

	"github.com/maicoin/max-exchange-api-go/types"
)

var (
	orderDescending = "desc"
	orderAscending  = "asc"
)

type Options map[string]interface{}

func defaultOptions() Options {
	return make(map[string]interface{})
}

// ----------------------------------------------------------------------------

// CallOption represents the API parameters
type CallOption func(opt map[string]interface{})

// AsksLimit represents the asks_limit parameter
func AsksLimit(limit int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["asks_limit"] = limit
	}
}

// BidsLimit represents the bids_limit parameter
func BidsLimit(limit int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["bids_limit"] = limit
	}
}

// Limit represents the limit parameter
func Limit(limit int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["limit"] = limit
	}
}

// Timestamp represents the timestamp parameter
func Timestamp(timestamp int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["timestamp"] = timestamp
	}
}

// Time represents the timestamp parameter in Go time.Time format
func Time(t time.Time) CallOption {
	return func(opt map[string]interface{}) {
		opt["timestamp"] = t.Unix()
	}
}

// From represents the from parameter
func From(from int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["from"] = from
	}
}

// FromTime represents the from parameter in Go time.Time format
func FromTime(from time.Time) CallOption {
	return func(opt map[string]interface{}) {
		opt["from"] = from.Unix()
	}
}

// To represents the to parameter
func To(to int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["to"] = to
	}
}

// ToTime represents the to parameter
func ToTime(to time.Time) CallOption {
	return func(opt map[string]interface{}) {
		opt["to"] = to.Unix()
	}
}

// OrderDesc represents the order_by=desc parameter
func OrderDesc() CallOption {
	return func(opt map[string]interface{}) {
		opt["order_by"] = orderDescending
	}
}

// OrderAsc represents the order_by=asc parameter
func OrderAsc() CallOption {
	return func(opt map[string]interface{}) {
		opt["order_by"] = orderAscending
	}
}

// Period represents the period parameter
func Period(period int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["period"] = period
	}
}

// PeriodDuration represents the period parameter in Go time.Duration format
func PeriodDuration(period time.Duration) CallOption {
	return func(opt map[string]interface{}) {
		opt["period"] = int64(period.Minutes())
	}
}

// Currency represents the currency parameter
func Currency(currency string) CallOption {
	return func(opt map[string]interface{}) {
		opt["currency"] = currency
	}
}

// Offset represents the offset parameter
func Offset(offset int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["offset"] = offset
	}
}

// DepositState represents the state parameter for deposit
func DepositState(state types.DepositState) CallOption {
	return func(opt map[string]interface{}) {
		opt["state"] = state
	}
}

// WithdrawalState represents the state parameter for withdrawal
func WithdrawalState(state types.WithdrawalState) CallOption {
	return func(opt map[string]interface{}) {
		opt["state"] = state
	}
}

// Price represents the price parameter
func Price(price types.Price) CallOption {
	return func(opt map[string]interface{}) {
		opt["price"] = price
	}
}

// Prices represents the orders[price] parameter
func Prices(prices []types.Price) CallOption {
	return func(opt map[string]interface{}) {
		opt["orders[price]"] = prices
	}
}

// StopPrice represents the stop_price parameter
func StopPrice(price types.Price) CallOption {
	return func(opt map[string]interface{}) {
		opt["stop_price"] = price
	}
}

// StopPrices represents the orders[stop_price] parameter
func StopPrices(prices []types.Price) CallOption {
	return func(opt map[string]interface{}) {
		opt["orders[stop_price]"] = prices
	}
}

// OrderType represents the ord_type parameter
func OrderType(t types.OrderType) CallOption {
	return func(opt map[string]interface{}) {
		opt["ord_type"] = t
	}
}

// OrderTypes represents the orders[ord_type] parameter
func OrderTypes(t []types.OrderType) CallOption {
	return func(opt map[string]interface{}) {
		opt["orders[ord_type"] = t
	}
}

// Pagination represents the pagination parameter
func Pagination(pagination bool) CallOption {
	return func(opt map[string]interface{}) {
		opt["pagination"] = pagination
	}
}

// Page represents the page parameter
func Page(page int32) CallOption {
	return func(opt map[string]interface{}) {
		opt["page"] = page
	}
}

// OrderSide represents the side parameter
func OrderSide(t types.OrderSide) CallOption {
	return func(opt map[string]interface{}) {
		opt["side"] = t
	}
}

// Market represents the market parameter
func Market(market string) CallOption {
	return func(opt map[string]interface{}) {
		opt["market"] = market
	}
}
