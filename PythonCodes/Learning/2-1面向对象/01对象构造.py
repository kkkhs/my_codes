class Student:
    name = None
    gender = None
    nationality = None
    native_place = None
    age = None

    def __init__(self, name, gender, nationality, native_place, age):
        self.name = name
        self.gender = gender
        self.nationality = nationality
        self.native_place = native_place
        self.age = age
        print("对象正在创建")

    def showinfo(self):
        print(f"{self.name}, {self.gender}, {self.nationality}, {self.native_place}, {self.age}")


stu_2 = Student("ljy", "Woman", "Japan", "Tokyon", 10)
stu_2.showinfo()

