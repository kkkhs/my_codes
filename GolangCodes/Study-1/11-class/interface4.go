package main

import "fmt"

//万能接口类型 arg interface{}
func myFunc(arg interface{}) {
	fmt.Println("myfunc is calling.....")
	fmt.Println(arg)
	fmt.Printf("type of arg = %T\n", arg)

	//如何区分引用的底层数据类型？
	//给interface提供"类型断言"机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string")
	} else {
		fmt.Println("arg is a string")
		fmt.Printf("Value type = %T\n", value)
	}
}

type Book struct {
	auth string
}

func main5() {
	book := Book{"khs"}
	myFunc(book)
	myFunc("abc")
	myFunc(100)
}
