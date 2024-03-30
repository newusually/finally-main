package account

import (
	"finally-main/okx-go/client"
	"finally-main/okx-go/consts"
)

// AccountAPI 结构体
type AccountAPI struct {
	*client.Client
}

// 创建新的AccountAPI实例
func NewAccountAPI(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *AccountAPI {
	return &AccountAPI{
		Client: client.NewClient(apiKey, apiSecretKey, passphrase, useServerTime, flag),
	}
}

// 获取仓位风险
func (a *AccountAPI) GetPositionRisk(instType string) (map[string]interface{}, error) {
	params := make(map[string]string)
	if instType != "" {
		params["instType"] = instType
	}
	return a.Request(consts.GET, consts.POSITION_RISK, params)
}

// 获取账户余额
func (a *AccountAPI) GetAccount(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	if ccy != "" {
		params["ccy"] = ccy
	}
	return a.Request(consts.GET, consts.ACCOUNT_INFO, params)
}

// 获取仓位
func (a *AccountAPI) GetPositions(instType, instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["instId"] = instId
	return a.Request(consts.GET, consts.POSITION_INFO, params)
}

// 获取账单明细（最近7天）
func (a *AccountAPI) GetBillsDetail(instType, ccy, mgnMode, ctType, billType, subType, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["ccy"] = ccy
	params["mgnMode"] = mgnMode
	params["ctType"] = ctType
	params["type"] = billType
	params["subType"] = subType
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return a.Request(consts.GET, consts.BILLS_DETAIL, params)
}

// 获取账单明细（最近3个月）
func (a *AccountAPI) GetBillsDetails(instType, ccy, mgnMode, ctType, billType, subType, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["ccy"] = ccy
	params["mgnMode"] = mgnMode
	params["ctType"] = ctType
	params["type"] = billType
	params["subType"] = subType
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return a.Request(consts.GET, consts.BILLS_ARCHIVE, params)
}

// 获取账户配置
func (a *AccountAPI) GetAccountConfig() (map[string]interface{}, error) {
	return a.Request(consts.GET, consts.ACCOUNT_CONFIG, nil)
}

// 获取持仓模式
func (a *AccountAPI) GetPositionMode(posMode string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["posMode"] = posMode
	return a.Request(consts.POST, consts.POSITION_MODE, params)
}

// 设置杠杆
func (a *AccountAPI) SetLeverage(lever, mgnMode, instId, ccy, posSide string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["lever"] = lever
	params["mgnMode"] = mgnMode
	params["instId"] = instId
	params["ccy"] = ccy
	params["posSide"] = posSide
	return a.Request(consts.POST, consts.SET_LEVERAGE, params)
}

// 获取合约的最大交易数量
func (a *AccountAPI) GetMaximumTradeSize(instId, tdMode, ccy, px string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["tdMode"] = tdMode
	params["ccy"] = ccy
	params["px"] = px
	return a.Request(consts.GET, consts.MAX_TRADE_SIZE, params)
}

// 获取最大可交易数量
func (a *AccountAPI) GetMaxAvailSize(instId, tdMode, ccy, reduceOnly string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["tdMode"] = tdMode
	params["ccy"] = ccy
	params["reduceOnly"] = reduceOnly
	return a.Request(consts.GET, consts.MAX_AVAIL_SIZE, params)
}

// 增加/减少保证金
func (a *AccountAPI) AdjustmentMargin(instId, posSide, marginType, amt string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["posSide"] = posSide
	params["type"] = marginType
	params["amt"] = amt
	return a.Request(consts.POST, consts.ADJUSTMENT_MARGIN, params)
}

// 获取杠杆
func (a *AccountAPI) GetLeverage(instId, mgnMode string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["mgnMode"] = mgnMode
	return a.Request(consts.GET, consts.GET_LEVERAGE, params)
}

// 获取最大贷款额度
func (a *AccountAPI) GetMaxLoan(instId, mgnMode, mgnCcy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["mgnMode"] = mgnMode
	params["mgnCcy"] = mgnCcy
	return a.Request(consts.GET, consts.MAX_LOAN, params)
}

// 获取费率
func (a *AccountAPI) GetFeeRates(instType, instId, uly, category string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["instId"] = instId
	params["uly"] = uly
	params["category"] = category
	return a.Request(consts.GET, consts.FEE_RATES, params)
}

// 获取计息记录
func (a *AccountAPI) GetInterestAccrued(instId, ccy, mgnMode, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["ccy"] = ccy
	params["mgnMode"] = mgnMode
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return a.Request(consts.GET, consts.INTEREST_ACCRUED, params)
}

// 获取利率
func (a *AccountAPI) GetInterestRate(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	return a.Request(consts.GET, consts.INTEREST_RATE, params)
}

// 设置希腊字母（PA/BS）
func (a *AccountAPI) SetGreeks(greeksType string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["greeksType"] = greeksType
	return a.Request(consts.POST, consts.SET_GREEKS, params)
}

// 获取最大提币额度
func (a *AccountAPI) GetMaxWithdrawal(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	return a.Request(consts.GET, consts.MAX_WITHDRAWAL, params)
}
