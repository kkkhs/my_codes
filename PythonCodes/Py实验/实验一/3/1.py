def foo(n):
    if n == 1 or n == 2:
        return 1
    else:
        return foo(n - 1) + foo(n - 2)


n = input("请输入参数n的值，将返回斐波那契数列中小于参数 n 的所有值：")
for i in range(1, int(n) + 2):
    a = foo(i)
    if a < int(n):
        print(a)
    else:
        break
