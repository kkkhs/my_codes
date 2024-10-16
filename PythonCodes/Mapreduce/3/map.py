import sys

# f = open("source.txt", "r", encoding="UTF-8")

for line in sys.stdin:
    # print(line)
    line = line.strip().split(",")
    # print(line)
    score = line[2:]
    score = [int(i) for i in score]
    # print(score)
    maxscore = max(score)
    # print('%.1f' % avg)
    print('%s %s %d' % (line[0], line[1], maxscore))
    