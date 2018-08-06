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
	"github.com/maicoin/max-exchange-api-go/types"
)

// API provides an interface to the MAX APIs
type API interface {
	PublicAPI
	PrivateAPI
}

// PublicAPI provides an interface to the public MAX APIs.
// Public APIs can be invoked without authentication and rate limits.
type PublicAPI interface {
	// Markets returns available markets on MAX.
	//
	// Available `CallOption`:
	//
	Markets(context.Context, ...CallOption) ([]*models.Market, error)

	// Markets returns available currencies on MAX.
	//
	// Available `CallOption`:
	//
	Currencies(context.Context, ...CallOption) ([]*models.Currency, error)

	// Ticker returns a ticker of specific market.
	//
	// Available `CallOption`:
	//
	Ticker(context.Context, string, ...CallOption) (*models.Ticker, error)

	// Tickers returns tickers of all markets.
	//
	// Available `CallOption`:
	//
	Tickers(context.Context, ...CallOption) (models.Tickers, error)

	// OrderBook returns order books of specific market.
	//
	// Available `CallOption`:
	//     AsksLimit(): returned sell orders limit, default to 20
	//     BidsLimit(): returned buy orders limit, default to 20
	OrderBook(context.Context, string, ...CallOption) (*models.OrderBook, error)

	// Depth returns depth of specific market.
	//
	// Available `CallOption`:
	//     Limit(): returned price levels limit, default to 300
	Depth(context.Context, string, ...CallOption) (*models.Depth, error)

	// Trades returns recent trades on market.
	//
	// Available `CallOption`:
	//     Timestamp(): the seconds elapsed since Unix epoch, set to return trades executed before the time only
	//     Time(): the time in Go format, set to return trades executed before the time only
	//     From(): trade id, set ot return trades created after the trade
	//     To(): trade id, set to return trades created before the trade
	//     OrderDesc(): use descending order by created time, default value
	//     OrderAsc(): use ascending order by created time
	//     Pagination(): do pagination & return metadata in header (default true)
	//     Page(): page number, applied for pagination (default 1)
	//     Limit(): returned limit (1~1000, default 50)
	//     Offset(): records to skip, not applied for pagination (default 0)
	Trades(context.Context, string, ...CallOption) ([]*models.Trade, error)

	// K returns OHLC chart of specific market.
	//
	// Available `CallOption`:
	//     Timestamp(): the seconds elapsed since Unix epoch, set to return data after the timestamp only
	//     Time(): the time in Go format, set to return data after the time only
	//     Period(): time period of K line in minute, default to 1
	//     PeriodDuration(): time period of K line in time.Duration format, default to 1*time.Minute
	//     Limit(): returned data points limit, default to 30
	K(context.Context, string, ...CallOption) ([]*models.Candle, error)

	// Time returns current sever time.
	//
	// Available `CallOption`:
	//
	Time(context.Context, ...CallOption) (time.Time, error)
}

