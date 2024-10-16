import sys
# f = open("B.txt", "r", encoding="UTF-8")
w = []
index = -1

for line in sys.stdin:
    # 将每一行按空格分开
    line = line.strip().split(" ")
    # print(line)
    line[2] = int(line[2])
    for i in range(len(w)):
        # 将该列的年份与月份与结果列表中对应
        if w[i][0] == line[0] and w[i][1] == line[1]:
            index = i
    if index == -1:
        w.append(line)
    if index != -1:
        # 如果该列月份的气温大于结果列表该月份的气温，则替换
        if line[2] > w[i][2]:
            w[i][2] = line[2]
        index = -1

# 输出年份 月份 最高气温
for i in w:
    print('%s %s %d' % (i[0], i[1], i[2]))
