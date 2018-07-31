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
	"encoding/json"
	"time"

	"github.com/maicoin/max-exchange-api-go/api"
	"github.com/maicoin/max-exchange-api-go/models"
	"github.com/maicoin/max-exchange-api-go/types"
)

type tickerJSON struct {
	// timestamp in seconds since Unix epoch
	At     json.Number `json:"at,omitempty"`
	Buy    json.Number `json:"buy,omitempty"`
	Sell   json.Number `json:"sell,omitempty"`
	Open   json.Number `json:"open,omitempty"`
	Last   json.Number `json:"last,omitempty"`
	High   json.Number `json:"high,omitempty"`
	Low    json.Number `json:"low,omitempty"`
	Volume json.Number `json:"vol,omitempty"`
}

func (t *tickerJSON) Ticker() (*models.Ticker, error) {
	ticker := &models.Ticker{}

	at, err := t.At.Int64()
	if err != nil {
		return nil, err
	}
	ticker.At = time.Unix(at, 0)
	ticker.Buy, err = t.Buy.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Sell, err = t.Sell.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Open, err = t.Open.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Last, err = t.Last.Float64()
	if err != nil {
		return nil, err
	}

	ticker.High, err = t.High.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Low, err = t.Low.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Volume, err = t.Volume.Float64()
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

type tmpTicker api.Ticker

func (t tmpTicker) Ticker() (result *models.Ticker, err error) {
	result = &models.Ticker{
		At: time.Unix(int64(t.At), 0),
	}

	result.Buy, err = types.ParsePrice(t.Buy)
	if err != nil {
		return nil, err
	}
	result.Sell, err = types.ParsePrice(t.Sell)
	if err != nil {
		return nil, err
	}
	result.Open, err = types.ParsePrice(t.Open)
	if err != nil {
		return nil, err
	}
	result.Last, err = types.ParsePrice(t.Last)
	if err != nil {
		return nil, err
	}
	result.High, err = types.ParsePrice(t.High)
	if err != nil {
		return nil, err
	}
	result.Low, err = types.ParsePrice(t.Low)
	if err != nil {
		return nil, err
	}
	result.Volume, err = types.ParseVolume(t.Vol)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type tmpTickers map[string]*tmpTicker

func (t tmpTickers) Tickers() (models.Tickers, error) {
	tickers := make(map[string]*models.Ticker)

	for name, tt := range t {
		ticker, err := tt.Ticker()
		if err != nil {
			return nil, err
		}

		tickers[name] = ticker
	}

	return tickers, nil
}

type depthJSON struct {
	Timestamp int             `json:"timestamp,omitempty"`
	Asks      [][]json.Number `json:"asks,omitempty"`
	Bids      [][]json.Number `json:"bids,omitempty"`
}

func (d *depthJSON) Depth() (*models.Depth, error) {
	depth := &models.Depth{
		Timestamp: time.Unix(int64(d.Timestamp), 0),
	}

	converter := func(n [][]json.Number) ([]*models.Bargain, error) {
		var err error
		results := make([]*models.Bargain, len(n))
		for i, x := range n {
			b := &models.Bargain{}

			b.Price, err = x[0].Float64()
			if err != nil {
				return nil, err
			}

			b.Volume, err = x[1].Float64()
			if err != nil {
				return nil, err
			}

			results[i] = b
		}

		return results, nil
	}

	var err error
	depth.Asks, err = converter(d.Asks)
	if err != nil {
		return nil, err
	}

	depth.Bids, err = converter(d.Bids)
	if err != nil {
		return nil, err
	}

	return depth, nil
}

type candleJSON []json.Number

func (c candleJSON) Candle() (*models.Candle, error) {
	candle := &models.Candle{}

	timestamp, err := c[0].Int64()
	if err != nil {
		return nil, err
	}

	candle.Time = time.Unix(timestamp, 0)
	candle.Open, err = c[1].Float64()
	if err != nil {
		return nil, err
	}
	candle.High, err = c[2].Float64()
	if err != nil {
		return nil, err
	}
	candle.Low, err = c[3].Float64()
	if err != nil {
		return nil, err
	}
	candle.Close, err = c[4].Float64()
	if err != nil {
		return nil, err
	}
	candle.Volume, err = c[5].Float64()
	if err != nil {
		return nil, err
	}

	return candle, nil
}

type candlesJSON []candleJSON

func (c candlesJSON) Candles() ([]*models.Candle, error) {
	candles := make([]*models.Candle, len(c))
	for i, cj := range c {
		candle, err := cj.Candle()
		if err != nil {
			return nil, err
		}

		candles[i] = candle
	}

	return candles, nil
}

type tickerEventJSON struct {
	At     json.Number `json:"at,omitempty"`
	Market string      `json:"market,omitempty"`
	Buy    json.Number `json:"buy,omitempty"`
	Sell   json.Number `json:"sell,omitempty"`
	Open   json.Number `json:"open,omitempty"`
	Last   json.Number `json:"last,omitempty"`
	High   json.Number `json:"high,omitempty"`
	Low    json.Number `json:"low,omitempty"`
	Volume json.Number `json:"vol,omitempty"`
}

func (t *tickerEventJSON) Ticker() (*models.TickerEvent, error) {
	ticker := &models.TickerEvent{}

	at, err := t.At.Int64()
	if err != nil {
		return nil, err
	}
	ticker.At = time.Unix(0, at*1000000)
	ticker.Market = t.Market
	ticker.Buy, err = t.Buy.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Sell, err = t.Sell.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Open, err = t.Open.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Last, err = t.Last.Float64()
	if err != nil {
		return nil, err
	}

	ticker.High, err = t.High.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Low, err = t.Low.Float64()
	if err != nil {
		return nil, err
	}

	ticker.Volume, err = t.Volume.Float64()
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

type tradeEventJSON struct {
	At     json.Number `json:"at,omitempty"`
	Market string      `json:"market,omitempty"`
	Volume json.Number `json:"volume,omitempty"`
	Price  json.Number `json:"price,omitempty"`
}

func (t tradeEventJSON) Trade() (*models.TradeEvent, error) {
	trade := &models.TradeEvent{}

	at, err := t.At.Int64()
	if err != nil {
		return nil, err
	}
	trade.At = time.Unix(0, at*1000000)
	trade.Market = t.Market
	trade.Price, err = t.Price.Float64()
	if err != nil {
		return nil, err
	}
	trade.Volume, err = t.Volume.Float64()
	if err != nil {
		return nil, err
	}

	return trade, nil
}
