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
	"time"

	"github.com/maicoin/max-exchange-api-go/models"
)

type publicClient = client

// Markets returns available markets on MAX.
//
// Available `CallOption`:
//
func (c *publicClient) Markets(ctx context.Context, opts ...CallOption) (results []*models.Market, err error) {
	markets, _, err := c.c.PublicApi.GetApiV2Markets(ctx)

	for _, m := range markets {
		m := m
		results = append(results, &m)
	}

	return results, err
}

// Markets returns available currencies on MAX.
//
// Available `CallOption`:
//
func (c *publicClient) Currencies(ctx context.Context, opts ...CallOption) (results []*models.Currency, err error) {
	currencies, _, err := c.c.PublicApi.GetApiV2Currencies(ctx)

	for _, c := range currencies {
		c := c
		results = append(results, &c)
	}

	return results, err
}

// Ticker returns a ticker of specific market.
//
// Available `CallOption`:
//
func (c *publicClient) Ticker(ctx context.Context, market string, opts ...CallOption) (*models.Ticker, error) {
	ticker, _, err := c.c.PublicApi.GetApiV2TickersMarket(ctx, market)
	if err != nil {
		return nil, err
	}

	return tmpTicker(ticker).Ticker()
}

// Tickers returns tickers of all markets.
//
// Available `CallOption`:
//
func (c *publicClient) Tickers(ctx context.Context, opts ...CallOption) (models.Tickers, error) {
	tickers, _, err := c.c.PublicApi.GetApiV2Tickers(ctx)
	if err != nil {
		return nil, err
	}

	tt := tmpTickers{}
	err = mapStruct(tickers, &tt)
	if err != nil {
		return nil, err
	}

	return tt.Tickers()
}

// OrderBook returns order books of specific market.
//
// Available `CallOption`:
//     AsksLimit(): returned sell orders limit, default to 20
//     BidsLimit(): returned buy orders limit, default to 20
func (c *publicClient) OrderBook(ctx context.Context, market string, opts ...CallOption) (*models.OrderBook, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	orderbook, _, err := c.c.PublicApi.GetApiV2OrderBook(ctx, market, o)

	return &orderbook, err
}

// Depth returns depth of specific market.
//
// Available `CallOption`:
//     Limit(): returned price levels limit, default to 300
func (c *publicClient) Depth(ctx context.Context, market string, opts ...CallOption) (*models.Depth, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	resp, err := c.c.PublicApi.GetApiV2Depth(ctx, market, o)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	depth := &depthJSON{}
	err = json.NewDecoder(resp.Body).Decode(&depth)
	if err != nil {
		return nil, err
	}

	return depth.Depth()
}

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
func (c *publicClient) Trades(ctx context.Context, market string, opts ...CallOption) (results []*models.Trade, err error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	trades, _, err := c.c.PublicApi.GetApiV2Trades(ctx, market, o)
	for _, t := range trades {
		t := t
		results = append(results, &t)
	}

	return results, err
}

// K returns OHLC chart of specific market.
//
// Available `CallOption`:
//     Timestamp(): the seconds elapsed since Unix epoch, set to return data after the timestamp only
//     Time(): the time in Go format, set to return data after the time only
//     Period(): time period of K line in minute, default to 1
//     PeriodDuration(): time period of K line in time.Duration format, default to 1*time.Minute
//     Limit(): returned data points limit, default to 30
func (c *publicClient) K(ctx context.Context, market string, opts ...CallOption) ([]*models.Candle, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	resp, err := c.c.PublicApi.GetApiV2K(ctx, market, o)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	candles := candlesJSON{}
	err = json.NewDecoder(resp.Body).Decode(&candles)
	if err != nil {
		return nil, err
	}

	return candles.Candles()
}

// Time returns current sever time.
//
// Available `CallOption`:
//
func (c *publicClient) Time(ctx context.Context, opts ...CallOption) (time.Time, error) {
	resp, err := c.c.PublicApi.GetApiV2Timestamp(ctx)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	var t int64
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(t, 0), nil
}
