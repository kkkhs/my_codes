# 只有抽象方法的类成为抽象类
class Animal:
    # 抽象方法
    def speak(self):
        pass


class Dog(Animal):
    def speak(self):
        print("汪汪汪")


class Cat(Animal):
    def speak(self):
        print("喵喵喵")


def noise(animal: Animal):
    animal.speak()


dog = Dog()
cat = Cat()

noise(dog)
noise(cat)

