package main

import (
	"fmt"
	"log"
)

func delete(id int) {
	ret, err := DB.Exec("delete from user where user_id = ?", id)
	if err != nil {
		log.Println("删除出现问题：", err)
	}

	affected, _ := ret.RowsAffected()
	fmt.Println("删除成功的行数:", affected)

}
