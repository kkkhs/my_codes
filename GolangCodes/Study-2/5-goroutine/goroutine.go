package main

import (
	"fmt"
	"sync"
)

func foo() {
	fmt.Println("foo finish")
}

func main1() { //main协程

	//使用WaitGroup等待子协程先完成
	wg := sync.WaitGroup{}
	wg.Add(2) //准备开辟2个协程

	go func() { //创建一个子协程来运行foo
		defer wg.Done() //最后让Add-1
		fmt.Println("foo finish")
	}()
	go func() {
		defer wg.Done() //最后让Add-1
		fmt.Println("匿名函数 finish")
	}()

	wg.Wait()

	//使main协程等待子协程完成
	//time.Sleep(1 * time.Second)
}
