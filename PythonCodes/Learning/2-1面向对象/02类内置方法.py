class Student:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    # __str__方法（直接输出类内容）
    def __str__(self):
        return f"Student类对象，name:{self.name}, age:{self.age}"

    # __lt__方法(类比较 > <)
    def __lt__(self, other):
        return self.age > other.age

    # __le__方法(类比较 >=  <= )
    def __le__(self, other):
        return self.age >= other.age

    # __eq__方法(类比较 == )
    def __eq__(self, other):
        return self.age == other.age


stu = Student("顶真", 18)
print(stu)

stu2 = Student("王源", 17)
print(stu.__lt__(stu2))
print(stu2 > stu)
print(stu2 < stu)
stu2.age = 18
print(stu2 <= stu)
print(stu2 == stu)

