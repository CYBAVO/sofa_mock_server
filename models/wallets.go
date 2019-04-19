// Copyright (c) 2018-2019 The Cybavo developers
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

type DepositWalletAddress struct {
	AddressID    int64  `orm:"pk;auto;unique;column(address_id)" json:"address_id"`
	Address      string `orm:"unique;column(address)" json:"address"`
	Currency     int64  `orm:"column(currency)" json:"currency"`
	TokenAddress string `orm:"column(token_address)" json:"token_address"`
	WalletID     int64  `orm:"column(wallet_id)" json:"wallet_id"`
}

func (m *DepositWalletAddress) TableName() string {
	return "walletaddress"
}

func (m *DepositWalletAddress) TableEngine() string {
	return "INNODB"
}

func (m *DepositWalletAddress) TableNameWithPrefix() string {
	return GetMockDatabasePrefix() + m.TableName()
}

func AddNewWalletAddresses(addresses []DepositWalletAddress) (succeededCount int64, err error) {
	o := orm.NewOrm()

	succeededCount, err = o.InsertMulti(len(addresses), addresses)
	if err != nil {
		logs.Error("Failed to add addresses to wallet =>", err)
		return
	}
	logs.Debug("Succeeded to add addresses")
	return
}
