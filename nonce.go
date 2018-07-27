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
	"sync"
	"time"
)

var latestNonce int64
var mu sync.Mutex

func nonce() (n int64) {
	mu.Lock()
	defer mu.Unlock()
	defer func() {
		latestNonce = n
	}()

	return maximum(latestNonce, time.Now().Unix()*1000) + 1
}

func maximum(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
