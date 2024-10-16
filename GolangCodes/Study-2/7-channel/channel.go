package main

import (
	"fmt"
	"time"
)

func main1() {
	ch := make(chan int, 1)
	fmt.Println(time.Now().Unix())

	go func() {
		a := <-ch
		fmt.Println(time.Now().Unix(), ": 读出", a)
	}()
	time.Sleep(2 * time.Second)
	go func() {
		ch <- 4
		fmt.Println(time.Now().Unix(), ": 写入4成功")
	}()

	time.Sleep(1 * time.Second)
}
