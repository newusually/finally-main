# -*- coding: gbk -*-
import base64
import datetime
import hashlib
import hmac
import json
import time
import urllib
from datetime import datetime

import numpy as np
import pandas as pd
import requests as r
import talib as ta


class DataPrice:
    def new_symbol_isbuy(minute, symbol):

        for i in range(5):
            try:

                if (minute == '1'):
                    minute = '1m'
                if (minute == '3'):
                    minute = '3m'
                if (minute == '5'):
                    minute = '5m'
                if (minute == '15'):
                    minute = '15m'

                t = time.time()

                # print (t)                       #原始时间数据
                # print (int(t))                  #秒级时间戳
                # print (int(round(t * 1000)))    #毫秒级时间戳
                # print (int(round(t * 1000000))) #微秒级时间戳
                tt = str((int(t * 1000)))
                ttt = str((int(round(t * 1000000))))

                # time.sleep(int(minute)/10)

                # ===获取close数据

                headers = {
                    'authority': 'www.okx.com',
                    'timeout': '1',
                    'x-cdn': 'https://static.okx.com',
                    'devid': '6ec23520-a48b-41f1-b35e-5dea795c61b8',
                    'accept-language': 'zh-CN',
                    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0',
                    'accept': 'application/json',
                    'x-utc': '8',
                    'sec-fetch-dest': 'empty',
                    'app-type': 'web',
                    'sec-fetch-site': 'same-origin',
                    'sec-fetch-mode': 'cors',
                    'referer': 'https://www.okx.com/trade-swap/' + symbol,
                    'cookie': 'locale=zh_CN; defaultLocale=zh_CN; _gcl_au=1.1.1514314517.' + str(
                        tt) + '; _ga=GA1.2.1025788009.' + str(tt) + '; _gid=GA1.2.1077289716.' + str(
                        tt) + '; amp_56bf9d=zTmKdiXyRK-5EUgHM2Qg_x...1fp5jebfd.1fp5jgo7d.2.0.2',
                }

                params = (
                    ('instId', symbol),
                    ('bar', minute),
                    ('after', ''),
                    ('limit', '1500'),
                    ('t', str(ttt)),
                )

                response = r.get('https://www.okx.com/priapi/v5/market/candles', headers=headers, params=params)

                if response.cookies.get_dict():  # 保持cookie有效
                    s = r.session()
                    c = r.cookies.RequestsCookieJar()  # 定义一个cookie对象
                    c.set('cookie-name', 'cookie-value')  # 增加cookie的值
                    s.cookies.update(c)  # 更新s的cookie
                    s.get(url='https://www.okx.com/priapi/v5/market/candles?instId=' + symbol + '&bar=' + str(
                        minute) + '&after=&limit=1500&t=' + tt, headers=headers)
                # print(eval(json.dumps(response.json()))['data'])
                new_df = pd.DataFrame(eval(json.dumps(response.json()))['data'])

                response.close()
                time.sleep(1)
                # print(new_df)
                df = pd.DataFrame()
                df['date'] = new_df[0]
                df['open'] = new_df[1]
                df['high'] = new_df[2]
                df['low'] = new_df[3]
                df['close'] = new_df[4]
                df['vol'] = new_df[5]

                # new_df.columns = ['date', 'open', 'high', 'low', 'close', 'vol', 'p', 'pp']
                datelist = []

                for timestamp in df['date']:
                    date = datetime.fromtimestamp(int(timestamp) / 1000)
                    date = date.strftime('%Y-%m-%d %H:%M:%S')
                    datelist.append(date)
                df['date'] = datelist
                # df['date'] = pd.to_datetime(df['date'], format='mixed')
                df['vol'] = df['vol'].astype('float')
                df['close'] = df['close'].astype('float')
                # print(new_df)
                df.sort_values(by=['date'], axis=0, ascending=True, inplace=True)

                if (minute == '1m'):
                    minute = '1'
                if (minute == '3m'):
                    minute = '3'
                if (minute == '5m'):
                    minute = '5'
                if (minute == '15m'):
                    minute = '15'

                df.to_csv(
                    f'../datas/new_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=False)

                df[['date', 'close', 'open', 'high', 'low', 'vol']].to_csv(
                    f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=0)
                break
            except:
                time.sleep(0.5)
                continue

    def eth_isbuy(minute, symbol):
        close = 0
        for i in range(5):
            try:

                if (minute == '1'):
                    minute = '1m'
                if (minute == '3'):
                    minute = '3m'
                if (minute == '5'):
                    minute = '5m'
                if (minute == '15'):
                    minute = '15m'

                t = time.time()

                # print (t)                       #原始时间数据
                # print (int(t))                  #秒级时间戳
                # print (int(round(t * 1000)))    #毫秒级时间戳
                # print (int(round(t * 1000000))) #微秒级时间戳
                tt = str((int(t * 1000)))
                ttt = str((int(round(t * 1000000))))

                # time.sleep(int(minute) / 10)

                # ===获取close数据

                headers = {
                    'authority': 'www.okx.com',
                    'timeout': '1',
                    'x-cdn': 'https://static.okx.com',
                    'devid': '6ec23520-a48b-41f1-b35e-5dea795c61b8',
                    'accept-language': 'zh-CN',
                    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0',
                    'accept': 'application/json',
                    'x-utc': '8',
                    'sec-fetch-dest': 'empty',
                    'app-type': 'web',
                    'sec-fetch-site': 'same-origin',
                    'sec-fetch-mode': 'cors',
                    'referer': 'https://www.okx.com/trade-swap/' + symbol,
                    'cookie': 'locale=zh_CN; defaultLocale=zh_CN; _gcl_au=1.1.1514314517.' + str(
                        tt) + '; _ga=GA1.2.1025788009.' + str(tt) + '; _gid=GA1.2.1077289716.' + str(
                        tt) + '; amp_56bf9d=zTmKdiXyRK-5EUgHM2Qg_x...1fp5jebfd.1fp5jgo7d.2.0.2',
                }

                params = (
                    ('instId', symbol),
                    ('bar', minute),
                    ('after', ''),
                    ('limit', '1500'),
                    ('t', str(ttt)),
                )

                response = r.get('https://www.okx.com/priapi/v5/market/candles', headers=headers, params=params)

                if response.cookies.get_dict():  # 保持cookie有效
                    s = r.session()
                    c = r.cookies.RequestsCookieJar()  # 定义一个cookie对象
                    c.set('cookie-name', 'cookie-value')  # 增加cookie的值
                    s.cookies.update(c)  # 更新s的cookie
                    s.get(url='https://www.okx.com/priapi/v5/market/candles?instId=' + symbol + '&bar=' + str(
                        minute) + '&after=&limit=1500&t=' + tt, headers=headers)
                # print(eval(json.dumps(response.json()))['data'])
                new_df = pd.DataFrame(eval(json.dumps(response.json()))['data'])

                response.close()
                time.sleep(1)
                # print(new_df)
                df = pd.DataFrame()
                df['date'] = new_df[0]
                df['open'] = new_df[1]
                df['high'] = new_df[2]
                df['low'] = new_df[3]
                df['close'] = new_df[4]
                df['vol'] = new_df[5]

                # new_df.columns = ['date', 'open', 'high', 'low', 'close', 'vol', 'p', 'pp']
                datelist = []

                for timestamp in df['date']:
                    date = datetime.fromtimestamp(int(timestamp) / 1000)
                    date = date.strftime('%Y-%m-%d %H:%M:%S')
                    datelist.append(date)
                df['date'] = datelist
                # df['date'] = pd.to_datetime(df['date'], format='mixed')
                df['vol'] = df['vol'].astype('float')
                df['close'] = df['close'].astype('float')
                # print(new_df)
                df.sort_values(by=['date'], axis=0, ascending=True, inplace=True)

                if (minute == '1m'):
                    minute = '1'
                if (minute == '3m'):
                    minute = '3'
                if (minute == '5m'):
                    minute = '5'
                if (minute == '15m'):
                    minute = '15'

                df.to_csv(
                    f'../datas/new_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=False)

                # 保存数据 变换数据格式
                old_df = pd.read_csv(f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')[
                         :-100]
                old_df[['date', 'close', 'open', 'high', 'low', 'vol']].to_csv(
                    f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=0)
                old_df = pd.read_csv(f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')

                old_df['date'] = pd.to_datetime(old_df['date'])
                # old_df['date'] = pd.to_datetime(df['date'], format='mixed')
                new_df = pd.read_csv(f'../datas/new_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')

                df = pd.DataFrame()

                # 确保在合并新旧数据后立即对数据进行排序，并且重置索引
                df = pd.concat([old_df, new_df], axis=0)
                df['date'] = pd.to_datetime(df['date'], format='%Y-%m-%d %H:%M:%S')
                df.sort_values(by=['date'], axis=0, ascending=True, inplace=True)
                df.reset_index(drop=True, inplace=True)  # 重置索引

                # 转换数据类型
                df['vol'] = df['vol'].astype('float')
                df['open'] = df['open'].astype('float')
                df['close'] = df['close'].astype('float')
                df['high'] = df['high'].astype('float')
                df['low'] = df['low'].astype('float')

                df.drop_duplicates(subset=['date'], keep='first', inplace=True)

                # 保存数据之前再次确保数据是按日期正序排列的
                df = df[['date', 'close', 'open', 'high', 'low', 'vol']]
                df.to_csv(f'../datas/old_data/{symbol}/{symbol}-{minute}min.csv', index=False)
                # df = getfulldata(df).dropna()
                df.to_csv(
                    f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=0)
                if symbol == 'ETH-USDT-SWAP':
                    # date=df['date'].values[-1]
                    close = df["close"].values[-1]
                    # print(symbol,minute,close)
                    # print("================================================")
                    if len(close) < 3:
                        SendDingding.sender(
                            "close null!!!--->>>" + "，\n我们是守护者，也是一群时刻对抗危险和疯狂的可怜虫！！！")

                break
            except:
                time.sleep(0.5)
                continue

        return close


def getfulldata(df):
    # 获取参数历史数据
    df['close5'] = ta.EMA(np.array(df['close'].values), timeperiod=5)
    df['close10'] = ta.EMA(np.array(df['close'].values), timeperiod=10)
    df['close20'] = ta.EMA(np.array(df['close'].values), timeperiod=200)
    df['close30'] = ta.EMA(np.array(df['close'].values), timeperiod=30)
    df['close60'] = ta.EMA(np.array(df['close'].values), timeperiod=60)
    df['close120'] = ta.EMA(np.array(df['close'].values), timeperiod=120)
    df['close250'] = ta.EMA(np.array(df['close'].values), timeperiod=250)

    df["MACD_macd"], df["MACD_macdsignal"], df["MACD_macdhist"] = ta.MACD(df['close'].values, fastperiod=12,
                                                                          slowperiod=26, signalperiod=60)
    df['macd'] = 2 * (df["MACD_macd"] - df["MACD_macdsignal"])

    df["MA"] = ta.MA(df['close'].values, timeperiod=30, matype=0)
    # EMA和MACD
    df['obv'] = ta.OBV(df['close'].values, df['vol'].values)
    df['maobv'] = ta.MA(df['obv'].values, timeperiod=30, matype=0)

    df['TRIX'] = ta.TRIX(np.array(df['close'].values), timeperiod=14)
    df['MATRIX'] = ta.MA(df['TRIX'].values, timeperiod=30, matype=0)
    df['SMA'] = ta.SMA(df['close'].values, timeperiod=25)  # SMA均线价格计算收盘价
    df['slowk'], df['slowd'] = ta.STOCH(df['high'].values, df['low'].values, df['close'].values, 5, 3, 0, 3, 0)

    # DEMA双移动平均线

    # 两条移动平均线来产生趋势信号，较长期者用来识别趋势，较短期者用来选择时机。正是两条平均线及价格三者的相互作用，才共同产生了趋势信号

    df['dema'] = ta.DEMA(df.close, timeperiod=30)

    # EMA指数平均数

    # 趋向类指标，其构造原理是仍然对价格收盘价进行算术平均，并根据计算结果来进行分析，用于判断价格未来走势的变动趋势

    df['ema'] = ta.EMA(df.close, timeperiod=30)

    # HT_TRENDLINE希尔伯特瞬时变换

    # 趋向类指标，其构造原理是仍然对价格收盘价进行算术平均，并根据计算结果来进行分析，用于判断价格未来走势的变动趋势。

    df['HT_TRENDLINE'] = ta.HT_TRENDLINE(df.close)

    # KAMA考夫曼的自适应移动平均线

    # 当价格沿一个方向快速移动时，短期的移动平均线是最合适的；当价格在横盘的过程中，长期移动平均线是合适的。

    df['KAMA'] = ta.KAMA(df.close, timeperiod=30)

    # MIDPOINT - 中期点

    df['MIDPOINT'] = ta.MIDPOINT(df.close, timeperiod=14)

    # MIDPRICE - 中期价格

    df['MIDPRICE'] = ta.MIDPRICE(df.high, df.low, timeperiod=14)

    # SAR抛物线指标

    # 停损点转向，利用抛物线方式，随时调整停损点位置以观察买卖点。 由于停损点（又称转向点SAR）以弧形的方式移动，故称之为抛物线转向指标 。

    df['SAR'] = ta.SAR(df.high, df.low, acceleration=0, maximum=0)

    # SAREXT 抛物面SAR扩展

    df['SAREXT'] = ta.SAREXT(df.high, df.low, startvalue=0, offsetonreverse=0, accelerationinitlong=0,
                             accelerationlong=0, accelerationmaxlong=0, accelerationinitshort=0,
                             accelerationshort=0,
                             accelerationmaxshort=0)

    # T3三重指数移动平均线

    # TRIX长线操作时采用本指标的讯号，长时间按照本指标讯号交易，获利百分比大于损失百分比，利润相当可观。

    df['T3'] = ta.T3(df.close, timeperiod=5, vfactor=0)

    # TEMA三重指数移动平均线（与T3无区别）

    df['TEMA'] = ta.TEMA(df.close, timeperiod=30)

    # TRIMA - 三角移动平均线

    df['TRIMA'] = ta.TRIMA(df.close, timeperiod=30)

    # WMA移动加权平均法

    # 以每次进货的成本加上原有库存存货的成本，除以每次进货数量与原有库存存货的数量之和，据以计算加权平均单位成本，以此为基础计算当月发出存货的成本和期末存货的成本的一种方法

    df['WMA'] = ta.WMA(df.close, timeperiod=30)

    # 1.2 Momentum Indicators(动量指标)
    # ADX平均趋向指数

    df['ADX'] = ta.ADX(df.high, df.low, df.close, timeperiod=14)

    # ADXR平均趋向指数的趋向指数

    # 使用ADXR指标，判断ADX趋势。

    df['ADXR'] = ta.ADXR(df.high, df.low, df.close, timeperiod=14)

    # APO - 绝对价格振荡器

    df['APO'] = ta.APO(df.close, fastperiod=12, slowperiod=26, matype=0)

    # AROON阿隆指标

    # 通过计算自价格达到近期最高值和最低值以来所经过的期间数，阿隆指标帮助你预测价格趋势到趋势区域（或者反过来，从趋势区域到趋势）的变化。

    df['aroondown'], df['aroonup'] = ta.AROON(df.high, df.low, timeperiod=14)

    # AROONOSC阿隆振荡

    df['AROONOSC'] = ta.AROONOSC(df.high, df.low, timeperiod=14)

    # BOP均势指标

    df['BOP'] = ta.BOP(df.open, df.high, df.low, df.close)

    # CCI顺势指标

    # 测量股价是否已超出常态分布范围

    df['CCI'] = ta.CCI(df.high, df.low, df.close, timeperiod=14)

    # CMO钱德动量摆动指标

    # 计算公式的分子中采用上涨日和下跌日的数据。

    df['CMO'] = ta.CMO(df.close, timeperiod=14)

    # DX动向指标或趋向指标

    # 分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

    df['DX'] = ta.DX(df.high, df.low, df.close, timeperiod=14)

    # MFI资金流量指标

    # 属于量价类指标，反映市场的运行趋势

    df['MFI'] = ta.MFI(df.high, df.low, df.close, df.vol, timeperiod=14)

    # MINUS_DI下升动向值(与DX相似)

    # 分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

    df['MINUS_DI'] = ta.MINUS_DI(df.high, df.low, df.close, timeperiod=14)

    # MINUS_DM上升动向值（与DX相似）

    # 分析股票价格在涨跌过程中买卖双方力量均衡点的变化情况，即多空双方的力量的变化受价格波动的影响而发生由均衡到失衡的循环过程，从而提供对趋势判断依据的一种技术指标。

    df['MINUS_DM'] = ta.MINUS_DM(df.high, df.low, timeperiod=14)

    # MOM上升动向值

    # 指股票(或经济指数)持续增长的能力。赢家组合在牛市中存在着正的动量效应，输家组合在熊市中存在着负的动量效应。

    df['MOM'] = ta.MOM(df.close, timeperiod=10)

    # PLUS_DI - Plus方向指示器

    df['PLUS_DI'] = ta.PLUS_DI(df.high, df.low, df.close, timeperiod=14)

    # PLUS_DM - Plus定向运动

    df['PLUS_DM'] = ta.PLUS_DM(df.high, df.low, timeperiod=14)

    # PPO价格震荡百分比指数（与MACD相似）

    df['PPO'] = ta.PPO(df.close, fastperiod=12, slowperiod=26, matype=0)

    # ROC变动率指标

    # 由当天的股价与一定的天数之前的某一天股价比较，其变动速度的大小,来反映股票市变动的快慢程度

    df['ROC'] = ta.ROC(df.close, timeperiod=10)

    # ROCP - 变化率百分比

    df['ROCP'] = ta.ROCP(df.close, timeperiod=10)

    # ROCR - 变化率比率

    df['ROCR'] = ta.ROCR(df.close, timeperiod=10)

    # ROCR100 - 变化率100比例

    df['ROCR100'] = ta.ROCR100(df.close, timeperiod=10)

    # RSI相对强弱指数

    # 比较一段时期内的平均收盘涨数和平均收盘跌数来分析市场买沽盘的意向和实力，从而作出未来市场的走势。

    df['RSI'] = ta.RSI(df.close, timeperiod=14)

    return df


# 发钉钉的类先声明
class SendDingding:
    def sender(txt):
        headers = {
            'Content-Type': 'application/json'
        }
        timestamp = str(round(time.time() * 1000))
        secret = "SEC050a3b2c9e5d8d0c777bbdd61270676a8bdad3608b36a086d70e95b712ad2db0"
        secret_enc = secret.encode('utf-8')
        string_to_sign = '{}\n{}'.format(timestamp, secret)
        string_to_sign_enc = string_to_sign.encode('utf-8')
        hmac_code = hmac.new(secret_enc, string_to_sign_enc, digestmod=hashlib.sha256).digest()
        sign = urllib.parse.quote_plus(base64.b64encode(hmac_code))
        today = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())

        sendtexts = "本地时间： " + today + "--->>>" + txt + "，\n" + "，\n我们是守护者，也是一群时刻对抗危险和疯狂的可怜虫！！！"

        params = {
            'sign': sign,

            'timestamp': timestamp
        }
        text_data = {
            "msgtype": "text",
            "text": {
                "content": sendtexts
            }
        }

        roboturl = 'https://oapi.dingtalk.com/robot/send?access_token=f8195c9e4ad6da4427d67e80dffed5d07ecaca1d1e79462fb5c0a9c6b12e90f2'
        r = requests.post(roboturl, data=json.dumps(text_data), params=params, headers=headers)
