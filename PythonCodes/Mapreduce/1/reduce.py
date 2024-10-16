
import sys
# f = open("B.txt", "r", encoding="UTF-8")
Class = []
index = -1

for line in sys.stdin:
    line = line.strip().split(" ")
    print(line)
    line[1] = int(line[1])
    line[2] = float(line[2])
    for i in range(len(Class)):
        if Class[i][0] == line[0]:
            index = i
    if index == -1:
        Class.append(line)
    if index != -1:
        Class[index][1] += 1
        Class[index][2] += line[2]
        index = -1

for i in Class:
    avg = (i[2]/i[1])
    print('%s %d %.1f'%(i[0], i[1], avg))
