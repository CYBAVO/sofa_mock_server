// Copyright (c) 2018-2020 The Cybavo developers
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
	beego.Router("/v1/mock/wallets/:wallet_id/pooladdress", &controllers.OuterController{}, "GET:GetDepositWalletPoolAddress")
	beego.Router("/v1/mock/wallets/:wallet_id/pooladdress/balance", &controllers.OuterController{}, "GET:GetDepositWalletPoolAddressBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/callback/resend", &controllers.OuterController{}, "POST:CallbackResend")
	beego.Router("/v1/mock/wallets/:wallet_id/withdraw", &controllers.OuterController{}, "POST:WithdrawAssets")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id/cancel", &controllers.OuterController{}, "POST:CancelWithdrawTransactions")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id", &controllers.OuterController{}, "GET:GetWithdrawTransactionState")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id/all", &controllers.OuterController{}, "GET:GetWithdrawTransactionStateAll")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/balance", &controllers.OuterController{}, "GET:GetWithdrawalWalletBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/apisecret", &controllers.OuterController{}, "GET:GetTxAPITokenStatus")
	beego.Router("/v1/mock/wallets/:wallet_id/apisecret/activate", &controllers.OuterController{}, "POST:ActivateAPIToken")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications", &controllers.OuterController{}, "GET:GetNotifications")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications/get_by_id", &controllers.OuterController{}, "POST:GetCallbackBySerial")
	beego.Router("/v1/mock/wallets/:wallet_id/receiver/notifications/txid/:txid/:vout_index", &controllers.OuterController{}, "GET:GetDepositCallback")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/notifications/order_id/:order_id", &controllers.OuterController{}, "GET:GetWithdrawalCallback")
	beego.Router("/v1/mock/wallets/:wallet_id/transactions", &controllers.OuterController{}, "GET:GetTransactionHistory")
	beego.Router("/v1/mock/wallets/:wallet_id/blocks", &controllers.OuterController{}, "GET:GetWalletBlockInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/invalid-deposit", &controllers.OuterController{}, "GET:GetInvalidDepositAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/info", &controllers.OuterController{}, "GET:GetWalletInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/verify", &controllers.OuterController{}, "POST:VerifyAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/autofee", &controllers.OuterController{}, "POST:GetAutoFee")
	beego.Router("/v1/mock/wallets/:wallet_id/receiver/balance", &controllers.OuterController{}, "GET:GetDepositWalletBalance")

	beego.Router("/v1/mock/wallets/callback", &controllers.OuterController{}, "POST:Callback")
	beego.Router("/v1/mock/wallets/withdrawal/callback", &controllers.OuterController{}, "POST:WithdrawalCallback")
}
