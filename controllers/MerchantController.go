// Copyright (c) 2018-2021 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/SOFA_MOCK_SERVER/api"
	"github.com/cybavo/SOFA_MOCK_SERVER/models"
)

type MerchantController struct {
	beego.Controller
}

func (c *MerchantController) getMerchantID() int64 {
	merchantID, err := strconv.ParseInt(c.Ctx.Input.Param(":merchant_id"), 10, 64)
	if err != nil {
		logs.Error("Invalid merchant ID =>", err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
	return merchantID
}

func (c *MerchantController) AbortWithError(status int, err error) {
	resp := api.ErrorCodeResponse{
		ErrMsg:  err.Error(),
		ErrCode: status,
	}
	c.Data["json"] = resp
	c.Abort(strconv.Itoa(status))
}

// @Title Set API token
// @router /merchant/:merchant_id/apitoken [post]
func (c *MerchantController) SetAPIToken() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()

	var request api.SetAPICodeRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	apiCodeParams := models.APICode{
		APICode:   request.APICode,
		ApiSecret: request.ApiSecret,
		WalletID:  merchantID,
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

// @Title Request a payment order
// @router /merchant/:merchant_id/order [post]
func (c *MerchantController) RequestPaymentOrder() {
	defer c.ServeJSON()

	req := api.RequestPaymentOrderRequest{}
	if err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &req); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	request := c.Ctx.Input.RequestBody
	if len(req.RedirectURL) > 0 {
		req.RedirectURL = url.QueryEscape(req.RedirectURL)
	} else {
		request, _ = json.Marshal(req)
	}

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "POST", fmt.Sprintf("/v1/merchant/%d/order", merchantID),
		nil, request)
	if err != nil {
		logs.Error("RequestPaymentOrder failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Query certain payment order
// @router /merchant/:merchant_id/order [get]
func (c *MerchantController) QueryPaymentOrder() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "GET", fmt.Sprintf("/v1/merchant/%d/order", merchantID),
		getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("QueryPaymentOrder failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	res := api.QueryPaymentOrderResponse{}
	if err := json.Unmarshal(resp, &res); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if len(res.RedirectURL) > 0 {
		res.RedirectURL, _ = url.QueryUnescape(res.RedirectURL)
		resp, _ = json.Marshal(res)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Update duration of certain payment order
// @router /merchant/:merchant_id/order/duration [post]
func (c *MerchantController) UpdateOrderDuration() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "POST", fmt.Sprintf("/v1/merchant/%d/order/duration", merchantID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("UpdateOrderDuration failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Update duration of certain payment order
// @router /merchant/:merchant_id/order [delete]
func (c *MerchantController) CancelPaymentOrder() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "DELETE", fmt.Sprintf("/v1/merchant/%d/order", merchantID),
		getQueryString(c.Ctx), nil)
	if err != nil {
		logs.Error("CancelPaymentOrder failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Get merchant API token status
// @router /merchant/:merchant_id/apisecret [get]
func (c *MerchantController) GetMerchantAPITokenStatus() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "GET", fmt.Sprintf("/v1/merchant/%d/apisecret", merchantID),
		nil, nil)
	if err != nil {
		logs.Error("GetMerchantAPITokenStatus failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Activate merchant API token
// @router /merchant/:merchant_id/apisecret/activate [post]
func (c *MerchantController) ActivateMerchantAPIToken() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "POST", fmt.Sprintf("/v1/merchant/%d/apisecret/activate", merchantID),
		nil, nil)
	if err != nil {
		logs.Error("ActivateMerchantAPIToken failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Refresh merchant API code and secret
// @router /merchant/:merchant_id/apisecret/refreshsecret [post]
func (c *MerchantController) RefreshMerchantSecret() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "POST", fmt.Sprintf("/v1/merchant/%d/apisecret/refreshsecret", merchantID),
		nil, c.Ctx.Input.RequestBody)
	if err != nil {
		logs.Error("RefreshMerchantSecret failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Resend failed callback
// @router /merchant/:merchant_id/notifications/manual [post]
func (c *MerchantController) ResendFailedMerchantCallbacks() {
	defer c.ServeJSON()

	merchantID := c.getMerchantID()
	resp, err := api.MakeRequest(merchantID, "POST", fmt.Sprintf("/v1/merchant/%d/notifications/manual", merchantID),
		nil, nil)
	if err != nil {
		logs.Error("ResendFailedMerchantCallbacks failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var m map[string]interface{}
	json.Unmarshal(resp, &m)
	c.Data["json"] = m
}

// @Title Merchant callback
// @router /merchant/callback [post]
func (c *MerchantController) Callback() {
	var cb api.MerchantCallbackStruct
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cb)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	apiCodeObj, err := models.GetWalletAPICode(cb.MerchantID)
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

	logs.Debug("Merchant Callback => %s", c.Ctx.Input.RequestBody)

	if cb.State == api.MerchantOrderStateSuccess {
	} else if cb.State == api.MerchantOrderStateExpired {
	} else if cb.State == api.MerchantOrderStateInsufficient {
	} else if cb.State == api.MerchantOrderStateExcess {
	} else if cb.State == api.MerchantOrderStateCancel {
	}

	// reply 200 OK to confirm the callback has been processed
	c.Ctx.WriteString("OK")
}
