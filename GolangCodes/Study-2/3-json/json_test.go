package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s = Student{"张三", 18, true}
	c = Class{
		Id:       "1(2)班",
		Students: []Student{s, s, s},
	}
)

// 单元测试的目的是证明一段代码的逻辑是否正确
func TestJson(t *testing.T) {
	//json序列化
	bytes, err := json.Marshal(c)
	if err != nil {
		//测试失败Fail
		t.Fail()
	}
	//str := string(bytes)
	//fmt.Println(str)

	//json反序列化
	var c2 Class
	err = json.Unmarshal(bytes, &c2) //传入指针修改Class
	if err != nil {
		t.Fail()
	}
	//验证c和c2是否相等
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	//fmt.Printf("%+v\n", c2) //详细信息 %+v
	fmt.Println("Json没毛病")
}

func TestSonic(t *testing.T) {
	//Sonic序列化
	bytes, err := sonic.Marshal(c)
	if err != nil {
		//测试失败Fail
		t.Fail()
	}
	//str := string(bytes)
	//fmt.Println(str)

	//sonic反序列化
	var c2 Class
	err = sonic.Unmarshal(bytes, &c2) //传入指针修改Class
	if err != nil {
		t.Fail()
	}
	//验证c和c2是否相等
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	//fmt.Printf("%+v\n", c2) //详细信息 %+v
	fmt.Println("Sonic没毛病")
}

// 性能测试benchmark
func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//json序列化
		bytes, _ := json.Marshal(c)

		//json反序列化
		var c2 Class
		json.Unmarshal(bytes, &c2) //传入指针修改Class

	}

}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//sonic序列化
		bytes, _ := sonic.Marshal(c)

		//sonic反序列化
		var c2 Class
		sonic.Unmarshal(bytes, &c2) //传入指针修改Class

	}
}
