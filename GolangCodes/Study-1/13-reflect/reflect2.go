package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is calling...")
	fmt.Printf("%v\n", this)
}

func main2() {
	user := User{1, "ksh", 18}

	DoFiledAndMothod(user)
}
func DoFiledAndMothod(intput interface{}) {
	intputType := reflect.TypeOf(intput)
	fmt.Println("input type =", intputType)
	intputValue := reflect.ValueOf(intput)
	fmt.Println("input value =", intputValue)

	//通过type获取里面的字段
	for i := 0; i < intputType.NumField(); i++ {
		field := intputType.Field(i)
		value := intputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	//通过type获取里面的方法，调用
	for i := 0; i < intputType.NumMethod(); i++ {
		m := intputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}
