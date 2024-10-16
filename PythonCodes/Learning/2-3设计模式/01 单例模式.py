from StrTools import str_tools

# 避免重复构建类, 节省内存
s1 = str_tools
s2 = str_tools

# s1, s2 当前是同一个id和地址
print(id(s1))
print(id(s2))
print(s1)
print(s2)