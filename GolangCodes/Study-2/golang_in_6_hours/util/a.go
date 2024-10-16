package util

import "fmt"

var (
	Name = "khs"
)

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("init util package1")
}
