class Vehicle:
    def __init__(self, MaxSpeed, Weight):
        self.__Maxspeed = MaxSpeed
        self.__Weight = Weight

    def SetMaxSpeed(self, MaxSpeed):
        self.__Maxspeed = MaxSpeed

    def show(self):
        print(self.__class__.__name__, self.__Maxspeed, self.__Weight)


class Bicycle(Vehicle):
    def __init__(self, MaxSpeed, Weight, Height):
        super().__init__(MaxSpeed, Weight)  # 调用父类的成员函数
        # Vehicle.__init__(self,MaxSpeed,Weight)
        self.__Height = Height

    def show(self):
        # print("Bicycle:")
        super().show()
        print(self.__Height)

    def SetFaMaxSpeed(self, Father_Name, Maxspeed):
        Vehicle.SetMaxSpeed(Father_Name, Maxspeed)

    def __GetHeight(self):
        return self.__Height

    def __SetHeight(self, newHeight):
        self.__Height = newHeight

    Height = property(__GetHeight, __SetHeight)  # 只读，可修改属性


if __name__ == '__main__':
    Ve1 = Vehicle(1000, 2000)
    Ve1.show()
    Bi1 = Bicycle(3000, 4000, 5000)
    Bi1.show()
    Bi1.SetFaMaxSpeed(Ve1, 6000)  # 改变父类私有属性
    Ve1.show()
