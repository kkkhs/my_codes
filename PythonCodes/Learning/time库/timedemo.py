import time
import calendar

# 获取当前时间
print(time.time())

# 获取当前时间元组localtime
localtime = time.localtime(time.time())
print("当前时间为：", localtime)

# 获取格式化的时间 asctime
localtime = time.asctime(time.localtime(time.time()))
print("当前时间为：", localtime)

# 格式化日期time.strftime(format[, t])
print(time.strftime("%Y-%M-%d %H:%M:%S", time.localtime()))

# 打印某月日历
cal = calendar.month(2023, 5)
print("以下为5月份日历")
print(cal)
