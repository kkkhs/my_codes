class Myqueue:
    def __init__(self):
        self.maxsize = 10  # 最多存储元素个数
        self.data = []  # 利用列表实现
        self.current = 0  # 当前几个元素

    def isEmpty(self):
        if len(self.data) == 0:
            return True
        return False

    def isFull(self):
        if len(self.data) == 10:
            return True
        return False

    def GetFirst(self):
        if self.current == 0:
            return "队列为空"
        return self.data[0]

    def EnterQueue(self, x):
        if self.current == 10:
            return "队列为空"
        else:
            self.data.append(x)
            self.current += 1
            print(f"入队成功: {x}")

    def OutQueue(self):
        if self.current == 0:
            return "队列为空"
        self.data.pop()
        self.current -= 1
        print("出队成功")

    def show(self):
        print(self.data)


if __name__ == '__main__':
    myqueue = Myqueue()
    for i in range(1, 5):
        myqueue.EnterQueue(i)
        myqueue.show()
    print(myqueue.GetFirst())
    myqueue.OutQueue()
    myqueue.show()
