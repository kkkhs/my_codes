import time

f = open("D:/测试.txt", "r", encoding="UTF-8")
print(type(f))

# 读取操作read
print(f.read(10))       # 读取10字节
print(f.read())         # 第二次读取从上一个结尾处开始

# 读取操作readLines 读取全部行封装到列表里
f = open("D:/测试.txt", "r", encoding="UTF-8")
print(f.readlines())
print(type(f.readlines()))

# 读取操作readline
f = open("D:/测试.txt", "r", encoding="UTF-8")
line1 = f.readline()
line2 = f.readline()
line3 = f.readline()
print(line1)
print(line2)
print(line3)

# for循环读取
f = open("D:/测试.txt", "r", encoding="UTF-8")
for line in f:
    print(f"每一行：{line}")

# 文件的关闭close
f.close()
print("文件关闭成功！")

# with open 语法（执行完语句块自动关闭文件）
with open("D:/测试.txt", "r", encoding="UTF-8") as f:
    for line in f:
        print(f"每一行：{line}")







