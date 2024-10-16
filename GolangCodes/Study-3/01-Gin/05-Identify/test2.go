package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"` //用户名
	Age  int    `json:"age" binding:"required"  msg:"请输入年龄"`        //年龄
}

// 获取结构体中的msg参数
func _GetValidMsg(err error, obj any) string {
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

func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"khs", "khx", "hsh"}

	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}

	return true
}
func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		ctx.JSON(200, gin.H{"msg": user})
	})

	router.Run(":8080")
}
