import sys
w = []
index = -1

for line in sys.stdin:
    # 将每一行按空格分开
    line = line.strip().split(" ")
    line[2] = int(line[2])
    for i in range(len(w)):
        # 将该列的年份与月份与结果列表中对应
        if w[i][0] == line[0] and w[i][1] == line[1]:
            index = i
    if index == -1:
        w.append(line)
    if index != -1:
        # 将改行的气温加入到结果列表中
        w[index].append(line[2])
        index = -1

for i in w:
    # 切割温度
    t = i[2:]
    # 将温度从大到小排序
    t.sort(reverse=True)
    # 切割出前三的温度
    t = t[:3]
    i[2] = t
    print('%s %s [%d,%d,%d]' % (i[0], i[1], i[2][0], i[2][1], i[2][2]))
