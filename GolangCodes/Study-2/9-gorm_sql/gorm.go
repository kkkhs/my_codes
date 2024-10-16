package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库表映射为结构体 // UserName ----> user_name
type User struct { // 1、User驼峰型结构体名 ------>>> user蛇型表名
	Id      int    //id	   主键
	Keyword string `gorm:"column:keywords"` //keywords
	City    string //city
}

// 2、显式告诉表名
func (User) TableName() string {
	return "user"
}

func main() {
	//连接mysql数据库
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil) //引入gorm
	checkError(err)

	user := User{
		Id:      5858,
		Keyword: "golang",
		City:    "天津",
	}
	client.Create(user)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
