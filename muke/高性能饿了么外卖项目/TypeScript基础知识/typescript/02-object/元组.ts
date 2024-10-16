// 存储的元素数据类型不同, 元组是可变的
// var tuple_name = [value1,value2,value3,…value n]

var mytuple1 = [10, "Runoob"];

var mytuple2 = [];
mytuple2[0] = 120
mytuple2[1] = 234

// push() 向元组添加元素，添加在最后面。
// pop() 从元组中移除元素（最后一个），并返回移除的元素。

// 更新元组元素
mytuple2[0] = 121

// 解构元组
// 我们也可以把元组元素赋值给变量，如下所示：
var a = [10, "Runoob"]
var [b, c] = a