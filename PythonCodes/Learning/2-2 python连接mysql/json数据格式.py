# 数据交互格式 用于组织和封装数据
import json

# 将python转换为json
data = [{"name": "khs", "age": 11}, {"name": "khw", "age": 12}, {"name": "khx", "age": 13}]
json_str = json.dumps(data, ensure_ascii=False)

print(type(json_str))
print(json_str)

d = {"name": "周杰伦", "addr": "台北"}

json_str = json.dumps(d, ensure_ascii=False)
print(type(json_str))
print(json_str)

# 将json转换为python
s = '[{"name": "khs", "age": 11}, {"name": "khw", "age": 12}, {"name": "khx", "age": 13}]'
l = json.loads(s)
print(type(l))
print(l)

d = '{"name": "周杰伦", "addr": "台北"}'
a = json.loads(d)
print(type(a))
print(a)

