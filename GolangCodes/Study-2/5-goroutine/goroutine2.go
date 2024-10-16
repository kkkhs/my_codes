package main

import (
	"fmt"
	"time"
)

func f2() {
	time.Sleep(2 * time.Second)
	fmt.Println("f2 finish")
}

func f1() {
	fmt.Println("f1 finish")
	go f2() //创建父子协程
}

// main->f1->f2
func main2() { //main协程
	go f1()
	time.Sleep(3 * time.Second)
}
