package main

import "fmt"

//固定长度数组传参(值传递)
func printArray(myArray [10]int) {
	for index, value := range myArray {
		fmt.Println("index =", index, ", value =", value)
	}
}

func main0() {
	//固定长度的数组
	var myArray1 [10]int

	myArray2 := [8]int{1, 2, 3, 4, 5}

	//数组遍历方式1
	for i := 0; i < len(myArray1); i++ {
		myArray1[i] = i + 1
		fmt.Println(myArray1[i])
	}

	//数组遍历方式2
	for index, value := range myArray2 {
		fmt.Println("index =", index, ", value =", value)
	}

	//查看数组数据类型
	fmt.Printf("Type of myArray1 = %T\n", myArray1)
	fmt.Printf("Type of myArray2 = %T\n", myArray2)

	//固定长度数组传参(值传递)
	printArray(myArray1)
	// printArray(myArray2)

}
