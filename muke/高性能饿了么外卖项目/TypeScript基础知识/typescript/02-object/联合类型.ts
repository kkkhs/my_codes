// 联合类型（Union Types）可以通过管道(|)将变量设置多种类型，赋值时可以根据设置的类型来赋值。
// Type1|Type2|Type3 

var val: string | number
val = 12
console.log("数字为 " + val)
val = "Runoob"
console.log("字符串为 " + val)

// 联合类型数组
// 我们也可以将数组声明为联合类型
var arr2: number[] | string[];
var i: number;
arr2 = [1, 2, 4]
console.log("**数字数组**")

arr2 = ["Runoob", "Google", "Taobao"]
console.log("**字符串数组**")
