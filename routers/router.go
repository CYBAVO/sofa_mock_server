// Copyright (c) 2018-2022 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/cybavo/SOFA_MOCK_SERVER/controllers"
)

func init() {
	beego.Router("/v1/mock/wallets/:wallet_id/apitoken", &controllers.OuterController{}, "POST:SetAPIToken")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses", &controllers.OuterController{}, "POST:CreateDepositWalletAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses", &controllers.OuterController{}, "GET:GetDepositWalletAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/label", &controllers.OuterController{}, "POST:UpdateDepositWalletAddressLabel")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/get_labels", &controllers.OuterController{}, "POST:GetDepositWalletAddressesLabel")
	beego.Router("/v1/mock/wallets/:wallet_id/receiver/addresses/verify", &controllers.OuterController{}, "POST:VerifyDepositAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/pooladdress", &controllers.OuterController{}, "GET:GetDepositWalletPoolAddress")
	beego.Router("/v1/mock/wallets/:wallet_id/pooladdress/balance", &controllers.OuterController{}, "GET:GetDepositWalletPoolAddressBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/collection/notifications/manual", &controllers.OuterController{}, "POST:ResendDepositCollectionCallbacks")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions", &controllers.OuterController{}, "POST:WithdrawAssets")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id/cancel", &controllers.OuterController{}, "POST:CancelWithdrawTransactions")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id", &controllers.OuterController{}, "GET:GetWithdrawTransactionState")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/:order_id/all", &controllers.OuterController{}, "GET:GetWithdrawTransactionStateAll")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/eventlog", &controllers.OuterController{}, "GET:GetTransactionEventLog")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions", &controllers.OuterController{}, "GET:GetSenderTransactionHistory")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/balance", &controllers.OuterController{}, "GET:GetWithdrawalWalletBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/apisecret", &controllers.OuterController{}, "GET:GetTxAPITokenStatus")
	beego.Router("/v1/mock/wallets/:wallet_id/apisecret/activate", &controllers.OuterController{}, "POST:ActivateAPIToken")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications", &controllers.OuterController{}, "GET:GetNotifications")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications/get_by_id", &controllers.OuterController{}, "POST:GetCallbackBySerial")
	beego.Router("/v1/mock/wallets/:wallet_id/notifications/inspect", &controllers.OuterController{}, "POST:NotificationsInspect")
	beego.Router("/v1/mock/wallets/:wallet_id/receiver/notifications/txid/:txid/:vout_index", &controllers.OuterController{}, "GET:GetDepositCallback")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/notifications/order_id/:order_id", &controllers.OuterController{}, "GET:GetWithdrawalCallback")
	beego.Router("/v1/mock/wallets/:wallet_id/transactions", &controllers.OuterController{}, "GET:GetTransactionHistory")
	beego.Router("/v1/mock/wallets/:wallet_id/blocks", &controllers.OuterController{}, "GET:GetWalletBlockInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/invalid-deposit", &controllers.OuterController{}, "GET:GetInvalidDepositAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/info", &controllers.OuterController{}, "GET:GetWalletInfo")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/verify", &controllers.OuterController{}, "POST:VerifyAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/autofee", &controllers.OuterController{}, "POST:GetAutoFee")
	beego.Router("/v1/mock/wallets/:wallet_id/autofees", &controllers.OuterController{}, "POST:GetAutoFees")
	beego.Router("/v1/mock/wallets/:wallet_id/receiver/balance", &controllers.OuterController{}, "GET:GetDepositWalletBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/vault/balance", &controllers.OuterController{}, "GET:GetVaultWalletBalance")
	beego.Router("/v1/mock/wallets/:wallet_id/addresses/contract_txid", &controllers.OuterController{}, "GET:GetDeployedContractCollectionAddresses")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/transactions/acl", &controllers.OuterController{}, "POST:SetWithdrawalACL")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/notifications/manual", &controllers.OuterController{}, "POST:ResendWithdrawalCallbacks")
	beego.Router("/v1/mock/wallets/:wallet_id/refreshsecret", &controllers.OuterController{}, "POST:RefreshSecret")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/whitelist", &controllers.OuterController{}, "GET:GetSenderWhitelist")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/whitelist", &controllers.OuterController{}, "POST:AddSenderWhitelist")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/whitelist", &controllers.OuterController{}, "DELETE:RemoveSenderWhitelist")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/whitelist/config", &controllers.OuterController{}, "GET:QuerySenderWhitelistConfig")
	beego.Router("/v1/mock/wallets/:wallet_id/sender/whitelist/check", &controllers.OuterController{}, "POST:CheckSenderWhitelist")
	beego.Router("/v1/mock/wallets/:wallet_id/signmessage", &controllers.OuterController{}, "POST:SignMessage")
	beego.Router("/v1/mock/wallets/:wallet_id/contract/read", &controllers.OuterController{}, "GET:CallContractRead")
	beego.Router("/v1/mock/wallets/readonly/walletlist", &controllers.OuterController{}, "GET:GetReadOnlyWalletList")
	beego.Router("/v1/mock/wallets/readonly/walletlist/balances", &controllers.OuterController{}, "GET:GetReadOnlyWalletListBalances")
	beego.Router("/v1/mock/currency/prices", &controllers.OuterController{}, "GET:GetCurrencyPrices")

	beego.Router("/v1/mock/wallets/callback", &controllers.OuterController{}, "POST:Callback")
	beego.Router("/v1/mock/wallets/withdrawal/callback", &controllers.OuterController{}, "POST:WithdrawalCallback")

	beego.Router("/v1/mock/merchant/:merchant_id/apitoken", &controllers.MerchantController{}, "POST:SetAPIToken")
	beego.Router("/v1/mock/merchant/:merchant_id/order", &controllers.MerchantController{}, "POST:RequestPaymentOrder")
	beego.Router("/v1/mock/merchant/:merchant_id/order", &controllers.MerchantController{}, "GET:QueryPaymentOrder")
	beego.Router("/v1/mock/merchant/:merchant_id/order/duration", &controllers.MerchantController{}, "POST:UpdateOrderDuration")
	beego.Router("/v1/mock/merchant/:merchant_id/order", &controllers.MerchantController{}, "DELETE:CancelPaymentOrder")
	beego.Router("/v1/mock/merchant/:merchant_id/apisecret", &controllers.MerchantController{}, "GET:GetMerchantAPITokenStatus")
	beego.Router("/v1/mock/merchant/:merchant_id/apisecret/activate", &controllers.MerchantController{}, "POST:ActivateMerchantAPIToken")
	beego.Router("/v1/mock/merchant/:merchant_id/apisecret/refreshsecret", &controllers.MerchantController{}, "POST:RefreshMerchantSecret")
	beego.Router("/v1/mock/merchant/:merchant_id/notifications/manual", &controllers.MerchantController{}, "POST:ResendFailedMerchantCallbacks")

	beego.Router("/v1/mock/merchant/callback", &controllers.MerchantController{}, "POST:Callback")
}
