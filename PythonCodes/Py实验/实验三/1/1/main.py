objectList = []


class Student:
    def __init__(self, id, name, classid, sex, gradedict) -> None:
        self.id = id
        self.name = name
        self.classid = classid
        self.sex = sex
        self.gradedict = gradedict

    def SetGrade(self, sub_name, newgrade):
        self.gradedict[sub_name] = newgrade

    def show(self):
        print(self.id, self.name, self.sex, self.gradedict, self.Get_average())

    def Get_average(self):
        temp = self.gradedict.values()
        sum = 0
        for x in temp:
            sum += int(x)
        return sum / 3


def Sub_dict(alist):
    blist = "".join(alist).split(" ")
    sub_dict = {}
    for x in blist:
        x = ''.join(x).split(":")
        # print(x)
        sub_dict[x[0]] = x[1]
    # print(sub_dict)
    return sub_dict


def Open_Txt():
    with open(r'E:\Codes\PythonCodes\Py实验\实验三\1.txt', 'r', encoding='UTF-8') as fp:
        # 清空全局变量
        objectList.clear()
        # 读取文件
        items = fp.readlines()
        # print(items)
        for item in items:
            item = item.replace('\n', '')
            # print(item)
            # 分隔成绩前的元素
            itemlist = item.split(" ", 4)
            # 成绩转字典
            itemlist[4] = Sub_dict(itemlist[4])
            # 输出信息
            print(itemlist)
            # 创造对象
            itemlist[1] = Student(itemlist[0], itemlist[1], itemlist[2], itemlist[3], itemlist[4])
            # 记录对象
            objectList.append(itemlist[1])


def Set():
    name = input("请输入要修改学生的姓名: ")

    with open(r'E:\Codes\PythonCodes\Py实验\实验三\1.txt', 'r', encoding='UTF-8') as fp:
        items = fp.readlines()
        # print(items)
        newitems = []
        flagname = 0
        flagsub = 0
        # print(1,items)
        for item in items:
            item = item.split(" ")
            # print(item[1])
            if item[1] == name:
                flagname = 1
                print(item)
                sub = input("请输入科目")
                newgrade = input("请输入成绩")
                for i in range(4, 7):
                    temp = item[i].split(":")
                    if temp[0] == sub:
                        flagsub = 1
                        item[i] = sub + ":" + newgrade
                        print("修改成功", item[i])
            newitems.append(item)
        if flagname == 0:
            print("没有该同学")

        if flagsub == 0:
            print("没有这个科目")

    with open(r'E:\Codes\PythonCodes\Py实验\实验三1.txt', 'w', encoding='UTF-8') as fp:
        for item in newitems:
            # print(item)
            i = 0
            for it in item:
                temp = ''.join(it)
                fp.write(temp)
                if i != 6:  # 防止写到行末尾出错
                    fp.write(" ")
                    i += 1


def Sort():
    # 对象字典，id:平均成绩
    gradeDict = {}
    # 字典值集合
    DictValues = []
    for x in objectList:
        avarge = Student.Get_average(x)
        gradeDict[x] = avarge
        DictValues.append(avarge)
    # 字典的值列表，排序
    DictValues.sort()
    for i in range(len(DictValues) - 1, -1, -1):
        for x in objectList:
            # 列表最大的值和字典匹配
            if gradeDict.get(x) == DictValues[i]:
                # 排名
                print(len(DictValues) - i)
                # 输出所有信息
                Student.show(x)
                # 删除已输出，避免和后序成绩相同时重复
                del gradeDict[x]


def main():
    while True:
        print("1,查看最新学生信息")
        print("2,修改成绩")
        print("3,按照成绩排名:")
        print("4,退出")
        n = int(input("input:"))
        if n == 1:
            Open_Txt()
        if n == 2:
            Set()
        if n == 3:
            Sort()
        if n == 4:
            break


if __name__ == "__main__":
    main()
