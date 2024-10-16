import random


def random_string():  # 随机生成字符串函数
    s = [chr(i) for i in range(33, 126)]
    return ''.join(random.sample(s, random.randint(1, 80)))


str1 = random_string()
str2 = random_string()
str1 = str1 + str2
print(len(str1))
with open(r'E:\Codes\PythonCodes\Py实验\实验三\5.txt', 'w') as fp:
    fp.write(str1)
with open(r'E:\Codes\PythonCodes\Py实验\实验三\5.txt', 'r') as fp:
    lines = fp.read()
    print(len(lines))
