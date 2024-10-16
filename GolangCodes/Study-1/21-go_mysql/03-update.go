package main

import (
	"fmt"
	"log"
)

func update(username string, id int) {
	ret, err := DB.Exec("update user set username = ? where user_id = ?", username, id)
	if err != nil {
		log.Println("更新出现问题:", err)
		return
	}
	affected, err := ret.RowsAffected()
	fmt.Println("更新成功的行数 :", affected)
}
