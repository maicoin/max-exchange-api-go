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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/maicoin/max-exchange-api-go/models"
	"github.com/maicoin/max-exchange-api-go/types"
)

type privateClient = client

// Me returns user profile and accounts information
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) Me(ctx context.Context, opts ...CallOption) (*models.Member, error) {
	member, _, err := c.c.PrivateApi.GetApiV2MembersMe(ctx, "", "", "")

	return &member, err
}

// Deposit returns details of the deposit with specific transaction ID.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) Deposit(ctx context.Context, txid string, opts ...CallOption) (*models.Deposit, error) {
	deposit, _, err := c.c.PrivateApi.GetApiV2Deposit(ctx, "", "", "", txid)

	return &deposit, err
}

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
func (c *privateClient) Deposits(ctx context.Context, opts ...CallOption) (results []*models.Deposit, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	deposits, _, err := c.c.PrivateApi.GetApiV2Deposits(ctx, "", "", "", o)
	for _, d := range deposits {
		d := d
		results = append(results, &d)
	}

	return results, err
}

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
func (c *privateClient) DepositAddress(ctx context.Context, opts ...CallOption) (results []*models.PaymentAddress, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	deposits, _, err := c.c.PrivateApi.GetApiV2DepositAddress(ctx, "", "", "", o)
	for _, d := range deposits {
		d := d
		results = append(results, &d)
	}

	return results, err
}

// DepositAddress returns the addresses that users are able to deposit.
//
// The address could be empty when a new one is generating, try again later in that case.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) DepositAddresses(ctx context.Context, opts ...CallOption) (results []*models.PaymentAddress, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	deposits, _, err := c.c.PrivateApi.GetApiV2DepositAddresses(ctx, "", "", "", o)
	for _, d := range deposits {
		d := d
		results = append(results, &d)
	}

	return results, err
}

// CreateDepositAddresses creates new addresses for deposit.
//
// Address creation is asynchronous, please call DepositAddresses later to get generated addresses
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) CreateDepositAddresses(ctx context.Context, currency string, opts ...CallOption) (results []*models.PaymentAddress, err error) {
	deposits, _, err := c.c.PrivateApi.PostApiV2DepositAddresses(ctx, "", "", "", currency)
	for _, d := range deposits {
		d := d
		results = append(results, &d)
	}

	return results, err
}

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
func (c *privateClient) Withdrawals(ctx context.Context, opts ...CallOption) (results []*models.Withdrawal, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	withdrawals, _, err := c.c.PrivateApi.GetApiV2Withdrawals(ctx, "", "", "", o)
	for _, w := range withdrawals {
		w := w
		results = append(results, &w)
	}

	return results, err
}

// Withdrawal returns the details of specific withdrawal.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) Withdrawal(ctx context.Context, uuid string, opts ...CallOption) (*models.Withdrawal, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	withdrawal, _, err := c.c.PrivateApi.GetApiV2Withdrawal(ctx, "", "", "", uuid)

	return &withdrawal, err
}

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
func (c *privateClient) CreateOrder(ctx context.Context, market string, side string, volumes types.Volume, opts ...CallOption) (*models.Order, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	order, _, err := c.c.PrivateApi.PostApiV2Orders(ctx, "", "", "", market, side, fmt.Sprintf("%v", volumes), o)

	return &order, err
}

// CreateOrders creates multiple sell/buy orders.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) CreateOrders(ctx context.Context, market string, orderRequests []*models.OrderRequest, opts ...CallOption) (results []*models.Order, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	body := make(map[string]interface{})

	body["market"] = market
	body["orders"] = orderRequests

	r, err := c.c.PrepareRequest(ctx,
		c.cfg.BasePath+"/api/v2/orders/multi",
		http.MethodPost,
		body, make(map[string]string), url.Values{}, url.Values{}, "", nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	resp, err := c.c.CallAPI(r)
	if err != nil || resp == nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Status: %v, Body: %s", resp.Status, bodyBytes)
	}

	if err = json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return results, err
	}

	return results, err
}

// CancelOrder cancels a sell/buy order.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) CancelOrder(ctx context.Context, id int32, opts ...CallOption) (*models.Order, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	order, _, err := c.c.PrivateApi.PostApiV2OrderDelete(ctx, "", "", "", id)

	return &order, err
}

// CancelOrders cancels a series of sell/buy orders.
//
// Available `CallOption`:
//     OrderSide(): set tp cancel only sell (asks) or buy (bids) orders
//     Market(): specify market like btctwd / ethbtc
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) CancelOrders(ctx context.Context, opts ...CallOption) (results []*models.Order, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	orders, _, err := c.c.PrivateApi.PostApiV2OrdersClear(ctx, "", "", "", o)
	for _, order := range orders {
		results = append(results, &order)
	}

	return results, err
}

// Order returns details of a specific order.
//
// Available `CallOption`:
//
// Note:
//     Use AuthToken() to pass your auth tokens.
func (c *privateClient) Order(ctx context.Context, id int32, opts ...CallOption) (*models.Order, error) {
	order, _, err := c.c.PrivateApi.GetApiV2Order(ctx, "", "", "", id)

	return &order, err
}

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
func (c *privateClient) Orders(ctx context.Context, market string, opts ...CallOption) (results []*models.Order, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	orders, _, err := c.c.PrivateApi.GetApiV2Orders(ctx, "", "", "", market, o)
	for _, order := range orders {
		order := order
		results = append(results, &order)
	}

	return results, err
}

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
func (c *privateClient) MyTrades(ctx context.Context, market string, opts ...CallOption) (results []*models.Trade, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	trades, _, err := c.c.PrivateApi.GetApiV2TradesMy(ctx, "", "", "", market, o)
	for _, t := range trades {
		t := t
		results = append(results, &t)
	}

	return results, err
}
