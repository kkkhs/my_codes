
import sys
# f = open("B.txt", "r", encoding="UTF-8")
Class = []
index = -1

for line in sys.stdin:
    line = line.strip().split(" ")
    line[1] = float(line[1])
    for i in range(len(Class)):
        if Class[i][0] == line[0]:
            index = i
    if index == -1:
        Class.append(line)
    if index != -1:
        Class[index].append(line[1])
        index = -1

for i in Class:
    cource = i[0]
    result = i[1:]
    # print(cource)
    # print(result)
    result.sort()
    result.reverse()
    # print(result)
    o = open('%s.txt' % cource, 'w')
    for score in result:
        o.write('%.1f\n' % score)
    o.close()
