// Copyright (c) 2018-2020 The Cybavo developers
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
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/SOFA_MOCK_SERVER/api"
	"github.com/cybavo/SOFA_MOCK_SERVER/models"
)

type OuterController struct {
	beego.Controller
}

func getQueryString(ctx *context.Context) []string {
	var qs []string
	tokens := strings.Split(ctx.Request.URL.RawQuery, "&")
	for _, token := range tokens {
		qs = append(qs, token)
	}
	return qs
}

var debugPrint = func(ctx *context.Context) {
	var params string
	qs := getQueryString(ctx)
	if qs != nil {
		params = strings.Join(qs, "&")
	}
	logs.Debug(fmt.Sprintf("Recv requst => %s, params: %s, body: %s", ctx.Input.URL(), params, ctx.Input.RequestBody))
}

func init() {
	beego.InsertFilter("/v1/mock/*", beego.BeforeExec, debugPrint)
}

func (c *OuterController) getWalletID() int64 {
	walletID, err := strconv.ParseInt(c.Ctx.Input.Param(":wallet_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid wallet ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
	return walletID
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

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/addresses", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("CreateDepositWalletAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get deposit wallet addresses
// @router /wallets/:wallet_id/addresses [get]
func (c *OuterController) GetDepositWalletAddresses() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/addresses", walletID),
		getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetDepositWalletAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get deposit wallet pool address
// @router /wallets/:wallet_id/pooladdress [get]
func (c *OuterController) GetDepositWalletPoolAddresses() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/pooladdress", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetDepositWalletPoolAddress failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
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

	logs.Debug("Callback => %s", c.Ctx.Input.RequestBody)

	c.Ctx.WriteString("OK")
}

// @Title Withdrawal Callback
// @router /wallets/withdrawal/callback [post]
func (c *OuterController) WithdrawalCallback() {
	// How to verify:
	// 1. Try to find corresponding API secret by request.Requests[0].OrderID
	// 2. Calculate checksum then compare to X-CHECKSUM header (refer to sample code bellow)
	// 3. If these two checksums match and the request is valid in your system,
	//    reply 200, OK otherwise reply 400 to decline the withdrawal

	// sample code to calculate checksum and verify
	// payload := string(c.Ctx.Input.RequestBody) + APISECRET
	// sha, _ := calcSHA256([]byte(payload))
	// checksumVerf := base64.URLEncoding.EncodeToString(sha)
	// checksum := c.Ctx.Input.Header("X-CHECKSUM")
	// if checksum != checksumVerf {
	//   c.AbortWithError(http.StatusBadRequest, errors.New("Bad checksum"))
	// }

	logs.Debug("Withdraw Callback => %s", c.Ctx.Input.RequestBody)

	c.Ctx.WriteString("OK")
}

// @Title Resend Callback
// @router /wallets/:wallet_id/callback/resend [post]
func (c *OuterController) CallbackResend() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/collection/notifications/manual", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("ResendCallback failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Withdraw transactions
// @router /wallets/:wallet_id/withdraw [post]
func (c *OuterController) WithdrawTransactions() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/sender/transactions", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("WithdrawTransactions failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get state of withdrawal transaction
// @router /wallets/:wallet_id/sender/transactions/:order_id [get]
func (c *OuterController) GetWithdrawTransactionState() {
	defer c.ServeJSON()

	orderID := c.Ctx.Input.Param(":order_id")
	if orderID == "" {
		logs.Error("Invalid order ID")
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid order id"))
		return
	}

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/sender/transactions/%s", walletID, orderID),
		nil, nil)
	if err != nil {
		logs.Error("GetWithdrawTransactionState failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get balance of withdrawal wallet
// @router /wallets/:wallet_id/sender/balance [get]
func (c *OuterController) GetWithdrawalWalletBalance() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/sender/balance", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetWithdrawalWalletBalance failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get API token status
// @router /wallets/:wallet_id/apisecret [get]
func (c *OuterController) GetTxAPITokenStatus() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/apisecret", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetTxAPITokenStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query notification history
// @router /wallets/:wallet_id/notifications [get]
func (c *OuterController) GetNotifications() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/notifications", walletID),
		getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetNotifications failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query notification by ID
// @router /wallets/:wallet_id/notifications/get_by_id [post]
func (c *OuterController) GetWalletNotificationsByID() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/notifications/get_by_id", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("GetWalletNotificationsByID failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query wallet transaction history
// @router /wallets/:wallet_id/transactions [get]
func (c *OuterController) GetTransactionHistory() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/transactions", walletID),
		getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("GetTransactionHistory failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query wallet block info
// @router /wallets/:wallet_id/blocks [get]
func (c *OuterController) GetWalletBlockInfo() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/blocks", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetWalletBlockInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query invalid deposit addresses
// @router /wallets/:wallet_id/addresses/invalid-deposit [get]
func (c *OuterController) GetInvalidDepositAddresses() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/addresses/invalid-deposit", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetInvalidDepositAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query wallet basic info
// @router /wallets/:wallet_id/info [get]
func (c *OuterController) GetWalletInfo() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "GET", fmt.Sprintf("/v1/sofa/wallets/%d/info", walletID),
		nil, nil)
	if err != nil {
		logs.Error("GetWalletInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Verify addresses
// @router /wallets/:wallet_id/addresses/verify [post]
func (c *OuterController) VerifyAddresses() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/addresses/verify", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("VerifyAddresses failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query wallet transaction avarage blockchain fee
// @router /wallets/:wallet_id/autofee [post]
func (c *OuterController) GetAutoFee() {
	defer c.ServeJSON()

	walletID := c.getWalletID()
	resp, err := api.MakeRequest(walletID, "POST", fmt.Sprintf("/v1/sofa/wallets/%d/autofee", walletID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("GetAutoFee failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}
