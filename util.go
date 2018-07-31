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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

func mapStruct(src interface{}, v interface{}) (err error) {
	pr, pw := io.Pipe()
	errCh := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer pw.Close()
		errCh <- json.NewEncoder(pw).Encode(src)
	}()

	go func() {
		defer wg.Done()
		errCh <- json.NewDecoder(pr).Decode(v)
	}()

	wg.Wait()
	close(errCh)

	e1 := <-errCh
	e2 := <-errCh

	if e1 != nil {
		err = fmt.Errorf("%v, %v", err, e1)
	}

	if e2 != nil {
		err = fmt.Errorf("%v, %v", err, e2)
	}

	return err
}

func signPayload(secret, payload []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}
