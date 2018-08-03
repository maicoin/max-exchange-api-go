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
	"github.com/maicoin/max-exchange-api-go/types"
)

var (
	DepositStateSubmitting      types.DepositState = "submitting"
	DepositStateCancelled       types.DepositState = "cancelled"
	DepositStateSubmitted       types.DepositState = "submitted"
	DepositStateSuspended       types.DepositState = "suspended"
	DepositStateRejected        types.DepositState = "rejected"
	DepositStateAccepted        types.DepositState = "accepted"
	DepositStateRefunded        types.DepositState = "refunded"
	DepositStateSuspect         types.DepositState = "suspect"
	DepositStateRefundCancelled types.DepositState = "refund_cancelled"
)

var (
	WithdrawalStateSubmitting types.WithdrawalState = "submitting"
	WithdrawalStateSubmitted  types.WithdrawalState = "submitted"
	WithdrawalStateRejected   types.WithdrawalState = "rejected"
	WithdrawalStateAccepted   types.WithdrawalState = "accepted"
	WithdrawalStateSuspect    types.WithdrawalState = "suspect"
	WithdrawalStateApproved   types.WithdrawalState = "approved"
	WithdrawalStateProcessing types.WithdrawalState = "processing"
	WithdrawalStateRetryable  types.WithdrawalState = "retryable"
	WithdrawalStateSent       types.WithdrawalState = "sent"
	WithdrawalStateCancelled  types.WithdrawalState = "cancelled"
	WithdrawalStateFailed     types.WithdrawalState = "failed"
	WithdrawalStatePending    types.WithdrawalState = "pending"
	WithdrawalStateConfirmed  types.WithdrawalState = "confirmed"
)

var (
	OrderTypeLimit      types.OrderType = "limit"
	OrderTypeMarket     types.OrderType = "market"
	OrderTypeStopLimit  types.OrderType = "stop_limit"
	OrderTypeStopMarket types.OrderType = "stop_market"
)

var (
	OrderSideSell types.OrderSide = "sell"
	OrderSideBuy  types.OrderSide = "buy"
)

var (
	OrderStateWait    types.OrderState = "wait"
	OrderStateDone    types.OrderState = "done"
	OrderStateConvert types.OrderState = "convert"
	OrderStateCancel  types.OrderState = "cancel"
)
