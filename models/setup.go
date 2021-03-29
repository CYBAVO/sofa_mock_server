// Copyright (c) 2018-2021 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package models

import (
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func GetMockDatabasePrefix() string {
	return beego.AppConfig.DefaultString("db_prefix", "mock_")
}

func RegisterDataBase() {
	logs.Info("Init Database Configuration.")

	adapter := beego.AppConfig.DefaultString("db_adapter", "")

	if adapter == "sqlite3" {
		orm.DefaultTimeLoc = time.UTC
		database := beego.AppConfig.DefaultString("db_database", "")

		dbPath := filepath.Dir(database)
		err := os.MkdirAll(dbPath, 0750)
		if err != nil {
			logs.Error("Failed to mkdir =>", err)
		}
		err = orm.RegisterDataBase("default", "sqlite3", database)

		if err != nil {
			logs.Error("sqlite3 Register Database Fail:", err)
		}
	} else {
		logs.Error("DB Non support type:", adapter)
		os.Exit(1)
	}
	logs.Info("Complete the database init:", adapter)
}

func RegisterModel() {

	orm.RegisterModelWithPrefix(GetMockDatabasePrefix(),
		new(APICode),
	)
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Warning("Failed to RunSyncdb => ", err)
	}

}
