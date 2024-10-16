package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine :i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main1() {
	//创建一个go程 去执行newTask()流程
	go newTask()

	fmt.Printf("main goroutine exi")
	i := 0

	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
