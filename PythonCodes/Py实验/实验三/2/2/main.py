import string


def check_password(pwd):
    if len(pwd) < 6 or len(pwd) > 12:  # 检测长度
        return "密码长度不符合! 应为6到12位"
    flag = [False] * 4
    for ch in pwd:
        if ch in string.ascii_lowercase:  # [a-z]之间至少有 1 个字母
            flag[0] = True
        if ch in string.ascii_uppercase:  # [A-Z]之间至少有 1 个字母
            flag[1] = True
        if ch in string.digits:  # [0-9]之间至少有 1 个数字
            flag[2] = True
        if ch in '$#@':  # [$＃@]中至少有 1 个字符
            flag[3] = True
    if flag.count(True) == 4:  # 如果通过了全部测试
        return pwd
    return "格式不对"


str1 = input("请输入密码，逗号分隔:")
alist = str1.split(",")  # 使用，分割
for pwd in alist:  # 循环遍历各个密码
    print(check_password(pwd))


