
def reverse(s):
    # 切片[起点：终点：步长]
    return s[::-1]

def substr(s, x, y):
    return s[x:y:1]

if __name__ == '__main__':
    print(reverse("abc"))
    print(substr("abcde", 2, 4))
