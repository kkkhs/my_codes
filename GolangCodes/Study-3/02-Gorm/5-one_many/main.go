package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Age      int
	Gender   bool
	UserInfo *UserInfo // 通过UserInfo可以拿到用户详情信息
}

type UserInfo struct {
	UserID uint // 外键
	User   User
	ID     uint
	Addr   string
	Like   string
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

	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接失败")
	}

	//DB.AutoMigrate(&User{}, &UserInfo{})

	// DB.Create(&User{
	// 	Name:     "khs",
	// 	Age:      18,
	// 	Gender:   true,
	// 	UserInfo: UserInfo{},
	// })

	var user User
	DB.Take(&user)
	DB.Select("UserInfo").Delete(&user)

}
