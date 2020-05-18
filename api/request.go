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

type ProcessingState int8

const (
	ProcessingStateUndefined ProcessingState = -1
	ProcessingStateInPool    ProcessingState = 0
	ProcessingStateInChain   ProcessingState = 1
	ProcessingStateDone      ProcessingState = 2
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
	State           int64                  `json:"state"`
	ConfirmBlocks   int64                  `json:"confirm_blocks"`
	ProcessingState ProcessingState        `json:"processing_state"`
	Addon           map[string]interface{} `json:"addon"`
}
