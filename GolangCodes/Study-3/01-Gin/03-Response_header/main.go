package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/res", func(ctx *gin.Context) {
		//设置响应头
		ctx.Header("Token", "khs")
		ctx.Header("Content-Type", "application/text; charset=utf-8")

		ctx.JSON(0, gin.H{"data": "看看响应头"})
	})

	router.Run(":8080")
}
