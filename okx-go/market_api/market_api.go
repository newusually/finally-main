package market_api

import (
	"finally-main/okx-go/client"
	"finally-main/okx-go/consts"
)

// MarketAPI 结构体
type MarketAPI struct {
	*client.Client
}

// 创建新的MarketAPI实例
func NewMarketAPI(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *MarketAPI {
	return &MarketAPI{
		Client: client.NewClient(apiKey, apiSecretKey, passphrase, useServerTime, flag),
	}
}

// 获取行情
func (m *MarketAPI) GetTickers(instType, uly string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	if uly != "" {
		params["uly"] = uly
	}
	return m.Request(consts.GET, consts.TICKERS_INFO, params)
}

// 获取单个行情
func (m *MarketAPI) GetTicker(instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	return m.Request(consts.GET, consts.TICKER_INFO, params)
}

// 获取指数行情
func (m *MarketAPI) GetIndexTicker(quoteCcy, instId string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["quoteCcy"] = quoteCcy
	params["instId"] = instId
	return m.Request(consts.GET, consts.INDEX_TICKERS, params)
}

// 获取订单簿
func (m *MarketAPI) GetOrderbook(instId, sz string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["sz"] = sz
	return m.Request(consts.GET, consts.ORDER_BOOKS, params)
}

// 获取K线数据
func (m *MarketAPI) GetCandlesticks(instId, after, before, bar, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["after"] = after
	params["before"] = before
	params["bar"] = bar
	params["limit"] = limit
	return m.Request(consts.GET, consts.MARKET_CANDLES, params)
}

// 获取历史K线数据（仅限顶级币种）
func (m *MarketAPI) GetHistoryCandlesticks(instId, after, before, bar, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["after"] = after
	params["before"] = before
	params["bar"] = bar
	params["limit"] = limit
	return m.Request(consts.GET, consts.HISTORY_CANDLES, params)
}

// 获取指数K线数据
func (m *MarketAPI) GetIndexCandlesticks(instId, after, before, bar, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["after"] = after
	params["before"] = before
	params["bar"] = bar
	params["limit"] = limit
	return m.Request(consts.GET, consts.INDEX_CANDLES, params)
}

// 获取标记价格K线数据
func (m *MarketAPI) GetMarkPriceCandlesticks(instId, after, before, bar, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["after"] = after
	params["before"] = before
	params["bar"] = bar
	params["limit"] = limit
	return m.Request(consts.GET, consts.MARK_PRICE_CANDLES, params)
}

// 获取交易数据
func (m *MarketAPI) GetTrades(instId, limit string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instId"] = instId
	params["limit"] = limit
	return m.Request(consts.GET, consts.MARKET_TRADES, params)
}

// 获取交易量
func (m *MarketAPI) GetVolume() (map[string]interface{}, error) {
	return m.Request(consts.GET, consts.VOLUME, nil)
}

// 获取预言机
func (m *MarketAPI) GetOracle() (map[string]interface{}, error) {
	return m.Request(consts.GET, consts.ORACLE, nil)
}

// 获取层级
func (m *MarketAPI) GetTier(instType, tdMode, uly, instId, ccy, tier string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["instType"] = instType
	params["tdMode"] = tdMode
	params["uly"] = uly
	params["instId"] = instId
	params["ccy"] = ccy
	params["tier"] = tier
	return m.Request(consts.GET, consts.TIER, params)
}
