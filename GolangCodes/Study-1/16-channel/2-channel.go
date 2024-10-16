package main

import (
	"fmt"
	"time"
)

func main2() {
	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c) =", len(c), ", cap(c) =", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子goroutine正在运行:发送的元素为：", i, ", len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num =", num)
	}

	fmt.Println("main结束")
}
