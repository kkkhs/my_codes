package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func main() {
	//数据库连接参数:
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "test"     //数据库名
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/test?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接失败")
	}

	var studentList []Student
	for i := 0; i < 100; i++ {
		email := fmt.Sprintf("%d@qq.com", i+1)
		studentList = append(studentList, Student{
			Name:   fmt.Sprintf("机器人%d号", i+1),
			Age:    21,
			Gender: true,
			Email:  &email,
		})
	}
	//db.Create(&studentList)

	db.Find(&studentList)
	fmt.Println(studentList)
	db.Delete(&studentList)
	fmt.Println(studentList)

}
