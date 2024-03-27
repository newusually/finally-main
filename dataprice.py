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

                # print (t)                       #ԭʼʱ������
                # print (int(t))                  #�뼶ʱ���
                # print (int(round(t * 1000)))    #���뼶ʱ���
                # print (int(round(t * 1000000))) #΢�뼶ʱ���
                tt = str((int(t * 1000)))
                ttt = str((int(round(t * 1000000))))

                # time.sleep(int(minute)/10)

                # ===��ȡclose����

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

                if response.cookies.get_dict():  # ����cookie��Ч
                    s = r.session()
                    c = r.cookies.RequestsCookieJar()  # ����һ��cookie����
                    c.set('cookie-name', 'cookie-value')  # ����cookie��ֵ
                    s.cookies.update(c)  # ����s��cookie
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

                # print (t)                       #ԭʼʱ������
                # print (int(t))                  #�뼶ʱ���
                # print (int(round(t * 1000)))    #���뼶ʱ���
                # print (int(round(t * 1000000))) #΢�뼶ʱ���
                tt = str((int(t * 1000)))
                ttt = str((int(round(t * 1000000))))

                # time.sleep(int(minute) / 10)

                # ===��ȡclose����

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

                if response.cookies.get_dict():  # ����cookie��Ч
                    s = r.session()
                    c = r.cookies.RequestsCookieJar()  # ����һ��cookie����
                    c.set('cookie-name', 'cookie-value')  # ����cookie��ֵ
                    s.cookies.update(c)  # ����s��cookie
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

                # �������� �任���ݸ�ʽ
                old_df = pd.read_csv(f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')[
                         :-100]
                old_df[['date', 'close', 'open', 'high', 'low', 'vol']].to_csv(
                    f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv', index=0)
                old_df = pd.read_csv(f'../datas/old_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')

                old_df['date'] = pd.to_datetime(old_df['date'])
                # old_df['date'] = pd.to_datetime(df['date'], format='mixed')
                new_df = pd.read_csv(f'../datas/new_data/' + symbol + '/' + symbol + '-' + str(minute) + 'min.csv')

                df = pd.DataFrame()

                # ȷ���ںϲ��¾����ݺ����������ݽ������򣬲�����������
                df = pd.concat([old_df, new_df], axis=0)
                df['date'] = pd.to_datetime(df['date'], format='%Y-%m-%d %H:%M:%S')
                df.sort_values(by=['date'], axis=0, ascending=True, inplace=True)
                df.reset_index(drop=True, inplace=True)  # ��������

                # ת����������
                df['vol'] = df['vol'].astype('float')
                df['open'] = df['open'].astype('float')
                df['close'] = df['close'].astype('float')
                df['high'] = df['high'].astype('float')
                df['low'] = df['low'].astype('float')

                df.drop_duplicates(subset=['date'], keep='first', inplace=True)

                # ��������֮ǰ�ٴ�ȷ�������ǰ������������е�
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
                            "close null!!!--->>>" + "��\n�������ػ��ߣ�Ҳ��һȺʱ�̶Կ�Σ�պͷ��Ŀ����棡����")

                break
            except:
                time.sleep(0.5)
                continue

        return close


