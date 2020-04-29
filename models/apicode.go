// Copyright (c) 2018-2020 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type APICode struct {
	APICodeID int64  `orm:"pk;auto;unique;column(api_code_id)" json:"api_code_id"`
	APICode   string `orm:"unique;column(api_code)" json:"api_code"`
	ApiSecret string `orm:"unique;column(api_secret)" json:"api_secret"`
	WalletID  int64  `orm:"unique;column(wallet_id)" json:"wallet_id"`
}

func (m *APICode) TableName() string {
	return "apicode"
}

func (m *APICode) TableEngine() string {
	return "INNODB"
}

func (m *APICode) TableNameWithPrefix() string {
	return GetMockDatabasePrefix() + m.TableName()
}

func SetAPICode(apiCodeObj *APICode) (err error) {
	o := orm.NewOrm()

	existedAPICodeObj, err := GetWalletAPICode(apiCodeObj.WalletID)
	if err != nil {
		apiCodeObj.APICodeID, err = o.Insert(apiCodeObj)
		if err != nil {
			logs.Error("Failed to insert API token =>", err)
			return
		}
	} else {
		apiCodeObj.APICodeID = existedAPICodeObj.APICodeID
		_, err = o.Update(apiCodeObj, "api_secret", "api_code")
		if err != nil {
			logs.Warning("Failed to update API secret =>", err)
			return
		}
	}
	logs.Info("Succeeded to set API token =>", apiCodeObj)
	return
}

func GetWalletAPICode(walletID int64) (apiCodeObj *APICode, err error) {
	o := orm.NewOrm()

	apiCodeObj = &APICode{}
	err = o.QueryTable(apiCodeObj.TableNameWithPrefix()).
		Filter("wallet_id", walletID).
		One(apiCodeObj)
	return
}
