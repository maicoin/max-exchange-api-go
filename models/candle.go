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

type Candle struct {
	Time   time.Time   `json:"timestamp,omitempty"`
	Open   types.Price `json:"open,omitempty"`
	High   types.Price `json:"high,omitempty"`
	Low    types.Price `json:"low,omitempty"`
	Close  types.Price `json:"close,omitempty"`
	Volume types.Price `json:"volume,omitempty"`
}

func (c *Candle) Array() []float64 {
	result := make([]float64, 6)
	result[0] = float64(c.Time.Unix())
	result[1] = c.Open
	result[2] = c.High
	result[3] = c.Low
	result[4] = c.Close
	result[5] = c.Volume

	return result
}
