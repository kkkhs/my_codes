package main

import "fmt"

//方法/类大写表示其他包也可以访问
type Hero struct {
	//属性:大写类似于public,小写为private
	Name  string
	Ad    int
	level int
}

func (this Hero) GetName() {
	fmt.Println("Name =", this.Name)
}

//当前this是拷贝副本:值传递
func (this Hero) setName(NewName string) {
	this.Name = NewName
}

//使用指针：引用传递
func (this *Hero) setName1(NewName string) {
	this.Name = NewName
}

func main1() {
	hero := Hero{Name: "khs", Ad: 10086, level: 6}
	fmt.Println(hero)
	hero.setName("hhhhh")
	hero.GetName()
	fmt.Println("--------------")
	hero.setName1("hhhhh")
	hero.GetName()
}
