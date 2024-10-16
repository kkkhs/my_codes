package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 100)

	//两个生产者
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	//1个消费者
	//wg2 := sync.WaitGroup{}
	//wg2.Add(1)
	mc := make(chan struct{}, 0) //使用管道阻塞main，起到协调同步作用
	go func() {
		sum := 0
		for {
			a, ok := <-ch //channel被close，且channel为空，ok为false
			if !ok {
				break
			} else {
				sum += a
			}
		}
		fmt.Println("sum=", sum)
		mc <- struct{}{} //当不关心channel写入的元素类型，可用空结构体填充，防止内存消耗，增强可读性
		//wg2.Done()
	}()

	wg.Wait()
	close(ch) //使for循环结束，关闭管道，使数据无法写入管道，但可以继续读

	<-mc
	//wg2.Wait()
}
