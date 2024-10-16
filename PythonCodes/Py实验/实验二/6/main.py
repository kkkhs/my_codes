a_list = input("str1= ").split(" ")
a_dict = {}
for ch in a_list:  # 遍历a
    a_dict[ch] = a_dict.get(ch, 0)+1  # 利用字典塞入key值累加value
print(a_dict)
