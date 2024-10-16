package main

import "fmt"

//接口
type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book...")
}
func (this *Book) WriteBook() {
	fmt.Println("Write a book...")
}
func main() {
	// b:pair<type:Book, value:Book{}地址>
	b := &Book{}

	//r:pair<type: , value: >
	var r Reader

	//r:pair<type:Book, value:Book{}地址>
	r = b
	r.ReadBook()

	//w:pair<type: , value: >
	var w Writer
	//w:pair<type:Book , value:Book{}地址 >
	w = r.(Writer) //断言为何能成功？ 因为w r 具体的type是一致的

	w.WriteBook()
}
