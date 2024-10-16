package main

import "fmt"

const (
	//可以在const()添加一个关键字，每行的iota都会累加1，第一行iota的默认值为0
	BEIJING  = 10 * iota
	SHANGHAI //iota = 10 * 1
	SHENZHEN //iota = 10 * 2
)

const (
	a, b = iota + 1, iota + 2 //iota = 0,  /1, 2
	c, d                      //iota = 1,  /2, 3
	e, f                      //iota = 2,  /3, 4

	g, h = iota * 2, iota * 3 //iota = 3,  /6, 9
	i, k                      //iota = 4,  /8, 12
)

func main() {
	//常量（只读属性）
	const length int = 10

	fmt.Println("length =", length)
	// length = 100       不可修改

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

	//var a int = iota  //iota只能配合const()使用
}
