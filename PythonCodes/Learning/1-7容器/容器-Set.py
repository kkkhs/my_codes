# 特点：不重复 无序（不支持下标访问）
name = {"123", 123, "abc", "iiii", 123}
# 空集合的定义
name2 = set()
print(name)
print(name2)

print(type(name))
print(type(name2))

# add添加
name.add("Python")
name.add(123)
print(name)

# remove移除
name.remove(123)
print(name)

# 随机取出元素方法 pop()
print(name.pop())

print(name)

# 清空clear
name.clear()
print(name)

set1 = {1, 2, 3}
set2 = {2, 3, 4}
# 取两个集合的差集
set3 = set1.difference(set2)
print(set3)
set3 = set2.difference(set1)
print(set3)

# 消除2个集合的差集
set1.difference_update(set2)
print(set1)
print(set2)

set1 = {1, 2, 3}
set2 = {2, 3, 4}
# 合并两个集合 union
set3 = set1.union(set2)
print(set3)

# 数量len
print(len(set1))

# 集合的遍历
set1 = {1, 2, 3, 4, 5, 6}
for i in set1:
    print(i)








