// Copyright (c) 2018-2019 The Cybavo developers
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
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"
)

type CommonResponse struct {
	Result int64 `json:"result"`
}

type ErrorCodeResponse struct {
	ErrMsg  string `json:"error,omitempty"`
	ErrCode int    `json:"error_code,omitempty"`
}

type SetAPICodeRequest struct {
	APICode   string `json:"api_code"`
	ApiSecret string `json:"api_secret"`
}

type GetAPICodeResponse struct {
	APICode   string `json:"api_code"`
	ApiSecret string `json:"api_secret"`
}

type CreateDepositWalletAddressesRequest struct {
	Count int64 `json:"count"`
}

type CreateDepositWalletAddressesResponse struct {
	Addresses []string `json:"addresses"`
}

type GetDepositWalletAddressesResponse struct {
	WalletId      int64           `json:"wallet_id"`
	WalletAddress []WalletAddress `json:"wallet_address"`
}

type GetDepositWalletPoolAddressResponse struct {
	Address string `json:"address"`
}

type WalletAddress struct {
	Currency     int64  `json:"currency"`
	TokenAddress string `json:"token_address"`
	Address      string `json:"address"`
}

type CallbackRequest struct {
	Type        int                    `json:"type"`
	Serial      int64                  `json:"serial"`
	OrderID     string                 `json:"order_id"`
	Currency    string                 `json:"currency"`
	TXID        string                 `json:"txid"`
	BlockHeight int64                  `json:"block_height"`
	TIndex      int                    `json:"tindex"`
	VOutIndex   int                    `json:"vout_index"`
	Amount      string                 `json:"amount"`
	Fees        string                 `json:"fees"`
	BroadcastAt int64                  `json:"broadcast_at"`
	ChainAt     int64                  `json:"chain_at"`
	Address     string                 `json:"address"`
	WalletID    int64                  `json:"wallet_id"`
	Addon       map[string]interface{} `json:"addon"`
}

type WithdrawTransaction struct {
	OrderID string `json:"order_id"`
	Address string `json:"address"`
	Amount  string `json:"amount"`
	Memo    string `json:"memo"`
}

type WithdrawTransactionRequest struct {
	Requests []WithdrawTransaction `json:"requests"`
}

type WithdrawTransactionResponse struct {
	Results map[int64]int64 `json:"results"`
}

type CallbackResendRequest struct {
	NotificationID int64 `json:"notification_id"`
}

type CallbackResendResponse struct {
	Count int `json:"count"`
}

func CreateDepositWalletAddresses(walletID int64, request *CreateDepositWalletAddressesRequest) (response *CreateDepositWalletAddressesResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/addresses", walletID)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}
	resp, err := makeRequest(walletID, "POST", uri, nil, jsonRequest)
	if err != nil {
		return
	}

	response = &CreateDepositWalletAddressesResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("CreateDepositWalletAddresses() => ", response)
	return
}

func GetDepositWalletAddresses(walletID int64, startIndex int, requestNumber int) (response *GetDepositWalletAddressesResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/addresses", walletID)

	params := []string{}
	params = append(params, fmt.Sprintf("start_index=%d", startIndex))
	params = append(params, fmt.Sprintf("request_number=%d", requestNumber))
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetDepositWalletAddressesResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetDepositWalletAddresses() => ", response)
	return
}

func GetDepositWalletPoolAddress(walletID int64) (response *GetDepositWalletPoolAddressResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/pooladdress", walletID)

	params := []string{}
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetDepositWalletPoolAddressResponse{}
	err = json.Unmarshal(resp, response)

	logs.Error("GetDepositWalletPoolAddress() => ", response)
	return
}

func ResendCallback(walletID int64, request *CallbackResendRequest) (response *CallbackResendResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/collection/notifications/manual", walletID)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := makeRequest(walletID, "POST", uri, nil, jsonRequest)
	if err != nil {
		return
	}

	response = &CallbackResendResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return
	}

	logs.Debug("ResendCallback() => ", response)
	return
}

func WithdrawTransactions(walletID int64, request *WithdrawTransactionRequest) (response *WithdrawTransactionResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/sender/transactions", walletID)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := makeRequest(walletID, "POST", uri, nil, jsonRequest)
	if err != nil {
		result := &ErrorCodeResponse{}
		_ = json.Unmarshal(resp, result)
		err = errors.New(result.ErrMsg)
		return
	}

	response = &WithdrawTransactionResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return
	}

	logs.Debug("WithdrawTransactions() => ", response)
	return
}
