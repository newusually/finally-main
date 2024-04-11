# -*- coding: gbk -*-
from mvc import MVC
from userinfo import User

if __name__ == '__main__':
    api_key, secret_key, passphrase, flag = User.get_userinfo()
    MVC.orderbuy(api_key, secret_key, passphrase, flag, "REN-USDT-SWAP", "15m")
