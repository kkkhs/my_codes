package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserInfo struct {
	Name       string `json:"name" binding:"min=4,max=8"`             //用户名
	Age        int    `json:"age" binding:"gt=18,lt=30"`              //年龄
	Sex        string `json:"sex" binding:"oneof=man woman"`          //性别
	Password   string `json:"password"`                               //密码
	RePassword string `json:"re_password" binding:"eqfield=Password"` //确认密码
}

// 获取结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用时，需要穿obj的指针
	getObj := reflect.TypeOf(obj)
	//1、将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//2、断言成功,循环每一个错误信息
		for _, e := range errs {
			//根据报错字段名 获取结构体的具体字段
			if field, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := field.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
func main0() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var user UserInfo
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"msg": user})
	})

	router.POST("/mistake", func(ctx *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"用户名校验失败"` //用户名
			Age  int    `json:"age" binding:"required"  msg:"请输入年龄"`   //年龄
		}
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
			return
		}
		ctx.JSON(200, gin.H{"msg": user})
	})

	router.Run(":8080")
}
