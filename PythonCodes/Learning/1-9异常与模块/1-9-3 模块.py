# 模块的导入方法1
import time

print("开始")
time.sleep(1)
print("结束")

# 模块的导入方法2(只用某一函数/类/变量/模块/*)
from time import sleep
print("我好")
sleep(1)
print("你好")

from time import *
sleep(1)
print("结束")

# 模块的导入方法3 as别名
import time as t
print("你好")
t.sleep(1)
print("结束")

from time import sleep as sl
sl(1)



