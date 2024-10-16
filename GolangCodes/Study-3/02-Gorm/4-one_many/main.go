package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"size:4"`
	Name     string    `gorm:"size:8"`
	Articles []Article // 用户拥有的文章列表
}

type Article struct {
	ID     uint   `gorm:"size:4"`
	Title  string `gorm:"size:16"`
	UserID uint   `gorm:"size:4"` // 属于   这里的类型要和引用的外键类型一致，包括大小
	User   User   // 属于
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

	//DB.AutoMigrate(&User{}, &Article{}) //自动建表

	//创建用户，带上文章
	// DB.Create(&User{
	// 	Name: "khs",
	// 	Articles: []Article{
	// 		{
	// 			Title: "python",
	// 		},
	// 		{
	// 			Title: "golang",
	// 		},
	// 	},
	// })

	//创建文章，关联已有用户
	// DB.Create(&Article{
	// 	Title:  "c#",
	// 	UserID: 1,
	// })

	// DB.Create(&Article{
	// 	Title: "java",
	// 	User: User{
	// 		Name: "张三",
	// 	},
	// })

	// var user User
	// DB.Take(&user, 2)
	// DB.Create(&Article{
	// 	Title: "java/3",
	// 	User:  user,
	// })

	// var user User
	// DB.Take(&user, 1)

	// var article Article
	// DB.Take(&article, 6)

	//给用户绑定文章
	//DB.Model(&user).Association("Articles").Append(&article)

	//给文章绑定用户
	//DB.Model(&article).Association("User").Append(&user)

	//普通查询文章
	// var article Article
	// DB.Preload("User").Take(&article) 预加载
	// fmt.Printf("%v\n", article)

	// var user User
	// DB.Preload("Articles.User.Articles").Take(&user) //预加载(可以嵌套)
	// fmt.Println(user)

	// var user User
	// DB.Preload("Articles", "id > ?", 2).Take(&user) //待条件的预加载
	// fmt.Println(user)

	// var user User
	// DB.Preload("Articles", func(db *gorm.DB) *gorm.DB { //自定义预加载
	// 	return db.Where("id > ?", 2)
	// }).Take(&user)
	// fmt.Println(user)

	//清除外键关系
	var article []Article
	DB.Find(&article)
	DB.Delete(&article)
}
