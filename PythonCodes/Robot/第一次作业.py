def hear(msg):
    i = 0
    for i in range(len(msg)):
        if msg[i] == 'hear':
            break
    x = i + 1
    time = msg[x]
    print("在" + time + "周期" + "  hear ")


def send(msg, i):
    x = 0
    for x in range(len(msg)):
        if msg[x] == 'passto':
            break

    print("从" + msg[i + 2] + "方向听到了" + "passto(" + msg[x + 1] + "," + msg[x + 2] + ")")


def see(msg, i):
    time = msg[i + 1]
    print("在" + time + "周期" + " see ")


def ball(msg, i):
    i = i + 2
    print(msg[i])

    Distance = msg[i]
    Direction = msg[i + 1]
    DistChng = msg[i + 2]
    DirChng = msg[i + 3]
    print(
        " Ball 距离我的 Distance 是 " + Distance + " Direction是" + Direction + " DistChng是" + DirChng + "  DirChng是" + DirChng)


def player(msg, i):
    x = i + 4

    Distance = msg[x]
    Direction = msg[x + 1]
    DistChng = msg[x + 2]
    DirChng = msg[x + 3]
    BodyDir = msg[x + 4]
    HeadDir = msg[x + 5]

    print(msg[i + 1] + " " + msg[
        i + 2] + "  距离我的 Distance 是 " + Distance + "  Direction是" + Direction + "  DistChng是" + DistChng + "  DirChng是" + DirChng)


def goal(msg, i):
    x = i + 3

    Distance = msg[x]
    Direction = msg[x + 1]

    print("goal " + msg[i + 1] + "  距离我的 Distance 是 " + Distance + "  Direction是" + Direction)


if __name__ == '__main__':
    msg = input("请输入你需要解析的字符串信息：")
    # msg = '(hear 1022 -30 passto(23,24))(see 1022 ((ball) -20 20 1 -2) ((player hfut1 2) 45 23 0.5 1 22 40 ) ((goal r) 12 20) ((Line r) -30))'
    # msg = '(see 2022  ((player hfut2 8) 40 23 1  6 21  33 ) ((goal r) 15 30)  ((f r t 20)5   24 ))(hear 2022 -10 “passball”)'
    # msg = '(hear 2022 -10 “passball”)'
    msg = msg.replace('(', '').replace(')', '').split(',')
    msg = msg[0].split(' ')
    # print(msg)
    for i in range(len(msg)):
        if msg[i] == 'hear':
            hear(msg)
            send(msg, i)
        elif msg[i] == 'see':
            see(msg, i)
        elif msg[i] == 'ball':
            ball(msg, i)
        elif msg[i] == 'player':
            player(msg, i)
        elif msg[i] == 'goal':
            goal(msg, i)

    """
样例输入：
(hear 1022 -30 passto(23,24))(see 1022 ((ball) -20 20 1 -2) ((player hfut1 2) 45 23 0.5 1 22 40 ) ((goal r) 12 20) ((Line r) -30))
样例输出：
在 1022 周期 hear 从 -30 方向 听到了 passto(23,24)；
在 1022 周期 see   ball 距离我的 Distance 是 -20, Direction 是 20, DirChng 是 1, DistChng 是 -2,
player hfut1 2 距离我的 Distance 是 45, Direction 是 23, DirChng 是 0.5, DistChng 是 1, 它的 BodyDir 是 22 和 HeadDir 是 40;
goal r 距离我的 Distance 是 12, Direction 是 20,
Line r 距离我的 Distance 是 -30,
"""
