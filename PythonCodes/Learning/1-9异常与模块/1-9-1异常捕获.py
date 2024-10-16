# 基本捕获异常
try:
    f = open("D:/abc.txt", "r")
except:
    print("出现异常")
    f = open("D:/abc.txt", "w")

# 捕获指定的异常
try:
    print(name)
except NameError as e:
    print("出现了变量未定义的异常")
    print(e)

# 捕获多个异常
try: 
    1/0
    name
except (NameError, ZeroDivisionError) as e:
    print("多个异常")
    print(e)

# 捕获所有异常
try:
    a = 1
except Exception as e:
    print(e)
else:   # 没有出现异常
    print("没有出现异常")
finally:  # 必须执行
    print("这句话必须打印")