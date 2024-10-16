package maths

import "fmt"

func sub(a, b int) int {
	return a - b
}

func init() {
	fmt.Println("init maths package1")
}

func init() {
	fmt.Println("init maths package2")
}