// PrivateAPI provides an interface the private MAX APIs which
// has authtication requirements and rate limits.
type PrivateAPI interface {
	// Me returns user profile and accounts information
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Me(context.Context, ...CallOption) (*models.Member, error)

	// Deposit returns details of the deposit with specific transaction ID.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Deposit(context.Context, string, ...CallOption) (*models.Deposit, error)

	// Deposits returns the history of your deposits.
	//
	// Available `CallOption`:
	//    Currency(): unique currency id, use Currencies() for available currencies.
	//    From(): target period start (Epoch time in seconds)
	//    FromTime(): target period start
	//    To(): target period end (Epoch time in seconds)
	//    ToTime(): target period end
	//    State(): the state of deposit
	//    Pagination(): do pagination & return metadata in header (default false)
	//    Page(): page number, applied for pagination (default 1)
	//    Limit(): returned limit (1~1000, default 50)
	//    Offset(): records to skip, not applied for pagination (default 0)
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Deposits(context.Context, ...CallOption) ([]*models.Deposit, error)

	// Deprecated: Use DepositAddresses instead.
	//
	// DepositAddress returns the addresses which are able to deposit.
	//
	// Available `CallOption`:
	//    Currency(): unique currency id, use Currencies() for available currencies.
	//
	// The address could be empty when a new one is generating, try again later in that case.
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	DepositAddress(context.Context, ...CallOption) ([]*models.PaymentAddress, error)

	// DepositAddress returns the addresses that users are able to deposit.
	//
	// The address could be empty when a new one is generating, try again later in that case.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	DepositAddresses(context.Context, ...CallOption) ([]*models.PaymentAddress, error)

	// CreateDepositAddresses creates new addresses for deposit.
	//
	// Address creation is asynchronous, please call DepositAddresses later to get generated addresses
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	CreateDepositAddresses(context.Context, string, ...CallOption) ([]*models.PaymentAddress, error)

	// Withdrawal returns the details of specific withdrawal.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Withdrawal(context.Context, string, ...CallOption) (*models.Withdrawal, error)

	// Withdrawals returns the withdrawals history.
	//
	// Available `CallOption`:
	//     Currency(): unique currency id, check Currencies() for available currencies
	//     From(): target period start (Epoch time in seconds)
	//     FromTime(): target period start
	//     To(): target period end (Epoch time in seconds)
	//     ToTime(): target period end
	//     State(): the state of withdrawals
	//     Pagniation(): do pagination & return metadata in header (default false)
	//     Page(): page number, applied for pagination (default 1)
	//     Limit(): returned limit (1~1000, default 50)
	//     Offset(): records to skip, not applied for pagination (default 0)
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Withdrawals(context.Context, ...CallOption) ([]*models.Withdrawal, error)

	// CreateOrder creates a sell/buy order.
	//
	// markets: unique market id, check Markets() for available markets
	// side: 'sell' or 'buy'
	// volume: total amount to sell/buy, an order could be partially executed
	//
	// Available `CallOption`:
	//    Price(): price per unit
	//    StopPrice(): price per unit to trigger a stop order
	//    OrderType(): `OrderTypeLimit`, `OrderTypeMarket`, `OrderTypeStopLimit`, or `OrderTypeStopMarket`
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	CreateOrder(context.Context, string, string, types.Volume, ...CallOption) (*models.Order, error)

	// CreateOrders creates multiple sell/buy orders.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	CreateOrders(context.Context, string, []*models.OrderRequest, ...CallOption) ([]*models.Order, error)

	// CancelOrder cancels a sell/buy order.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	CancelOrder(context.Context, int32, ...CallOption) (*models.Order, error)

	// CancelOrders cancels a series of sell/buy orders.
	//
	// Available `CallOption`:
	//     OrderSide(): set tp cancel only sell (asks) or buy (bids) orders
	//     Market(): specify market like btctwd / ethbtc
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	CancelOrders(context.Context, ...CallOption) ([]*models.Order, error)

	// Order returns details of a specific order.
	//
	// Available `CallOption`:
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Order(context.Context, int32, ...CallOption) (*models.Order, error)

	// Orders returns your orders.
	//
	// Available `CallOption`:
	//     State(): filter by state, default to 'OrderStateWait'
	//     OrderDesc(): use descending order by created time
	//     OrderAsc(): use ascending order by created time, default value
	//     Pagination(): do pagination & return metadata in header (default true)
	//     Page(): page number, applied for pagination (default 1)
	//     Limit(): returned limit (1~1000, default 100)
	//     Offset(): records to skip, not applied for pagination (default 0)
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	Orders(context.Context, string, ...CallOption) ([]*models.Order, error)

	// MyTrades returns the executed trades which are sorted in reverse creation order.
	//
	// Available `CallOption`:
	//     Timestamp(): the seconds elapsed since Unix epoch, set to return trades executed before the time only
	//     Time(): the time in Go format, set to return trades executed before the time only
	//     From(): trade id, set ot return trades created after the trade
	//     To(): trade id, set to return trades created before the trade
	//     OrderDesc(): use descending order by created time, default value
	//     OrderAsc(): use ascending order by created time
	//     Pagination(): do pagination & return metadata in header (default true)
	//     Page(): page number, applied for pagination (default 1)
	//     Limit(): returned limit (1~1000, default 50)
	//     Offset(): records to skip, not applied for pagination (default 0)
	//
	// Note:
	//     Use AuthToken() to pass your auth tokens.
	MyTrades(context.Context, string, ...CallOption) ([]*models.Trade, error)
}
