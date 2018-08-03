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
	"fmt"
	"sync"
	"testing"
)

func TestNonce(t *testing.T) {
	var pivot sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			n := nonce()
			key := fmt.Sprintf("%d", n)

			if _, loaded := pivot.LoadOrStore(key, n); loaded {
				t.Errorf("Got equivalent nonce: %d\n", n)
			}
		}()
	}

	wg.Wait()
}
