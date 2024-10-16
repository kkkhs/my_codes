import random

n = input('请输入列表长度')
a = []
b = []
for i in range(0, int(n)):
    ran = random.randint(1, 100)
    a.append(ran)

b.append(sum(a)/len(a))
for i in a:
    if i > b[0]:
        b.append(i)

b = tuple(b)
print(b)