package main

import (
	"fmt"
	"sync"
	"sync/atomic" //直接利用atomic原子类更加便利
)

var (
	n    int32
	lock = sync.Mutex{} //创建锁
)

func foo() {
	for i := 0; i < 100000; i++ {
		//n ++ 为临界区，需要考虑并发安全性，利用锁
		// lock.Lock()
		// n++
		// lock.Unlock()
		atomic.AddInt32(&n, 1) //变原子操作
	}
	fmt.Println(n)
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		foo()
	}()
	go func() {
		defer wg.Done()
		foo()
	}()
	wg.Wait()
	fmt.Println("main n=", n)
}
