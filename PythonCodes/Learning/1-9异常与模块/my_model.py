# 自定义模块
# all变量 导包 只会有testa
__all__ = ['testa']

def testa(a, b):
    print(a + b)

def testb(a, b):
    print(a - b)

# 用于测试模块代码  导包不会执行下段代码
if __name__ == '__main__':
    testa(1, 2)



