package funding_api

import (
	"finally-main/okx-go/client"
	"finally-main/okx-go/consts"
)

// FundingAPI 结构体
type FundingAPI struct {
	*client.Client
}

// 创建新的FundingAPI实例
func NewFundingAPI(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *FundingAPI {
	return &FundingAPI{
		Client: client.NewClient(apiKey, apiSecretKey, passphrase, useServerTime, flag),
	}
}

// 获取充值地址
func (f *FundingAPI) GetDepositAddress(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	return f.Request(consts.GET, consts.DEPOSIT_ADDRESS, params)
}

// 获取余额
func (f *FundingAPI) GetBalances(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	return f.Request(consts.GET, consts.GET_BALANCES, params)
}

// 资金划转
func (f *FundingAPI) FundsTransfer(ccy, amt, froms, to, transferType, subAcct, instId, toInstId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["amt"] = amt
	params["from"] = froms
	params["to"] = to
	params["type"] = transferType
	params["subAcct"] = subAcct
	params["instId"] = instId
	params["toInstId"] = toInstId
	return f.Request(consts.POST, consts.FUNDS_TRANSFER, params)
}

// 提币
func (f *FundingAPI) CoinWithdraw(ccy, amt, dest, toAddr, pwd, fee string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["amt"] = amt
	params["dest"] = dest
	params["toAddr"] = toAddr
	params["pwd"] = pwd
	params["fee"] = fee
	return f.Request(consts.POST, consts.WITHDRAWAL_COIN, params)
}

// 获取充值历史
func (f *FundingAPI) GetDepositHistory(ccy, state, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["state"] = state
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return f.Request(consts.GET, consts.DEPOSIT_HISTORY, params)
}

// 获取提币历史
func (f *FundingAPI) GetWithdrawalHistory(ccy, state, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["state"] = state
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return f.Request(consts.GET, consts.WITHDRAWAL_HISTORY, params)
}

// 获取货币信息
func (f *FundingAPI) GetCurrency() (map[string]interface{}, error) {
	return f.Request(consts.GET, consts.CURRENCY_INFO, nil)
}

// 存取款
func (f *FundingAPI) PurchaseRedempt(ccy, amt, side string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["amt"] = amt
	params["side"] = side
	return f.Request(consts.POST, consts.PURCHASE_REDEMPT, params)
}

// 获取账单信息
func (f *FundingAPI) GetBills(ccy, billType, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	params["type"] = billType
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return f.Request(consts.GET, consts.BILLS_INFO, params)
}
