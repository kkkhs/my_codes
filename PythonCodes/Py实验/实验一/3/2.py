a = []

n = input('输入一个大于 2 的自然数:')
if int(n) <= 2:
    print('请重新输入！')
else:
    for i in range(2, int(n)):
        for j in range(2, i):
            if i % j == 0:
                break
        else:
            a.append(i)
    print(a)
