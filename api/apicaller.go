// Copyright (c) 2018-2021 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/SOFA_MOCK_SERVER/models"
)

var baseURL = beego.AppConfig.DefaultString("api_server_url", "")

func buildChecksum(params []string, secret string, time int64, r string) string {
	params = append(params, fmt.Sprintf("t=%d", time))
	params = append(params, fmt.Sprintf("r=%s", r))
	sort.Strings(params)
	params = append(params, fmt.Sprintf("secret=%s", secret))
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(params, "&"))))
}

func MakeRequest(targetID int64, method string, api string, params []string, postBody []byte) ([]byte, error) {
	if targetID < 0 || method == "" || api == "" {
		return nil, errors.New("invalid parameters")
	}

	client := &http.Client{}
	r := RandomString(8)
	if r == "" {
		return nil, errors.New("can't generate random byte string")
	}
	t := time.Now().Unix()
	url := fmt.Sprintf("%s%s?t=%d&r=%s", baseURL, api, t, r)
	if len(params) > 0 {
		url += fmt.Sprintf("&%s", strings.Join(params, "&"))
	}
	var req *http.Request
	var err error
	if postBody == nil || len(postBody) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(postBody))
		params = append(params, string(postBody))
	}
	if err != nil {
		return nil, err
	}

	apiCodeObj, err := models.GetWalletAPICode(targetID)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-CODE", apiCodeObj.APICode)
	req.Header.Set("X-CHECKSUM", buildChecksum(params, apiCodeObj.ApiSecret, t, r))
	if postBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	logs.Debug("Request URL:", url)
	logs.Debug("\tX-CHECKSUM:\t", req.Header.Get("X-CHECKSUM"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		result := &ErrorCodeResponse{}
		_ = json.Unmarshal(body, result)
		msg := fmt.Sprintf("%s, Error: %s", res.Status, result.String())
		return body, errors.New(msg)
	}
	return body, nil
}
