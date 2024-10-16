import sys

words = []

for line in sys.stdin:
    line = line.strip()
    word = line.split(" ")
    words.append(word)

for i in words:
    for j in i:
        print("%s 1" % j)
