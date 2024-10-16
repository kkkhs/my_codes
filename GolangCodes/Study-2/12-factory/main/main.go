package main

import (
	"fmt"
	"khs/Study-2/12-factory/model"
)

func main() {
	//创建Student实例
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }

	stu := model.NewStudent("tom~", 88.8) //指针
	fmt.Println(*stu)
}
