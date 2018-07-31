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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	HeaderAccessKey  = "X-MAX-ACCESSKEY"
	HeaderPayloadKey = "X-MAX-PAYLOAD"
	HeaderSignature  = "X-MAX-SIGNATURE"
)

func newAuthMiddleware(accessKey, secretKey string) middleware {
	return func(n http.RoundTripper) http.RoundTripper {
		return authMiddleware{
			accessKey: accessKey,
			secretKey: secretKey,
			next:      n,
		}
	}
}

type authMiddleware struct {
	accessKey string
	secretKey string
	next      http.RoundTripper
}

func (m authMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	params := make(map[string]interface{})
	params["path"] = req.URL.Path
	params["nonce"] = nonce()

	if req.Body != nil {
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			return nil, err
		}
	}

	body, _ := json.Marshal(params)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	payload := base64.StdEncoding.EncodeToString(body)

	// Write auth headers
	req.Header.Set(HeaderAccessKey, m.accessKey)
	req.Header.Set(HeaderPayloadKey, payload)
	req.Header.Set(HeaderSignature, signPayload([]byte(m.secretKey), []byte(payload)))

	req.RequestURI = req.URL.Path
	req.ContentLength = int64(len(body))

	return m.next.RoundTrip(req)
}
