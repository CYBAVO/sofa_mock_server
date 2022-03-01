// Copyright (c) 2018-2022 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package preinit

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	ConfigurationFile = "conf/mockserver.app.conf"
)

func init() {
	logs.Info("LoadAppConfig %s", ConfigurationFile)
	err := beego.LoadAppConfig("ini", ConfigurationFile)

	if err != nil {
		logs.Error("LoadAppConfig,An error occurred:", err)
		os.Exit(1)
	}

}
