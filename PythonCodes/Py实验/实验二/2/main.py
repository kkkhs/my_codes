# 输入两个集合
setA = set(input('请输入集合A：'))
setB = set(input('请输入集合B：'))

# 并集
s1 = setA | setB
print(s1)
# 交集
s2 = setA & setB
print(s2)
# 差集
s3 = setA - setB
print(s3)
