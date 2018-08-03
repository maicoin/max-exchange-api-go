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
	"log"
	"net/http"
	"net/http/httputil"
)

func newLogMiddleware(logger *log.Logger) middleware {
	return func(n http.RoundTripper) http.RoundTripper {
		return logMiddleware{
			next:   n,
			logger: logger,
		}
	}
}

type logMiddleware struct {
	next   http.RoundTripper
	logger *log.Logger
}

func (m logMiddleware) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	if dump, err := httputil.DumpRequest(req, true); err == nil {
		m.logger.Println(string(dump))
	}

	defer func() {
		if dump, err := httputil.DumpResponse(resp, true); err == nil {
			m.logger.Println(string(dump))
		}
	}()

	return m.next.RoundTrip(req)
}
