// Copyright (c) 2018-2019 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/SOFA_MOCK_SERVER/api"
	"github.com/cybavo/SOFA_MOCK_SERVER/models"
)

type OuterController struct {
	beego.Controller
}

func (c *OuterController) AbortWithError(status int, err error) {
	resp := api.ErrorCodeResponse{
		ErrMsg:  err.Error(),
		ErrCode: status,
	}
	c.Data["json"] = resp
	c.Abort(strconv.Itoa(status))
}

// @Title Set API token
// @router /wallets/:wallet_id/apitoken [post]
func (c *OuterController) SetAPIToken() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid walled ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request api.SetAPICodeRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	apiCodeParams := models.APICode{
		APICode:   request.APICode,
		ApiSecret: request.ApiSecret,
		WalletID:  walletID,
	}
	err = models.SetAPICode(&apiCodeParams)
	if err != nil {
		logs.Error("SetAPICode failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	response := &api.CommonResponse{
		Result: 1,
	}
	c.Data["json"] = response
}

// @Title Create deposit wallet addresses
// @router /wallets/:wallet_id/addresses [post]
func (c *OuterController) CreateDepositWalletAddresses() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var request api.CreateDepositWalletAddressesRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.CreateDepositWalletAddresses(walletID, &request)
	if err != nil {
		logs.Error("CreateDepositWalletAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// var walletAddresses []models.DepositWalletAddress
	// for _, address := range resp.Addresses {
	// 	walletAddresses = append(walletAddresses, models.DepositWalletAddress{
	// 		Address:  address,
	// 		WalletID: walletID,
	// 	})
	// }
	// _, err = models.AddNewWalletAddresses(walletAddresses)
	// if err != nil {
	// 	logs.Error("AddNewWalletAddresses failed", err)
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// }

	c.Data["json"] = resp
}

// @Title Get deposit wallet addresses
// @router /wallets/:wallet_id/addresses [get]
func (c *OuterController) GetDepositWalletAddresses() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	startIndex, _ := c.GetInt("start_index", 0)
	requestNumber, _ := c.GetInt("request_number", 1000)

	resp, err := api.GetDepositWalletAddresses(walletID, startIndex, requestNumber)
	if err != nil {
		logs.Error("GetDepositWalletAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Data["json"] = resp
}

// @Title Get deposit wallet pool address
// @router /wallets/:wallet_id/pooladdress [get]
func (c *OuterController) GetDepositWalletPoolAddresses() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := api.GetDepositWalletPoolAddress(walletID)
	if err != nil {
		logs.Error("GetDepositWalletPoolAddress failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Data["json"] = resp
}

func calcSHA256(data []byte) (calculatedHash []byte, err error) {
	sha := sha256.New()
	_, err = sha.Write(data)
	if err != nil {
		return
	}
	calculatedHash = sha.Sum(nil)
	return
}

// @Title Callback
// @router /wallets/callback [post]
func (c *OuterController) Callback() {
	var request api.CallbackStruct
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	apiCodeObj, err := models.GetWalletAPICode(request.WalletID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	checksum := c.Ctx.Input.Header("X-CHECKSUM")
	payload := string(c.Ctx.Input.RequestBody) + apiCodeObj.ApiSecret
	sha, _ := calcSHA256([]byte(payload))
	checksumVerf := base64.URLEncoding.EncodeToString(sha)

	if checksum != checksumVerf {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad checksum"))
	}

	logs.Debug("Callback => %s\n%#v", c.Ctx.Input.RequestBody, request)

	c.Ctx.WriteString("OK")
}

// @Title Resend Callback
// @router /wallets/:wallet_id/callback/resend [post]
func (c *OuterController) CallbackResend() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request api.CallbackResendRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.ResendCallback(walletID, &request)
	if err != nil {
		logs.Error("ResendCallback failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Withdraw transactions
// @router /wallets/:wallet_id/sender/transactions [post]
func (c *OuterController) WithdrawTransactions() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request api.WithdrawTransactionRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.WithdrawTransactions(walletID, &request)
	if err != nil {
		logs.Error("WithdrawTransactions failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get state of withdrawal transaction
// @router /wallets/:wallet_id/sender/transactions/:order_id [get]
func (c *OuterController) GetWithdrawTransactionState() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderID := c.Ctx.Input.Param(":order_id")
	if orderID == "" {
		logs.Error("Invalid order ID")
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid order id"))
		return
	}

	resp, err := api.GetWithdrawTransactionState(walletID, orderID)
	if err != nil {
		logs.Error("GetWithdrawTransactionState failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get API token status
// @router /wallets/:wallet_id/apisecret [get]
func (c *OuterController) GetTxAPITokenStatus() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := api.GetTxAPITokenStatus(walletID)
	if err != nil {
		logs.Error("GetTxAPITokenStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Query notification history
// @router /wallets/:wallet_id/notifications [get]
func (c *OuterController) GetNotifications() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fromTime, _ := c.GetInt64("from_time", -1)
	toTime, _ := c.GetInt64("to_time", -1)
	notificationType, _ := c.GetInt("type", -1)
	if fromTime == -1 || toTime == -1 || notificationType == -1 {
		logs.Error("Invalid parameters")
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid parameters"))
		return
	}

	resp, err := api.GetNotifications(walletID, fromTime, toTime, notificationType)
	if err != nil {
		logs.Error("GetNotifications failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Query wallet transaction history
// @router /wallets/:wallet_id/transactions [get]
func (c *OuterController) GetTransactionHistory() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fromTime, _ := c.GetInt64("from_time", -1)
	toTime, _ := c.GetInt64("to_time", -1)
	startIndex, _ := c.GetInt("start_index", 0)
	requestNumber, _ := c.GetInt("request_number", 0)
	state, _ := c.GetInt("state", -1)

	if fromTime == -1 || toTime == -1 {
		logs.Error("Invalid parameters")
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid parameters"))
		return
	}

	resp, err := api.GetTransactionHistory(walletID, fromTime, toTime, startIndex, requestNumber, state)
	if err != nil {
		logs.Error("QueryWalletTransactionHistory failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Query wallet block info
// @router /wallets/:wallet_id/blocks [get]
func (c *OuterController) GetWalletBlockInfo() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := api.GetWalletBlockInfo(walletID)
	if err != nil {
		logs.Error("GetBlockInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Query invalid deposit addresses
// @router /wallets/:wallet_id/addresses/invalid-deposit [get]
func (c *OuterController) GetInvalidDepositAddresses() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := api.GetInvalidDepositAddresses(walletID)
	if err != nil {
		logs.Error("GetInvalidDepositAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Query wallet basic info
// @router /wallets/:wallet_id/info [get]
func (c *OuterController) GetWalletInfo() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := api.GetWalletInfo(walletID)
	if err != nil {
		logs.Error("GetBlockInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Verify addresses
// @router /wallets/:wallet_id/addresses/verify [post]
func (c *OuterController) VerifyAddresses() {
	defer c.ServeJSON()

	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid walled ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var request api.VerifyAddressesRequest
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.VerifyAddresses(walletID, &request)
	if err != nil {
		logs.Error("VerifyAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Data["json"] = resp
}
