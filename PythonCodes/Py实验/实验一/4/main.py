str1 = input("请输入串1:")
str2 = input("请输入串2:")

length = min(len(str1), len(str2))
f = max(filter(lambda i: str1[len(str1) - i:] == str2[:i], range(length+1)))
print(str1+str2[f:])
