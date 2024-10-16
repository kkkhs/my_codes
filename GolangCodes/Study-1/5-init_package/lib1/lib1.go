package lib1

import "fmt"

//当前lib1包提供的接口
func Lib1test() {
	fmt.Println("lib1test().......")
}

func init() {
	fmt.Println("lib1, init().......")
}
