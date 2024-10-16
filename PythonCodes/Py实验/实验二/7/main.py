import re
str1 = input("str1= ")
print(re.findall(r"\b[a-zA-Z]{3}\b", str1))  # 返回string中所有与pattern匹配的全部字符串,返回形式为数组。
