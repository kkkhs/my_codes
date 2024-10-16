package main

import (
	"errors"
	"log"
)

type User struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func query(id int) (*User, error) {
	rows, err := DB.Query("select * from user where user_id = ? limit 1", id)
	if err != nil {
		log.Println("查询出现错误:", err)
		return nil, errors.New(err.Error())
	}
	user := new(User)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.Username, &user.Sex, &user.Email); err != nil {
			log.Println("scan err:", err)
			return nil, errors.New(err.Error())
		}
	}
	return user, err
}
