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
	beego.Router("/v1/mock/wallets/:wallet_id/apitoken", &controllers.OuterController{}, "POST:SetAPIToken")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses", &controllers.OuterController{}, "POST:CreateDepositWalletAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses", &controllers.OuterController{}, "GET:GetDepositWalletAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/pooladdress", &controllers.OuterController{}, "GET:GetDepositWalletPoolAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/callback/resend", &controllers.OuterController{}, "POST:CallbackResend")
	beego.Router("/v1/mock/wallets/:wallet_id/withdraw", &controllers.OuterController{}, "POST:WithdrawTransactions")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id", &controllers.OuterController{}, "GET:GetWithdrawTransactionState")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/balance", &controllers.OuterController{}, "GET:GetWithdrawalWalletBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/apisecret", &controllers.OuterController{}, "GET:GetTxAPITokenStatus")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications", &controllers.OuterController{}, "GET:GetNotifications")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications/get_by_id", &controllers.OuterController{}, "POST:GetWalletNotificationsByID")
	beego.Router("/v1/mock/wallets/:wallet_id/transactions", &controllers.OuterController{}, "GET:GetTransactionHistory")
	beego.Router("/v1/mock/wallets/:wallet_id/blocks", &controllers.OuterController{}, "GET:GetWalletBlockInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/invalid-deposit", &controllers.OuterController{}, "GET:GetInvalidDepositAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/info", &controllers.OuterController{}, "GET:GetWalletInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/verify", &controllers.OuterController{}, "POST:VerifyAddresses")

	beego.Router("/v1/mock/wallets/callback", &controllers.OuterController{}, "POST:Callback")
	beego.Router("/v1/mock/wallets/withdrawal/callback", &controllers.OuterController{}, "POST:WithdrawalCallback")
}
