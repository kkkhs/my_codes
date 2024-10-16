package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func m1(ctx *gin.Context) {
	fmt.Println("1")
	//中间件传参
	ctx.Set("name", "khs")
}

func m2(ctx *gin.Context) {
	fmt.Println("2")
}

func main() {
	router := gin.Default()

	//Use使用全局中间件
	router.Use(m1, m2)

	router.GET("/1", func(ctx *gin.Context) {
		fmt.Println(ctx.Get("name"))
		ctx.JSON(200, gin.H{"msg": "index"})
	})

	router.GET("/2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "2"})
	})
	router.Run()
}
