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

    symbollist = MVC.getsymbollist()

    if len(symbollist) < 10:
        SendDingding.sender("symbol null!!!--->>>" + "，\n我们是守护者，也是一群时刻对抗危险和疯狂的可怜虫！！！")
    else:

        p1 = multiprocessing.Process(target=data.getbuyinfo, args=[symbollist[:50], minute])
        p2 = multiprocessing.Process(target=data.getbuyinfo, args=[symbollist[50:100], minute])
        p3 = multiprocessing.Process(target=data.getbuyinfo, args=[symbollist[100:150], minute])
        p4 = multiprocessing.Process(target=data.getbuyinfo, args=[symbollist[150:], minute])

        p1.start()
        p2.start()
        p3.start()
        p4.start()

        p1.join()
        p2.join()
        p3.join()
        p4.join()
