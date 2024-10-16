package main

import "fmt"

func fakeChange(p int) {
	p = 12345
}
func changeValue(p *int) {
	*p = 12345
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 10
	var b int = 20
	fakeChange(a)
	fmt.Println("a =", a)

	changeValue(&a)
	fmt.Println("a =", a)

	fmt.Println("------------")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	swap(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
