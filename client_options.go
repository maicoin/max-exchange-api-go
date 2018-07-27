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
	"time"
)

type ClientOption func(*client)

// AuthToken passess the access key and secret key to the API client.
func AuthToken(accessKey, secretKey string) ClientOption {
	return func(c *client) {
		c.middlewares = append(c.middlewares, newAuthMiddleware(accessKey, secretKey))
	}
}

// UserAgent sets the User-Agent in HTTP header.
func UserAgent(agent string) ClientOption {
	return func(c *client) {
		c.cfg.UserAgent = agent
	}
}

// BasePath sets base path of the API endpoint
func BasePath(path string) ClientOption {
	return func(c *client) {
		c.cfg.BasePath = path
	}
}

// Timeout sets the HTTP request timeout
func Timeout(t time.Duration) ClientOption {
	return func(c *client) {
		c.requestTimeout = t
	}
}

// Logging enables HTTP request logging
func Logging(logger *log.Logger) ClientOption {
	return func(c *client) {
		c.middlewares = append(c.middlewares, newLogMiddleware(logger))
	}
}
