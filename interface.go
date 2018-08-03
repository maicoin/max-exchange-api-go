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
	"context"
	"time"

	"github.com/maicoin/max-exchange-api-go/models"
)

// API provides an interface to the MAX APIs
type API interface {
	PublicAPI
	PrivateAPI
}

// PublicAPI provides an interface to the public MAX APIs.
// Public APIs can be invoked without authentication and rate limits.
type PublicAPI interface {
	Markets(context.Context, ...CallOption) ([]*models.Market, error)
	Currencies(context.Context, ...CallOption) ([]*models.Currency, error)
	Ticker(context.Context, string, ...CallOption) (*models.Ticker, error)
	Tickers(context.Context, ...CallOption) (models.Tickers, error)
	OrderBook(context.Context, string, ...CallOption) (*models.OrderBook, error)
	Depth(context.Context, string, ...CallOption) (*models.Depth, error)
	Trades(context.Context, string, ...CallOption) ([]*models.Trade, error)
	K(context.Context, string, ...CallOption) ([]*models.Candle, error)
	Time(context.Context, ...CallOption) (time.Time, error)
}

// PrivateAPI provides an interface the private MAX APIs which
// has authtication requirements and rate limits.
type PrivateAPI interface {
	Me(context.Context, ...CallOption) (*models.Member, error)
	Deposit(context.Context, string, ...CallOption) (*models.Deposit, error)
	Deposits(context.Context, ...CallOption) ([]*models.Deposit, error)
	DepositAddress(context.Context, ...CallOption) ([]*models.PaymentAddress, error)
	CreateDepositAddresses(context.Context, string, ...CallOption) ([]*models.PaymentAddress, error)
	Withdrawal(context.Context, string, ...CallOption) (*models.Withdrawal, error)
	Withdrawals(context.Context, ...CallOption) ([]*models.Withdrawal, error)
	CreateOrders(context.Context, string, []*models.OrderRequest, ...CallOption) ([]*models.Order, error)
	CancelOrder(context.Context, int32, ...CallOption) (*models.Order, error)
	CancelOrders(context.Context, ...CallOption) ([]*models.Order, error)
	Order(context.Context, int32, ...CallOption) (*models.Order, error)
	Orders(context.Context, string, ...CallOption) ([]*models.Order, error)
	MyTrades(context.Context, string, ...CallOption) ([]*models.Trade, error)
}
