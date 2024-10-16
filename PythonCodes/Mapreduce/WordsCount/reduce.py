import sys
from operator import itemgetter

words = []
index = -1

for line in sys.stdin:
    line = line.strip()
    line = line.split(" ")
    line[1] = int(line[1])
    for i in range(len(words)):
        if words[i][0] == line[0]:
            index = i
    if index == -1:
        words.append(line)
    if index != -1:
        words[index][1] += 1
        index = -1

for i in words:
    print(i)
