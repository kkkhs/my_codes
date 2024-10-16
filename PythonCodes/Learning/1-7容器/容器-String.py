# 不可修改原字符串， 只能得到新字符串
str1 = "itheima and itcast"
# str1[2] = '0'
print(str1[0])
print(str1[-1])

print(str1.index("and"))

# 替换 replace
str2 = str1.replace("and", "or")
print(str1)
print(str2)

# 分割split
name = str1.split(" ")
print(name)

# strip方法去除前后内容
str2 = " itheima and itcast "
print(str2)
str3 = str2.strip()
# 默认是去除前后空格
print(str3)

str1 = "12itheima and itcast21"
str2 = str1.strip("12")
print(str2)

# 序列切片操作 语法:序列[起始下标：结束下标：步长]  步长默认为1、不包含结束下标
mylist = [0, 1, 2, 3, 4, 5, 6]
result1 = mylist[1:4]
print(result1)

# 可以省略
result2 = mylist[:]
print(result2)

result3 = mylist[::2]
print(result3)

# 反转序列
result4 = mylist[::-1]
print(result4)

result5 = mylist[::-2]
print(result5)

