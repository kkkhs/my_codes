// 关闭channel的特点
package main

import "fmt"

func main3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close可以关闭一个channel
		close(c)
	}()

	/*
		for {
			//ok为真则channel没有关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/
	//简写-------->
	//使用 range 不断迭代操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished..........")
}
