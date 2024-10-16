package main

import "fmt"

//map传参(引用传递)
func printMap(cityMap map[string]string) {
	cityMap["CN"] = "BEIJING"
	cityMap["US"] = "NewYork"
	for key, value := range cityMap {
		fmt.Println("key =", key, ", value =", value)
	}
}

func main() {
	cityMap := make(map[string]string)
	//添加value
	cityMap["CN"] = "BEIJING"
	cityMap["US"] = "NewYork"
	cityMap["JP"] = "Tokyo"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key =", key, ", value =", value)
	}

	//删除
	delete(cityMap, "US")
	fmt.Println("--------------------------")
	for key, value := range cityMap {
		fmt.Println("key =", key, ", value =", value)
	}

	//修改
	cityMap["CN"] = "ZHUZHOU"
	fmt.Println("--------------------------")
	printMap(cityMap)
}
