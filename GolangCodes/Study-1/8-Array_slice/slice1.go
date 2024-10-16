package main

import "fmt"

//切片传参(引用传递)
func printarray(myArray []int) {
	myArray[0] = 100

	for _, value := range myArray {
		fmt.Println("value =", value)
	}

}

func main1() {
	//切片slice定义
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("Type of myArray = %T\n", myArray)

	printarray(myArray)
}
