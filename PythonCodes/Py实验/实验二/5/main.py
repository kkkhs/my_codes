import random
n = int(input("n="))
m = int(input("m="))
a_list = []
while len(a_list) != n:  # 只要a长度小于n
    x = random.randint(1, m)
    if x % 2 != 0:  # x不是偶数
        a_list.append(x)  # 将x加入a中
print(tuple(a_list))  # 转换为元组输出
