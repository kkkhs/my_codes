package main

import (
	"fmt"
)

// interface本质是一个指针
type AnimalIF interface {
	Sleep()
	getColor() string
	GetType() string
}

// 具体的类1
type Cat struct {
	//接口不需要继承
	//子类必须实现接口的所有方法
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleeping.........")
}

func (this *Cat) getColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类2
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleeping.........")
}

func (this *Dog) getColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
}

func main3() {
	var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep()
	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{"green"}
	dog := Dog{"dada"}
	showAnimal(&cat)
	showAnimal(&dog)
}
