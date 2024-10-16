class Employee:
    def __init__(self, name, id, salary):
        self.name = name
        self.id = id
        self.salary = salary

    def pay(self):
        print(f"Employee pay: {self.salary}")

    def show(self):
        print("Employee show:")
        print(f'Name: {self.name}\t Id: {self.id}\t Salary: {self.salary}')


class Manager(Employee):
    def __init__(self):
        pass

    def pay(self):
        print("Manager pay: 1000")

    def show(self):
        print("Manager show: S H O W")


class Salesman(Employee):
    def __init__(self):
        pass

    def pay(self):
        print("Salesman pay: 10000")

    def show(self):
        print("Salesman show: S h o w")


if __name__ == '__main__':
    em1 = Employee("张三", 2022216, 100)
    em1.pay()
    em1.show()
    ma1 = Manager()
    ma1.pay()
    ma1.show()
    Sa1 = Salesman()
    Sa1.pay()
    Sa1.show()
