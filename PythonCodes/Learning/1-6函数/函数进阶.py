# 一、多个返回值
def return0():
    return 1, 2


x, y = return0()
print(x)
print(y)


# 二、多种参数使用方法
# 1 位置参数（默认）
# 2 关键字参数（key = value)
def info(name, age, gender):
    print(f"姓名是：{name}，年龄是:{age}, 性别是：{gender} ")


info("john", 18, "男")
info(name="kim", gender="boy", age=19)
info("mike", gender="gilr", age=20)  # 混用


# 3 缺省参数
def info2(name, age, gender='男'):
    print(f"姓名是：{name}，年龄是:{age}, 性别是：{gender} ")


info2("john", 18)
info2("john", 18, '女')


# 4 不定长参数
# 4.1位置传递的不定长参数 *   以元组存在
def info3(*args):
    print(f"类型：{type(args)}, 内容：{args}")


info3(1, 2, 3, "boy", "kim")


# 4.2关键字传递的不定长参数 **  以字典存在
def info4(**kwargs):
    print(f"类型：{type(kwargs)}, 内容：{kwargs}")


info4(a=1, b=2, addr="beij")


# 三、以函数作为参数传递
def func(compute):
    result = compute(1, 2)
    print(type(compute))
    print(result)


def compute(x, y):
    return x + y


func(compute)

# lambda匿名函数  只能临时使用一次
def func(compute):
    result = compute(1, 2)
    print(type(compute))
    print(result)

func(lambda x, y: x * y)