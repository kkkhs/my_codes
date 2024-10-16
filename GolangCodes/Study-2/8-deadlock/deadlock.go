package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()

	go func() {
		time.Sleep(time.Hour) //Sleep阻塞会自己解除 在执行Sleep时不知道后面的代码>>>>>等待
		//wg.Done()
	}()
	//time.Sleep(time.Hour)  //1、Sleep阻塞会自己解除 在执行Sleep时不知道后面的代码>>>>>等待
	//ch <- 1 //2、所有协程都在写入，而没有协程读取 >>>>死锁
	wg.Wait() //3、互相死依赖 >>>>死锁

}
