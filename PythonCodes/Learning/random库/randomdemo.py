import random

#random() 方法返回一个随机数在半开放区间 [0,1) 范围
print(random.random())

# seed() 方法改变随机数生成器的种子
random.seed(10)
print(random.random())
random.seed(10)
print(random.random())

random.seed("hello", 2)
print(random.random())

random.seed('l')
print(random.random())

# 从 range(start, stop, step) 返回一个随机选择的元素。
print(random.randrange(0, 4, 2))

# 返回随机整数 N 满足 a <= N <= b。
print(random.randint(0, 5))

l = {1, 2, 3, 4, 5}
print(type(l))





