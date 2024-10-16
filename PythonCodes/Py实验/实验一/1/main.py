import random
ran = random.randint(1, 25)  # 设置随机种子数
print('猜数字游戏(只有一次机会哦！)')  # 提示玩家
guess = int(input('请输入您要猜的数字：'))
if guess == ran:
    print('right!')
elif guess > ran:
    print(f'too large! The right answer is {ran}')
else:
    print(f'too small! The right answer is {ran}')