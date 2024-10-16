package main

import (
	"fmt"
	"log"
)

func insertTX(username string) {
	tx, err := DB.Begin()
	if err != nil {
		log.Println("开启事务错误：", err)
		return
	}
	res, err := tx.Exec("insert into user (username, sex, email) value(?, ?, ?)", username, "man", "test@qq.com")
	tx.Exec("insert into user (username, sex, email) value(?, ?, ?)", username, "woman", "test2@qq.com")
	if err != nil {
		log.Println("事务sql执行错误:", err)
	}
	id, _ := res.LastInsertId()
	fmt.Println("插入成功 :", id)
	if username == " khs" {
		fmt.Println("回滚....")
		_ = tx.Rollback()
	} else {
		fmt.Println("提交..")
		_ = tx.Commit()
	}

}
