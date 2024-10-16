package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat...........")
}

func (this *Human) Run() {
	fmt.Println("Human Run...........")
}

type SuperMan struct {
	Human //该类公有继承了Human

	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eating...........")
}

func (this *SuperMan) fly() {
	fmt.Println("SuperMan flying..............")
}

func (this *SuperMan) print123() {
	fmt.Println(*this)
}

func main2() {
	h := Human{"Zhang3", "man"}
	h.Eat()
	h.Run()

	//定义一个子类对象1
	s := SuperMan{Human{"li4", "female"}, 3}
	fmt.Println(s)
	s.Run()
	s.Eat()

	//定义一个子类对象1
	var s2 SuperMan
	s2.name = "123"
	s2.sex = "123"
	s2.level = 123
	s2.print123()
}
