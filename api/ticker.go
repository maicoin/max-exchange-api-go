/*
 * MAX RESTful API List
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// get ticker of specific market
type Ticker struct {

	// timestamp in seconds since Unix epoch
	At int32 `json:"at,omitempty"`

	// highest buy price
	Buy string `json:"buy,omitempty"`

	// lowest sell price
	Sell string `json:"sell,omitempty"`

	// price before 24 hours
	Open string `json:"open,omitempty"`

	// lowest price within 24 hours
	Low string `json:"low,omitempty"`

	// highest price within 24 hours
	High string `json:"high,omitempty"`

	// last traded price
	Last string `json:"last,omitempty"`

	// traded volume within 24 hours
	Vol string `json:"vol,omitempty"`
}
