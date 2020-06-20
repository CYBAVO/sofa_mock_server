// Copyright (c) 2018-2020 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

package api

import (
	"encoding/json"
	"fmt"
)

type CommonResponse struct {
	Result int64 `json:"result"`
}

type ErrorCodeResponse struct {
	ErrMsg    string              `json:"error,omitempty"`
	ErrCode   int                 `json:"error_code,omitempty"`
	Message   string              `json:"message,omitempty"`
	Blacklist map[string][]string `json:"blacklist,omitempty"`
}

func (m *ErrorCodeResponse) String() string {
	if len(m.Blacklist) > 0 {
		blacklist, _ := json.Marshal(m.Blacklist)
		return fmt.Sprintf("%s (msg:%s) (code:%d)", m.ErrMsg, string(blacklist), m.ErrCode)
	}
	return fmt.Sprintf("%s (msg:%s) (code:%d)", m.ErrMsg, m.Message, m.ErrCode)
}

type SetAPICodeRequest struct {
	APICode   string `json:"api_code"`
	ApiSecret string `json:"api_secret"`
}

type CallbackType int

const (
	DepositCallback  CallbackType = 1
	WithdrawCallback CallbackType = 2
	CollectCallback  CallbackType = 3
	AirdropCallback  CallbackType = 4
)

type ProcessingState int8

const (
	ProcessingStateUndefined ProcessingState = -1
	ProcessingStateInPool    ProcessingState = 0
	ProcessingStateInChain   ProcessingState = 1
	ProcessingStateDone      ProcessingState = 2
)

type CallbackState int64

const (
	CallbackStateInit            CallbackState = 0  // enqueue (0)
	CallbackStateHolding         CallbackState = 1  // Processing batch in KMS (1)
	CallbackStateInPool          CallbackState = 2  // KMS process done, TXID created (2)
	CallbackStateInChain         CallbackState = 3  // TXID in chain (3)
	CallbackStateDone            CallbackState = 4  // TXID confirmed in N blocks (4)
	CallbackStateFailed          CallbackState = 5  // Failed (5)
	CallbackStateResended        CallbackState = 6  // Resended (6)
	CallbackStateRiskControl     CallbackState = 7  // blocked due to risk controlled (7)
	CallbackStateCancelled       CallbackState = 8  // cancelled
	CallbackStateUTXOUnavailable CallbackState = 9  // Retry for UTXO Temporarily Not Available
	CallbackStateDropped         CallbackState = 10 // Dropped
	CallbackStateInChainFailed   CallbackState = 11 // Transaction Failed (11)
	CallbackStatePaused          CallbackState = 12 // Paused (12)
)

type CallbackStruct struct {
	Type            int                    `json:"type"`
	Serial          int64                  `json:"serial"`
	OrderID         string                 `json:"order_id"`
	Currency        string                 `json:"currency"`
	TXID            string                 `json:"txid"`
	BlockHeight     int64                  `json:"block_height"`
	TIndex          int                    `json:"tindex"`
	VOutIndex       int                    `json:"vout_index"`
	Amount          string                 `json:"amount"`
	Fees            string                 `json:"fees"`
	Memo            string                 `json:"memo"`
	BroadcastAt     int64                  `json:"broadcast_at"`
	ChainAt         int64                  `json:"chain_at"`
	FromAddress     string                 `json:"from_address"`
	ToAddress       string                 `json:"to_address"`
	WalletID        int64                  `json:"wallet_id"`
	State           CallbackState          `json:"state"`
	ConfirmBlocks   int64                  `json:"confirm_blocks"`
	ProcessingState ProcessingState        `json:"processing_state"`
	Addon           map[string]interface{} `json:"addon"`
}
