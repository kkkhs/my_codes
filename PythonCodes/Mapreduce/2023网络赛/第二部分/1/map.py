import sys

# f = open("weather.txt", "r", encoding="UTF-8")

for line in sys.stdin:
    # 将每一行按制表符分开
    line = line.strip().split("\t")

    # 将年份与月份分开
    line[0] = line[0].split("-")
    line[1] = line[0][1]
    line[0] = line[0][0]
    # print(line)
    print('%s %s %s' % (line[0], line[1], line[2]))
