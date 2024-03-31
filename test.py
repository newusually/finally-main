# -*- coding: utf-8 -*-
import pandas as pd

# 读取文件
df = pd.read_csv("..\\datas\\log\\buylog.txt", header=None)

# 获取最后10行
last_10_lines = df.tail(10)

# 保存到原文件，替换源文件
last_10_lines.to_csv("..\\datas\\log\\buylog.txt", header=False, index=False)
