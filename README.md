[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/maicoin/max-exchange-api-go?status.svg)](https://godoc.org/github.com/maicoin/max-exchange-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/maicoin/max-exchange-api-go)](https://goreportcard.com/report/github.com/maicoin/max-exchange-api-go)

# max-exchange-api-go
MAX (Maicoin Assets eXchange) Go SDK

## Examples

See the `examples` directory.

## Documentation for API Endpoints

### Notes

Private APIs require authentication. Pass your API tokens by `AuthToken()` or `WSAuthToken()` before using them.

### RESTful APIs

All URIs are relative to *https://max-api.maicoin.com*

Class | Go Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Public* | [**Currencies**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/currencies |
*Public* | [**Depth**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/depth |
*Public* | [**K**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/k |
*Public* | [**Markets**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/markets |
*Public* | [**OrderBook**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/order_book |
*Public* | [**Tickers**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/tickers |
*Public* | [**Ticker**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/tickers/{market} |
*Public* | [**Timestamp**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/timestamp |
*Public* | [**Trades**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/trades |


Class | Go Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Private* | [**Deposit**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/deposit |
*Private* | [**DepositAddress**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/deposit_address | Deprecated
*Private* | [**CreateDepositAddresses**](https://max.maicoin.com/documents/api_list#/) | **POST** /api/v2/deposit_addresses | create deposit addresses
*Private* | [**DepositAddresses**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/deposit_addresses | where to deposit
*Private* | [**Deposits**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/deposits |
*Private* | [**Me**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/members/me |
*Private* | [**Order**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/order |
*Private* | [**Orders**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/orders |
*Private* | [**MyTrades**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/trades/my |
*Private* | [**Withdrawal**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/withdrawal |
*Private* | [**Withdrawals**](https://max.maicoin.com/documents/api_list#/) | **GET** /api/v2/withdrawals |
*Private* | [**CancelOrder**](https://max.maicoin.com/documents/api_list#/) | **POST** /api/v2/order/delete |
*Private* | [**CancelOrders**](https://max.maicoin.com/documents/api_list#/) | **POST** /api/v2/orders/clear |
*Private* | [**CreateOrder**](https://max.maicoin.com/documents/api_list#/) | **POST** /api/v2/orders |
*Private* | [**CreateOrders**](https://max.maicoin.com/documents/api_list#/) | **POST** /api/v2/orders/multi | create multiple sell/buy orders

### Websocket APIs (Beta)

Class | Go Method |  Description
------------ | ------------- | -------------
*Public* | [**SubscribeTicker**](https://max.maicoin.com/documents/websocket_api) | Subscribe the realtime price information
*Public* | [**SubscribeOrderBook**](https://max.maicoin.com/documents/websocket_api) | Subscribe the realtime changes on order books
*Public* | [**SubscribeTrade**](https://max.maicoin.com/documents/websocket_api) | Subscribe the realtime trades information

Class | Go Method |  Description
------------ | ------------- | -------------
*Private* | [**SubscribeAccount**](https://max.maicoin.com/documents/websocket_api) | Subscribe the accounts changes for an user

## API Reference

See [MAX RESTful API List](https://max.maicoin.com/documents/api_list#/)
and
[MAX Websocket API List (Beta)](https://max.maicoin.com/documents/websocket_api).
