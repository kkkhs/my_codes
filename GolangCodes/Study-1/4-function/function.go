package main

import "fmt"

//单值返回
func foo1(a string, b int) int {
	fmt.Println("a =", a, ", b =", b)

	c := 100

	return c
}

//多值返回(匿名)
func foo2(a string, b int) (int, int) {
	fmt.Println("a =", a, ", b =", b)

	return 666, 777
}

//多值返回(有形参名)
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("foo3---------------------")
	fmt.Println("a =", a, ", b =", b)

	fmt.Println("r1 =", r1, ", r2 =", r2) //默认形参返回值为0/空
	//有名称的返回值赋值
	r1 = 100
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("foo4---------------------")
	fmt.Println("a =", a, ", b =", b)

	//有名称的返回值赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c =", c)

	ret1, ret2 := foo2("hahaha", 123)
	fmt.Println("ret1 =", ret1, ", ret2 =", ret2)

	ret1, ret2 = foo3("hahahahaha", 23333)
	fmt.Println("ret1 =", ret1, ", ret2 =", ret2)

	ret1, ret2 = foo4("hahahahaha", 44444)
	fmt.Println("ret1 =", ret1, ", ret2 =", ret2)
}
