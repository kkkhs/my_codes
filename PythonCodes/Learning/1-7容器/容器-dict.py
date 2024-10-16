# 特带：一一对应 key:value

name = {"wlh": 99, "zjl": 88, "ljj": 77}
# 空字典
name1 = {}
name2 = dict()

print(type(name))
print(type(name1))
print(type(name2))

# 索引
print(name["wlh"])
print(name["zjl"])
print(name["ljj"])

# 新增/更新元素
name["khs"] = 100
name["zjl"] = 100
print(name)

# pop删除元素
value = name.pop("khs")
print(value)
print(name)

# 清空元素
name.clear()
print(name)

name = {"wlh": 99, "zjl": 88, "ljj": 77}
# 获取全部key
print(name.keys())

# 遍历字典
for key in name.keys():
    print(name[key])

for i in name:
    print(f"key :{i}")
    print(f"value :{name[i]}")

# 通用排序
print(sorted(name))
print(sorted(name, reverse=True))


