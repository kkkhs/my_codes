import random

num = [100, 1000, 10000, 100000, 1000000, 9000000]
for n in num:
    count = 0
    for i in range(n):
        x = random.random()
        y = random.random()
        if x**2 + y**2 <= 1:
            count += 1
    print(f"在{n}个点的试验下，Π的计算值为: {4 * count / n}")


