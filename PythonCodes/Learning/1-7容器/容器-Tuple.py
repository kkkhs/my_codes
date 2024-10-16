# 特点：不可以修改（除了其中嵌套的元组）

name1 = ("123")
print(type(name1))
name2 = tuple("123", )
print(type(name2))
# 单个元素的元组要加逗号
name = (123, "abc", "cde", 123)

print(name)
print(type(name))

# 其他方式与list一样

