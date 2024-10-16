package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		//return
		func() {
			defer fmt.Println("B.defer")
			//提前退出整个goroutine
			runtime.Goexit()

			fmt.Println("B")
		}() //后加()表示调用该函数

		fmt.Println("A")
	}() //后加()表示调用该函数

	//goroutine匿名函数无法得到返回值
	go func(a int, b int) bool {
		fmt.Println("a =", a, ", b =", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
