import threading
import time


def sing(msg):
    while True:
        print(msg)
        time.sleep(1)


def dance(msg):
    while True:
        print(msg)
        time.sleep(1)


if __name__ == '__main__':
    # 唱歌的线程(传入元组)args
    sing_thread = threading.Thread(target=sing, args=("我要唱歌 哈哈哈",))
    # 跳舞的线程(传入字典)kwargs
    dance_thread = threading.Thread(target=dance, kwargs={"msg": "我要跳舞 啦啦啦"})

    # 使用线程start
    sing_thread.start()
    dance_thread.start()
