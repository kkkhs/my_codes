/*
	四种变量的声明方式
*/

package main

import "fmt"

//声明全局变量，方法1，2，3可以使用
var gA int = 100
var gB int
var gC = 200

//用方法四声明全局变量,不可以
//gD := 100           :=只能在函数体内使用

func main() {
	//方法1:声明一个变量，默认值为0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	//方法2：声明一个变量并赋值
	var b string = "abcd"
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	//方法3：省去数据类型，通过值自动匹配类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcde"
	fmt.Println("cc =", cc)
	fmt.Printf("type of cc = %T\n", cc)

	//方法4：省去var关键字，直接自动匹配，
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abce"
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Printf("type of g = %T\n", g)

	//声明多个变量
	var xx, yy int = 10, 20 //同种变量
	fmt.Println("xx =", xx, ", yy =", yy)
	var kk, ll = 100, "abce" //不同变量
	fmt.Println("kk =", kk, ", ll =", ll)
	aa, bb := 100, "abce"
	fmt.Println("aa =", aa, ", bb =", bb)
	//多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, ", jj =", jj)
}
