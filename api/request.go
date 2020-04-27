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
	Message string `json:"message,omitempty"`
}

func (m *ErrorCodeResponse) String() string {
	return fmt.Sprintf("%s (msg:%s) (code:%d)", m.ErrMsg, m.Message, m.ErrCode)
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

type GetNotificationsResponse struct {
	Notifications []*CallbackStruct `json:"notifications"`
}

type GetNotificationsByIDRequest struct {
	IDs []int64 `json:"ids"`
}

type WithdrawTransaction struct {
	OrderID         string `json:"order_id"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	Memo            string `json:"memo"`
	UserID          string `json:"user_id"`
	Message         string `json:"message"`
	BlockAverageFee *int   `json:"block_average_fee"`
	ManualFee       *int   `json:"manual_fee"`
}

type WithdrawTransactionRequest struct {
	Requests []WithdrawTransaction `json:"requests"`
}

type WithdrawTransactionResponse struct {
	Results map[string]int64 `json:"results"`
}

type GetWithdrawTransactionStateResponse struct {
	OrderID      string    `json:"order_id"`
	Address      string    `json:"address"`
	Amount       string    `json:"amount"`
	Memo         string    `json:"memo"`
	InChainBlock int64     `json:"in_chain_block"`
	TxID         string    `json:"txid"`
	CreateTime   time.Time `json:"create_time"`
}

type GetWithdrawalWalletBalanceResponse struct {
	Currency              int64  `json:"currency"`
	WalletAddress         string `json:"wallet_address"`
	TokenAddress          string `json:"token_address"`
	Balance               string `json:"balance"`
	TokenBalance          string `json:"token_balance"`
	UnconfirmBalance      string `json:"unconfirm_balance"`
	UnconfirmTokenBalance string `json:"unconfirm_token_balance"`
	ErrReason             string `json:"err_reason,omitempty"`
}

type CallbackResendRequest struct {
	NotificationID int64 `json:"notification_id"`
}

type CallbackResendResponse struct {
	Count int `json:"count"`
}

type GetTxAPITokenStatusResponse struct {
	Valid       *WalletApiCode `json:"valid,omitempty"`
	Inactivated *WalletApiCode `json:"inactivated,omitempty"`
}

type WalletApiCode struct {
	APICode   string `json:"api_code,omitempty"`
	ApiSecret string `json:"api_secret,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Status    int    `json:"status,omitempty"`
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

type GetWalletBlockInfoResponse struct {
	LatestBlockHeight int64 `json:"latest_block_height"`
	SyncedBlockHeight int64 `json:"synced_block_height"`
}

type GetInvalidDepositAddressesResponse struct {
	Addresses []string `json:"addresses"`
}

type GetWalletInfoResponse struct {
	Currency             int64  `json:"currency"`
	CurrencyName         string `json:"currency_name"`
	Address              string `json:"address"`
	TokenName            string `json:"token_name,omitempty"`
	TokenSymbol          string `json:"token_symbol,omitempty"`
	TokenContractAddress string `json:"token_contract_address,omitempty"`
	TokenDecimals        string `json:"token_decimals,omitempty"`
}

type VerifyAddressesRequest struct {
	Addresses []string `json:"addresses"`
}

type AddressStatus struct {
	Address      string `json:"address"`
	Valid        bool   `json:"valid"`
	MustNeedMemo bool   `json:"must_need_memo"`
}

type VerifyAddressesResponse struct {
	Result []*AddressStatus `json:"result"`
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

func GetWithdrawTransactionState(walletID int64, orderID string) (response *GetWithdrawTransactionStateResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/sender/transactions/%s", walletID, orderID)

	resp, err := makeRequest(walletID, "GET", uri, nil, nil)
	if err != nil {
		return
	}

	response = &GetWithdrawTransactionStateResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return
	}

	logs.Debug("GetWithdrawTransactionState() => ", response)
	return
}

func GetWithdrawalWalletBalance(walletID int64) (response *GetWithdrawalWalletBalanceResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/sender/balance", walletID)

	resp, err := makeRequest(walletID, "GET", uri, nil, nil)
	if err != nil {
		return
	}

	response = &GetWithdrawalWalletBalanceResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return
	}

	logs.Debug("GetWithdrawalWalletBalance() => ", response)
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

func GetNotificationByID(walletID int64, request *GetNotificationsByIDRequest) (response *GetNotificationsResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/notifications/get_by_id", walletID)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := makeRequest(walletID, "POST", uri, nil, jsonRequest)
	if err != nil {
		return
	}

	response = &GetNotificationsResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetNotificationByID() => ", response)
	return
}

func GetTransactionHistory(walletID int64, fromTime int64, toTime int64, startIndex int, requestNumber int, state int) (response *GetTransactionHistoryResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/transactions", walletID)

	params := []string{}
	params = append(params, fmt.Sprintf("from_time=%d", fromTime))
	params = append(params, fmt.Sprintf("to_time=%d", toTime))
	params = append(params, fmt.Sprintf("start_index=%d", startIndex))
	params = append(params, fmt.Sprintf("request_number=%d", requestNumber))
	params = append(params, fmt.Sprintf("state=%d", state))
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetTransactionHistoryResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetTransactionHistory() => ", response)
	return
}

func GetWalletBlockInfo(walletID int64) (response *GetWalletBlockInfoResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/blocks", walletID)

	params := []string{}
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetWalletBlockInfoResponse{}
	err = json.Unmarshal(resp, response)

	logs.Error("GetWalletBlockInfo() => ", response)
	return
}

func GetInvalidDepositAddresses(walletID int64) (response *GetInvalidDepositAddressesResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/addresses/invalid-deposit", walletID)

	params := []string{}
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetInvalidDepositAddressesResponse{}
	err = json.Unmarshal(resp, response)

	logs.Error("GetInvalidDepositAddresses() => ", response)
	return
}

func GetWalletInfo(walletID int64) (response *GetWalletInfoResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/info", walletID)

	params := []string{}
	resp, err := makeRequest(walletID, "GET", uri, params, nil)
	if err != nil {
		return
	}

	response = &GetWalletInfoResponse{}
	err = json.Unmarshal(resp, response)

	logs.Error("GetWalletInfo() => ", response)
	return
}

func VerifyAddresses(walletID int64, request *VerifyAddressesRequest) (response *VerifyAddressesResponse, err error) {
	uri := fmt.Sprintf("/v1/sofa/wallets/%d/addresses/verify", walletID)

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := makeRequest(walletID, "POST", uri, nil, jsonRequest)
	if err != nil {
		return
	}

	response = &VerifyAddressesResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return
	}

	logs.Debug("VerifyAddresses() => ", response)
	return
}
