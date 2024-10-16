package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义一个中间件
func m1(ctx *gin.Context) {
	fmt.Println("1")
}

func m2(ctx *gin.Context) {
	fmt.Println("2")

}

func main() {
	router := gin.Default()

	router.GET("/", m1, func(ctx *gin.Context) {
		fmt.Println("index")
		ctx.JSON(200, gin.H{"msg": "index"})
	}, m2)
	router.Run()
}
