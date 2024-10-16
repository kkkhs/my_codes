name = ["asd-asd", "NASDAQ", "asd-asd", 123]
name2 = ["last1", "last2"]

# index操作
print(name.index(123))

# insert插入元素
name.insert(1, "itcast")
print(name)

# append追加单个元素
name.append("last one")
print(name)

# extent 追加一批元素
name.extend(name2)
print(name)

# del(下标) 删除元素
name = ["asd-asd", "NASDAQ", "asd-asd", 123]
del name[2]
print(name)

# .pop（下标）取出元素
a = name.pop(1)
print(a)
print(name)

#  remove(元素)
name = ["asd-asd", "NASDAQ", "asd-asd", 123, 123]
name.remove("asd-asd")
# 只会删除第一个元素
print(name)

# 清空列表
name.clear()
print(name)

# 统计某元素的数量
name = ["asd-asd", "NASDAQ", "asd-asd", 123, 123]
print(name.count(123))
print(name.count("123"))

# 统计全部元素的数量
print(len(name))

# 列表的遍历
# while
index = 0
while index < len(name):
    print(name[index])
    index += 1

# for
for i in name:
    print(i)











