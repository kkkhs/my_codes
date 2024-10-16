f1 = open("D:/test.txt", "r")
f2 = open("D:/bill.txt.bak", "w")

for line in f1:
    line = line.strip()

    if line.split(",")[4] == "测试" :
        continue

    f2.write(line)
    f2.write("\n")

# f2.flush() 可省略

f1.close()
f2.close()


