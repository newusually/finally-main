import json

with open(f'../../datas/api_bn.json', 'r', encoding='utf-8') as f:
    obj = json.loads(f.read())

api_key = str(obj['api_key'])
secret_key = str(obj['secret_key'])

import os
import sys
from pprint import pprint

root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
sys.path.append(root + '/python')

import ccxt  # noqa: E402

print('CCXT Version:', ccxt.__version__)

exchange = ccxt.binance({
    'apiKey': api_key,
    'secret': secret_key,
    'options': {
        'defaultType': 'future',
    },
})


def table(values):
    first = values[0]
    keys = list(first.keys()) if isinstance(first, dict) else range(0, len(first))
    widths = [max([len(str(v[k])) for v in values]) for k in keys]
    string = ' | '.join(['{:<' + str(w) + '}' for w in widths])
    return "\n".join([string.format(*[str(v[k]) for k in keys]) for v in values])


markets = exchange.load_markets()

symbol = 'BTC/USDT'  # YOUR SYMBOL HERE
market = exchange.market(symbol)

exchange.verbose = True  # UNCOMMENT THIS AFTER LOADING THE MARKETS FOR DEBUGGING

print('----------------------------------------------------------------------')

print('Fetching your balance:')
response = exchange.fetch_balance()
pprint(response['total'])  # make sure you have enough futures margin...
# pprint(response['info'])  # more details

print('----------------------------------------------------------------------')

# https://binance-docs.github.io/apidocs/futures/en/#position-information-v2-user_data

print('Getting your positions:')
response = exchange.fapiPrivateV2_get_positionrisk()
print(table(response))

print('----------------------------------------------------------------------')

# https://binance-docs.github.io/apidocs/futures/en/#change-position-mode-trade

print('Getting your current position mode (One-way or Hedge Mode):')
response = exchange.fapiPrivate_get_positionside_dual()
if response['dualSidePosition']:
    print('You are in Hedge Mode')
else:
    print('You are in One-way Mode')
