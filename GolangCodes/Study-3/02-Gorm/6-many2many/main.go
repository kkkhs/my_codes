package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags"`
}

type Tag struct {
	ID   uint
	Name string
	// Articles []Article `gorm:"many2many:article_tags"` 不需要反向引用
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
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

	//DB.AutoMigrate(&Article{}, &Tag{}) //自动生成第三张表

	//添加
	// DB.Create(&Article{
	// 	Title: "python",
	// 	Tags: []Tag{
	// 		{
	// 			Name: "python基础",
	// 		},
	// 		{
	// 			Name: "python高阶",
	// 		},
	// 	},
	// })

	//创建文章关联已有标签
	// var tags []Tag
	// DB.Find(&tags, "name in ?", []string{"python高阶"})

	// DB.Create(&Article{
	// 	Title: "Golang",
	// 	Tags:  tags,
	// })

	//查询
	// var article Article
	// DB.Preload("Tags").Take(&article)
	// fmt.Println(article)

	//更新Tags
	//1. 先删除原有的关联
	// var article Article
	// DB.Preload("Tags").Take(&article, 1)
	// //fmt.Println(article)
	// //DB.Model(&article).Association("Tags").Delete(article.Tags)

	// //2. 再添加新的关联Tags
	// var tag Tag
	// DB.Take(&tag, 2)
	// DB.Model(&article).Association("Tags").Append(&tag)

	//全部更新
	// var article Article
	// DB.Preload("Tags").Take(&article, 1)

	// var tag Tag
	// DB.Take(&tag, 1)
	// DB.Model(&article).Association("Tags").Replace(&tag)

	// DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	// err = DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
	// fmt.Println(err)

	var article Article
	DB.Preload("Tags").Take(&article, 2)
	fmt.Println(article)
}
