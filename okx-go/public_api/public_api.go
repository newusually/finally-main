package public_api

import (
	"finally-main/okx-go/client"
	"finally-main/okx-go/consts"
)

// PublicAPI 结构体
type PublicAPI struct {
	*client.Client
}

// 创建新的PublicAPI实例
func NewPublicAPI(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *PublicAPI {
	return &PublicAPI{
		Client: client.NewClient(apiKey, apiSecretKey, passphrase, useServerTime, flag),
	}
}

// 获取合约信息
func (p *PublicAPI) GetInstruments(instType, uly, instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["uly"] = uly
	params["instId"] = instId
	return p.Request(consts.GET, consts.INSTRUMENT_INFO, params)
}

// 获取交割/行权历史
func (p *PublicAPI) GetDeliverHistory(instType, uly, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["uly"] = uly
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return p.Request(consts.GET, consts.DELIVERY_EXERCISE, params)
}

// 获取未平仓合约数量
func (p *PublicAPI) GetOpenInterest(instType, uly, instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["uly"] = uly
	params["instId"] = instId
	return p.Request(consts.GET, consts.OPEN_INTEREST, params)
}

// 获取资金费率
func (p *PublicAPI) GetFundingRate(instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	return p.Request(consts.GET, consts.FUNDING_RATE, params)
}

// 获取资金费率历史
func (p *PublicAPI) FundingRateHistory(instId, after, before, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["after"] = after
	params["before"] = before
	params["limit"] = limit
	return p.Request(consts.GET, consts.FUNDING_RATE_HISTORY, params)
}

// 获取限价
func (p *PublicAPI) GetPriceLimit(instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	return p.Request(consts.GET, consts.PRICE_LIMIT, params)
}

// 获取期权市场数据
func (p *PublicAPI) GetOptSummary(uly, expTime string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["uly"] = uly
	params["expTime"] = expTime
	return p.Request(consts.GET, consts.OPT_SUMMARY, params)
}

// 获取预估交割/行权价格
func (p *PublicAPI) GetEstimatedPrice(instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	return p.Request(consts.GET, consts.ESTIMATED_PRICE, params)
}

// 获取折扣率和免息额度
func (p *PublicAPI) DiscountInterestFreeQuota(ccy string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["ccy"] = ccy
	return p.Request(consts.GET, consts.DISCOUNT_INTETEST_INFO, params)
}

// 获取系统时间
func (p *PublicAPI) GetSystemTime() (map[string]interface{}, error) {
	return p.Request(consts.GET, consts.SYSTEM_TIME, nil)
}

// 获取强平订单
func (p *PublicAPI) GetLiquidationOrders(instType, mgnMode, instId, ccy, uly, alias, state, before, after, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["mgnMode"] = mgnMode
	params["instId"] = instId
	params["ccy"] = ccy
	params["uly"] = uly
	params["alias"] = alias
	params["state"] = state
	params["before"] = before
	params["after"] = after
	params["limit"] = limit
	return p.Request(consts.GET, consts.LIQUIDATION_ORDERS, params)
}

// 获取标记价格
func (p *PublicAPI) GetMarkPrice(instType, uly, instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["uly"] = uly
	params["instId"] = instId
	return p.Request(consts.GET, consts.MARK_PRICE, params)
}

// 获取层级
func (p *PublicAPI) GetTier(instType, tdMode, uly, instId, ccy, tier string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["tdMode"] = tdMode
	params["uly"] = uly
	params["instId"] = instId
	params["ccy"] = ccy
	params["tier"] = tier
	return p.Request(consts.GET, consts.MARK_PRICE, params)
}
