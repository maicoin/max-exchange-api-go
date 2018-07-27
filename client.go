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
	"net/http"
	"time"

	"github.com/maicoin/max-exchange-api-go/api"
)

func NewClient(opts ...ClientOption) *client {
	c := &client{
		requestTimeout: 10 * time.Second,
		cfg:            api.NewConfiguration(),
		middlewares:    make([]middleware, 0),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.c = api.NewAPIClient(c.config())

	return c
}

// Interface check
var _ PublicAPI = &publicClient{}
var _ PrivateAPI = &privateClient{}

type client struct {
	c              *api.APIClient
	requestTimeout time.Duration
	middlewares    []middleware

	cfg *api.Configuration
}

type middleware func(http.RoundTripper) http.RoundTripper

func (c *client) config() *api.Configuration {
	s := http.DefaultTransport

	for _, m := range c.middlewares {
		s = m(s)
	}

	c.cfg.HTTPClient = http.DefaultClient
	c.cfg.HTTPClient.Transport = s
	c.cfg.HTTPClient.Timeout = c.requestTimeout

	return c.cfg
}
