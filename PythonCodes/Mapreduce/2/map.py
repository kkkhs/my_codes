import sys

# f = open("source.txt", "r", encoding="UTF-8")

for line in sys.stdin:
    # print(line)
    line = line.strip().split(",")
    # print(line)
    score = line[2:]
    score = [int(i) for i in score]
    # print(score)
    avg = float(sum(score)) / len(score)
    # print('%.1f' % avg)
    line[1] = avg
    print('%s %f' % (line[0], line[1]))