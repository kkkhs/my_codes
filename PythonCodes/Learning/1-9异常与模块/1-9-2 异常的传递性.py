def func1():
    print("func1开始执行")
    num = 1/0
    print("func1结束执行")

def func2():
    print("func2开始执行")
    func1()
    print("func2结束执行")

def main():
    func2()

main()