def getfulldata(df):
    # ��ȡ������ʷ����
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
    # EMA��MACD
    df['obv'] = ta.OBV(df['close'].values, df['vol'].values)
    df['maobv'] = ta.MA(df['obv'].values, timeperiod=30, matype=0)

    df['TRIX'] = ta.TRIX(np.array(df['close'].values), timeperiod=14)
    df['MATRIX'] = ta.MA(df['TRIX'].values, timeperiod=30, matype=0)
    df['SMA'] = ta.SMA(df['close'].values, timeperiod=25)  # SMA���߼۸�������̼�
    df['slowk'], df['slowd'] = ta.STOCH(df['high'].values, df['low'].values, df['close'].values, 5, 3, 0, 3, 0)

    # DEMA˫�ƶ�ƽ����

    # �����ƶ�ƽ���������������źţ��ϳ���������ʶ�����ƣ��϶���������ѡ��ʱ������������ƽ���߼��۸����ߵ��໥���ã��Ź�ͬ�����������ź�

    df['dema'] = ta.DEMA(df.close, timeperiod=30)

    # EMAָ��ƽ����

    # ������ָ�꣬�乹��ԭ������Ȼ�Լ۸����̼۽�������ƽ���������ݼ����������з����������жϼ۸�δ�����Ƶı䶯����

    df['ema'] = ta.EMA(df.close, timeperiod=30)

    # HT_TRENDLINEϣ������˲ʱ�任

    # ������ָ�꣬�乹��ԭ������Ȼ�Լ۸����̼۽�������ƽ���������ݼ����������з����������жϼ۸�δ�����Ƶı䶯���ơ�

    df['HT_TRENDLINE'] = ta.HT_TRENDLINE(df.close)

    # KAMA������������Ӧ�ƶ�ƽ����

    # ���۸���һ����������ƶ�ʱ�����ڵ��ƶ�ƽ����������ʵģ����۸��ں��̵Ĺ����У������ƶ�ƽ�����Ǻ��ʵġ�

    df['KAMA'] = ta.KAMA(df.close, timeperiod=30)

    # MIDPOINT - ���ڵ�

    df['MIDPOINT'] = ta.MIDPOINT(df.close, timeperiod=14)

    # MIDPRICE - ���ڼ۸�

    df['MIDPRICE'] = ta.MIDPRICE(df.high, df.low, timeperiod=14)

    # SAR������ָ��

    # ͣ���ת�����������߷�ʽ����ʱ����ͣ���λ���Թ۲������㡣 ����ͣ��㣨�ֳ�ת���SAR���Ի��εķ�ʽ�ƶ����ʳ�֮Ϊ������ת��ָ�� ��

    df['SAR'] = ta.SAR(df.high, df.low, acceleration=0, maximum=0)

    # SAREXT ������SAR��չ

    df['SAREXT'] = ta.SAREXT(df.high, df.low, startvalue=0, offsetonreverse=0, accelerationinitlong=0,
                             accelerationlong=0, accelerationmaxlong=0, accelerationinitshort=0,
                             accelerationshort=0,
                             accelerationmaxshort=0)

    # T3����ָ���ƶ�ƽ����

    # TRIX���߲���ʱ���ñ�ָ���Ѷ�ţ���ʱ�䰴�ձ�ָ��Ѷ�Ž��ף������ٷֱȴ�����ʧ�ٷֱȣ������൱�ɹۡ�

    df['T3'] = ta.T3(df.close, timeperiod=5, vfactor=0)

    # TEMA����ָ���ƶ�ƽ���ߣ���T3������

    df['TEMA'] = ta.TEMA(df.close, timeperiod=30)

    # TRIMA - �����ƶ�ƽ����

    df['TRIMA'] = ta.TRIMA(df.close, timeperiod=30)

    # WMA�ƶ���Ȩƽ����

    # ��ÿ�ν����ĳɱ�����ԭ�п�����ĳɱ�������ÿ�ν���������ԭ�п����������֮�ͣ����Լ����Ȩƽ����λ�ɱ����Դ�Ϊ�������㵱�·�������ĳɱ�����ĩ����ĳɱ���һ�ַ���

    df['WMA'] = ta.WMA(df.close, timeperiod=30)

    # 1.2 Momentum Indicators(����ָ��)
    # ADXƽ������ָ��

    df['ADX'] = ta.ADX(df.high, df.low, df.close, timeperiod=14)

    # ADXRƽ������ָ��������ָ��

    # ʹ��ADXRָ�꣬�ж�ADX���ơ�

    df['ADXR'] = ta.ADXR(df.high, df.low, df.close, timeperiod=14)

    # APO - ���Լ۸�����

    df['APO'] = ta.APO(df.close, fastperiod=12, slowperiod=26, matype=0)

    # AROON��¡ָ��

    # ͨ�������Լ۸�ﵽ�������ֵ�����ֵ�������������ڼ�������¡ָ�������Ԥ��۸����Ƶ��������򣨻��߷��������������������ƣ��ı仯��

    df['aroondown'], df['aroonup'] = ta.AROON(df.high, df.low, timeperiod=14)

    # AROONOSC��¡��

    df['AROONOSC'] = ta.AROONOSC(df.high, df.low, timeperiod=14)

    # BOP����ָ��

    df['BOP'] = ta.BOP(df.open, df.high, df.low, df.close)

    # CCI˳��ָ��

    # �����ɼ��Ƿ��ѳ�����̬�ֲ���Χ

    df['CCI'] = ta.CCI(df.high, df.low, df.close, timeperiod=14)

    # CMOǮ�¶����ڶ�ָ��

    # ���㹫ʽ�ķ����в��������պ��µ��յ����ݡ�

    df['CMO'] = ta.CMO(df.close, timeperiod=14)

    # DX����ָ�������ָ��

    # ������Ʊ�۸����ǵ�����������˫�����������ı仯����������˫���������ı仯�ܼ۸񲨶���Ӱ��������ɾ��⵽ʧ���ѭ�����̣��Ӷ��ṩ�������ж����ݵ�һ�ּ���ָ�ꡣ

    df['DX'] = ta.DX(df.high, df.low, df.close, timeperiod=14)

    # MFI�ʽ�����ָ��

    # ����������ָ�꣬��ӳ�г�����������

    df['MFI'] = ta.MFI(df.high, df.low, df.close, df.vol, timeperiod=14)

    # MINUS_DI��������ֵ(��DX����)

    # ������Ʊ�۸����ǵ�����������˫�����������ı仯����������˫���������ı仯�ܼ۸񲨶���Ӱ��������ɾ��⵽ʧ���ѭ�����̣��Ӷ��ṩ�������ж����ݵ�һ�ּ���ָ�ꡣ

    df['MINUS_DI'] = ta.MINUS_DI(df.high, df.low, df.close, timeperiod=14)

    # MINUS_DM��������ֵ����DX���ƣ�

    # ������Ʊ�۸����ǵ�����������˫�����������ı仯����������˫���������ı仯�ܼ۸񲨶���Ӱ��������ɾ��⵽ʧ���ѭ�����̣��Ӷ��ṩ�������ж����ݵ�һ�ּ���ָ�ꡣ

    df['MINUS_DM'] = ta.MINUS_DM(df.high, df.low, timeperiod=14)

    # MOM��������ֵ

    # ָ��Ʊ(�򾭼�ָ��)����������������Ӯ�������ţ���д��������Ķ���ЧӦ���������������д����Ÿ��Ķ���ЧӦ��

    df['MOM'] = ta.MOM(df.close, timeperiod=10)

    # PLUS_DI - Plus����ָʾ��

    df['PLUS_DI'] = ta.PLUS_DI(df.high, df.low, df.close, timeperiod=14)

    # PLUS_DM - Plus�����˶�

    df['PLUS_DM'] = ta.PLUS_DM(df.high, df.low, timeperiod=14)

    # PPO�۸��𵴰ٷֱ�ָ������MACD���ƣ�

    df['PPO'] = ta.PPO(df.close, fastperiod=12, slowperiod=26, matype=0)

    # ROC�䶯��ָ��

    # �ɵ���Ĺɼ���һ��������֮ǰ��ĳһ��ɼ۱Ƚϣ���䶯�ٶȵĴ�С,����ӳ��Ʊ�б䶯�Ŀ����̶�

    df['ROC'] = ta.ROC(df.close, timeperiod=10)

    # ROCP - �仯�ʰٷֱ�

    df['ROCP'] = ta.ROCP(df.close, timeperiod=10)

    # ROCR - �仯�ʱ���

    df['ROCR'] = ta.ROCR(df.close, timeperiod=10)

    # ROCR100 - �仯��100����

    df['ROCR100'] = ta.ROCR100(df.close, timeperiod=10)

    # RSI���ǿ��ָ��

    # �Ƚ�һ��ʱ���ڵ�ƽ������������ƽ�����̵����������г�����̵������ʵ�����Ӷ�����δ���г������ơ�

    df['RSI'] = ta.RSI(df.close, timeperiod=14)

    return df


# ����������������
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

        sendtexts = "����ʱ�䣺 " + today + "--->>>" + txt + "��\n" + "��\n�������ػ��ߣ�Ҳ��һȺʱ�̶Կ�Σ�պͷ��Ŀ����棡����"

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
