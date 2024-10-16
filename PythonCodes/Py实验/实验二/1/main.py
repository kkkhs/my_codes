import string
from random import choices
from collections import Counter

s = ''.join(choices(string.ascii_letters + string.digits, k=1000))  # 随机生成1000个字符
print("string:", s)


def meth1(s1):
    char_count = dict()  # 初始化一个字典
    for i in s1:  # 遍历
        char_count[i] = char_count.get(i, 0) + 1  # 如果指定键的值不存在时，返回0
    print(char_count)


def meth2(s2):
    cnt = Counter()  # 初始化一个计数器
    for i in s2:  # 遍历
        cnt[i] += 1
    print(cnt)


meth1(s)
meth2(s)

