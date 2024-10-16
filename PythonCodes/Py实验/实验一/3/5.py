import math
weeks = 365 / 7  # 求出一年多少周 52
d = 365 % 7  # 最后一天是星期几 1
de = 0.99 ** (2 * int(weeks))  # 求出下降的比例
increase = 37.78 / de  #  求出上升的总比例
n = math.pow(increase, 1 / (5 * int(weeks) + 1))
percent = (n - 1) * 100  # 反向求结果
print(f'每天努力{percent}%')  # 输出答案

