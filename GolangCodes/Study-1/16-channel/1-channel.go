package main

import "fmt"

func main1() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行....")

		c <- 666 //将666发送给c
	}()

	// <-c     //从c中接收数据并丢弃
	num := <-c //从c中接收数据并赋值给num
	fmt.Println("num =", num)
	fmt.Println("main goroutine 结束")
}
