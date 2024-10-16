import time

f = open("D:/test.txt", "w", encoding="UTF-8")

# write写入
f.write("帅哥在此")

# flush刷新内容到内存中
f.flush()

# close方法内置了flush功能
f.close()

