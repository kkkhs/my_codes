import random
ran = random.randint(1, 50)  # 设置随机种子数
n = 10  # 设置猜的次数
print(f'猜数字游戏(共有{n}次机会哦)')  # 提示玩家
count = 0     # 记录玩家猜测的次数
while True:
    guess = int(input('请输入您要猜的数字：'))
    if count >= n:
        print(f'你的次数已经用完了哦， 正确答案是{ran}')
        break
    else:
        count += 1
        if guess == ran:
            print('right!')
            break
        elif guess > ran:
            print('too large!')
        else:
            print('too small!')


