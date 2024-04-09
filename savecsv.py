# -*- coding: gbk -*-
import multiprocessing
import sys

from databuy import Databuy, SendDingding
from mvc import MVC

minute = str(sys.argv[1])
if __name__ == '__main__':

    if (minute == '1m'):
        minute = '1'
    if (minute == '3m'):
        minute = '3'
    if (minute == '5m'):
        minute = '5'
    if (minute == '15m'):
        minute = '15'

    data = Databuy()

    data.getethinfo(minute)