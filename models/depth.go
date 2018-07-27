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
	"encoding/json"
	"time"

	"github.com/maicoin/max-exchange-api-go/types"
)

// get depth of a specified market, sorted from highest price to lowest
type Depth struct {
	Timestamp time.Time  `json:"timestamp,omitempty"`
	Asks      []*Bargain `json:"asks,omitempty"`
	Bids      []*Bargain `json:"bids,omitempty"`
}

type Bargain struct {
	Price  types.Price
	Volume types.Volume
}

func (b *Bargain) MarshalJSON() ([]byte, error) {
	return json.Marshal([]float64{
		b.Price,
		b.Volume,
	})
}
