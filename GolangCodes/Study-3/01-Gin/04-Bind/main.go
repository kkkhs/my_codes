package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindJSON(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.POST("/query", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindQuery(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.POST("/uri/:name/:age/:sex", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindUri(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.POST("/form", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBind(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.Run(":8080")
}
