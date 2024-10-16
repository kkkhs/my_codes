from typing import Union

var_1: int = 10
var_2: float = 3.1415926
var_3: str = "khs"

my_list: list[int] = [1, 2, 3]
my_tuple: tuple[int, str, bool] = (1, "2", True)
my_set: set[int] = {1, 2, 3}
my_dict: dict[str, int] = {"ksh": 1}


class Student:
    pass


student: Student = Student()

# Union类型注解
my_list1: list[Union[int, str]] = [1, 2, 3, "khs"]


def func(data: Union[int, str]) -> Union[int, str]:
    pass


func(1)
