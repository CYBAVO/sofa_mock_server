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
	"time"

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
	Count int64    `json:"count"`
	Memos []string `json:"memos"`
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
	Memo         string `json:"memo"`
}

type CallbackStruct struct {
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
	Memo        string                 `json:"memo"`
	BroadcastAt int64                  `json:"broadcast_at"`
	ChainAt     int64                  `json:"chain_at"`
	FromAddress string                 `json:"from_address"`
	ToAddress   string                 `json:"to_address"`
	WalletID    int64                  `json:"wallet_id"`
	State       int64                  `json:"state"`
	Addon       map[string]interface{} `json:"addon"`
}

type GetNotificationsResponse struct {
	Notifications []*CallbackStruct `json:"notifications"`
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
	Results map[string]int64 `json:"results"`
}

type CallbackResendRequest struct {
	NotificationID int64 `json:"notification_id"`
}

type CallbackResendResponse struct {
	Count int `json:"count"`
}

type GetTxAPITokenStatusResponse struct {
	APICode   string `json:"api_code"`
	ExpiresAt int64  `json:"exp"`
}

type GetTransactionHistoryResponse struct {
	TransactionCount int               `json:"transaction_count"`
	TransactionItem  []TransactionItem `json:"transaction_item"`
}

type ApprovalItem struct {
	ApprovalId   int64  `json:"approval_id"`
	ApprovalUser string `json:"approval_user"`
	ApprovalTime int64  `json:"approval_time"`
	UserMessage  string `json:"user_message"`
	Level        int    `json:"level"`
	Owner        int    `json:"owner"`
	Confirm      int    `json:"confirm"`
	State        int    `json:"state"`
	ErrorCode    int    `json:"error_code"`
}

type TransactionBatchStats struct {
	TransactionId    int64     `json:"transaction_id"`
	TotalAmount      string    `json:"total_amount"`
	TransactionCount int       `json:"transaction_count"`
	OutgoingCount    int       `json:"outgoing_count"`
	SuccessCount     int       `json:"success_count"`
	FailCount        int       `json:"fail_count"`
	SuccessAmount    string    `json:"success_amount"`
	CreateTime       time.Time `json:"create_time"`
}

type TransactionItem struct {
	IssueUserId         int64                  `json:"issue_user_id"`
	IssueUserName       string                 `json:"issue_user_name"`
	Description         string                 `json:"description"`
	WalletId            int64                  `json:"wallet_id"`
	WalletName          string                 `json:"wallet_name"`
	WalletAddress       string                 `json:"wallet_address"`
	TokenAddress        string                 `json:"token_address"`
	TxId                string                 `json:"txid"`
	Currency            int64                  `json:"currency"`
	CurrencyName        string                 `json:"currency_name"`
	OutgoingAddress     string                 `json:"outgoing_address"`
	OutgoingAddressName string                 `json:"outgoing_address_name"`
	Amount              string                 `json:"amount"`
	Fee                 string                 `json:"fee"`
	TxNo                int64                  `json:"txno"`
	ApprovalItem        []ApprovalItem         `json:"approval_item"`
	State               int                    `json:"state"`
	CreateTime          int64                  `json:"create_time"`
	TransactionTime     int64                  `json:"transaction_time"`
	ScheduledName       string                 `json:"scheduled_name"`
	TransactionType     int                    `json:"transaction_type"`
	Batch               *TransactionBatchStats `json:"batch,omitempty"`
	EosTransactionType  int                    `json:"eos_transaction_type"`
	RealAmount          string                 `json:"real_amount"`
	ChainFee            string                 `json:"chain_fee"`
	PlatformFee         string                 `json:"platform_fee"`
	TxCategory          string                 `json:"tx_category"`
	Memo                string                 `json:"memo"`
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

func GetTxAPITokenStatus(walletID int64) (response *GetTxAPITokenStatusResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/apisecret", walletID)

	resp, err := makeRequest(walletID, "GET", uri, nil, nil)
	if err != nil {
		return
	}

	response = &GetTxAPITokenStatusResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetTxAPITokenStatus() => ", response)
	return
}

func GetNotifications(walletID int64, fromTime int64, toTime int64, notificationType int) (response *GetNotificationsResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/notifications", walletID)

	params := []string{}
	params = append(params, fmt.Sprintf("from_time=%d", fromTime))
	params = append(params, fmt.Sprintf("to_time=%d", toTime))
	params = append(params, fmt.Sprintf("type=%d", notificationType))
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetNotificationsResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetNotifications() => ", response)
	return
}

func GetTransactionHistory(walletID int64, fromTime int64, toTime int64, startIndex int, requestNumber int, state int) (response *GetTransactionHistoryResponse, err error) {
	params := []string{}
	params = append(params, fmt.Sprintf("from_time=%d", fromTime))
	params = append(params, fmt.Sprintf("to_time=%d", toTime))
	params = append(params, fmt.Sprintf("start_index=%d", startIndex))
	params = append(params, fmt.Sprintf("request_number=%d", requestNumber))
	params = append(params, fmt.Sprintf("state=%d", state))
	resp, err := makeRequest(walletID, "GET", "/v1/sofa/transactions", params, nil)
	if err != nil {
		return
	}

	response = &GetTransactionHistoryResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetTransactionHistory() => ", response)
	return
}
