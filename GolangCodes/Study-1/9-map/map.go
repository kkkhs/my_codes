package main

import "fmt"

func main1() {
	//1、声明空的map1 key为string value为string
	var map1 map[string]string

	if map1 == nil {
		fmt.Println("map1是一个空map")
	}

	//2、用make给map分配10个空间
	map1 = make(map[string]string, 10)

	map1["one"] = "java"
	map1["two"] = "c"
	map1["three"] = "python"
	fmt.Println(map1)

	//3、声明
	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c"
	fmt.Println(map2)

	//4、声明
	map3 := map[int]string{
		1: "php",
		2: "go",
	}
	fmt.Println(map3)

}
