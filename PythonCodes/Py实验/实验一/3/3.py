s = input("请输入字符串：")
j = len(s) - 1
for i in range(0, j):
    if s[i] != s[j]:
        print('不是回文')
        break
    j -= 1
else:
    print('是回文')
