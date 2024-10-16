package main

import "fmt"

func main2() {
	//1、声明slice1是一个切片，且有默认值，len为3
	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//2、声明slice2是切片但没有分配空间
	var slice2 []int
	slice2 = make([]int, 3) //通过make给slice开辟空间，默认值为0

	slice2[0] = 100

	//3、声明式通过make给slice开辟空间，默认值为0
	var slice3 []int = make([]int, 3)
	slice3[0] = 300

	//4、声明简写开辟空间
	slice4 := make([]int, 3)
	slice4[1] = 50

	//判断一个slice是否为空
	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	} else {
		fmt.Println("slice1不是一个空切片")
	}

}
