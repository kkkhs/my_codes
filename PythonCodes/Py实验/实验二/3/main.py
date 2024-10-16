import random
n = int(input("n="))  # 外层
m = int(input("m="))  # 内层
a_list = [[]]*n
for i in range(n):
    a_list[i] = [random.randint(1, 10) for i in range(random.randint(1, m))]

a_list.sort(key=lambda x: len(str(x)), reverse=True)  #  按照长度排列
print(a_list)
