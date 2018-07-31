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

import "log"

type WebsocketClientOption func(*wsClient)

// WSAuthToken passes API tokens to the websocket client
func WSAuthToken(accessKey, secretKey string) WebsocketClientOption {
	return func(c *wsClient) {
		c.accessKey = accessKey
		c.secretKey = secretKey
	}
}

// WSURL sets the websocket URL to connect to
func WSURL(url string) WebsocketClientOption {
	return func(c *wsClient) {
		c.URL = url
	}
}

// WSLogging sets logger of the websocket client
func WSLogging(logger *log.Logger) WebsocketClientOption {
	return func(c *wsClient) {
		c.logger = logger
	}
}
