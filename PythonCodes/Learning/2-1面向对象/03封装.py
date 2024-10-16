# 私有以__开头
class Phone:
    __current_v = 0.5

    def __keep_single_core(self):
        print("让cpu以单核模式运行")

    def call_by_5G(self):
        if self.__current_v >= 1:
            print("5G通话以开启")
        else:
            self.__keep_single_core()
            print("电量不足，正在以单核模式运行")


phone = Phone()
phone.call_by_5G()

