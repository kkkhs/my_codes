import time
# 计算n!
def cal(n):
    if n == 1:
        return 1
    else:
        return n * cal(n - 1)


print(cal(5))

# 尾递归是指递归函数在调用自身时直接传递其状态值


def cal2(n, res=1):
    if n == 1:
        return res
    else:
        return cal(n - 1, n * res)


print(cal(5))

# 斐波那契 普通递归
def fib(n):
    if n < 3:
        return 1
    else:
        return fib(n - 1) + fib(n - 2)

start_time = time.time()
print(fib(30))
end_time = time.time()
print(end_time - start_time)

# 斐波那契 尾递归
def fib_tail(n, res=0, temp=1):
    if n == 0:
        return res
    else:
        return fib_tail(n - 1, temp, res + temp)

start_time = time.time()
print(fib_tail(30))
end_time = time.time()
print(end_time - start_time)
