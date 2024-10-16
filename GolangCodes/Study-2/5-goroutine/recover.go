package main

import (
	"fmt"
	"time"
)

func f0() {
	defer func() { //最后匿名函数调用recover函数使子协程不影响main协程
		err := recover() //recover可以捕获panic
		if err != nil {
			fmt.Println("发生了panic:", err)
		}
	}()

	a, b := 3, 0
	fmt.Println(a, b)
	_ = a / b //此处会发生panic 会使整个go进程结束

	fmt.Println("f0 finish")
}
func main() {
	go f0()
	time.Sleep(3 * time.Second)
	fmt.Println("main finish")
}
