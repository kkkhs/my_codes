package main

import "fmt"

//定义一个别名
type myint int

//定义一个结构体
type Book struct {
	title string
	auth  string
}

//struct传参：值传递
func printBook(book Book) {
	book.auth = "123"
	fmt.Println(book)
}

func main() {
	var a myint = 10
	fmt.Printf("type of a = %T\n", a)

	//定义结构体
	var book1 Book
	book1.title = "golang"
	book1.auth = "khs"
	fmt.Println(book1)
	fmt.Printf("type of boo1 = %T\n", book1)
	fmt.Println("-----------------")
	printBook(book1)
	fmt.Println(book1)
}
