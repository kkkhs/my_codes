
import sys
# f = open("B.txt", "r", encoding="UTF-8")
Class = []
index = -1

for line in sys.stdin:
    line = line.strip().split(" ")
    line[2] = int(line[2])
    for i in range(len(Class)):
        if Class[i][0] == line[0]:
            index = i
    if index == -1:
        Class.append(line)
    if index != -1:
        if Class[index][2] < line[2]:
            Class[index][1] = line[1]
            Class[index][2] = line[2]
            index = -1

for i in Class:
    print('%s,%s,%s' % (i[0], i[1], i[2]))
