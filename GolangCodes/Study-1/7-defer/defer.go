package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func call....")
	return 0
}

func returnFunc() int {
	fmt.Println("return func call....")
	return 0
}

//return比defer先执行
func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}

func main() {
	//写入defer关键字 类似于final 打印使用栈机制
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	defer fmt.Println("main end3")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
