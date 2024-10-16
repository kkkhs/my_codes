package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero1 struct {
	Name string
	Age  int
}

type HeroSlice []Hero1

func (this HeroSlice) Len() int {
	return len(this)
}

func (this HeroSlice) Less(i, j int) bool {
	return this[i].Age < this[j].Age //从小到大
}

func (this HeroSlice) Swap(i, j int) {
	temp := this[i]
	this[i] = this[j]
	this[j] = temp
}

func main() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero1{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	for _, v := range heroes {
		fmt.Println(v)
	}
	sort.Sort(heroes)
	fmt.Println("------------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
