// Copyright (c) 2018-2019 The Cybavo developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Cybavo and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Cybavo
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Cybavo.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/cybavo/SOFA_MOCK_SERVER/controllers"
)

func init() {
	InitUpdateSRVNameSpace()
}

func InitUpdateSRVNameSpace() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/mock",
				beego.NSInclude(
					&controllers.OuterController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
