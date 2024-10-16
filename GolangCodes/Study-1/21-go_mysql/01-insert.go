package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/khsbase?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)

	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		db.Close()
		panic(err)
	}
	DB = db
}

func save() {
	r, err := DB.Exec("insert into user (username, sex, email) values(?, ?, ?)", "khs", "man", "khs@qq.com")
	if err != nil {
		fmt.Println("数据插入错误")
		panic(err)
	}
	id, err := r.LastInsertId() //返回最后插入行的ID或序列值
	if err != nil {
		panic(err)
	}
	fmt.Println("数据插入成功：", id)

}

func main() {
	defer DB.Close()
	//save()
	user, _ := query(4)
	fmt.Printf("id = 4 : %+v\n", *user)
	//update("hhh", 4)
	//delete(4)
	insertTX("kkkk")
}
