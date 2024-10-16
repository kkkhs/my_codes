package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//1、请求头的各种获取方式
	router.GET("/", func(ctx *gin.Context) {
		//获取单个请求头 (首字母大小写不区分)
		fmt.Println(ctx.GetHeader("User-Agent"))
		fmt.Println(ctx.GetHeader("user-agent"))

		//获取全部请求头
		v := ctx.Request.Header // Map类型
		fmt.Printf("%+v\n = ", v)

		//如果使用Get方法或 .GetHeader方法 可以不区分大小写，并且返回第一个Value
		m := ctx.Request.Header.Get("User-Agent")
		fmt.Printf("%+v\n = ", m)
		//如果是Map的取值方式，应注意大小写，并且返回所有Value
		s := ctx.Request.Header["User-Agent"]
		fmt.Printf("%+v\n = ", s)

		//自定义请求头,用Get方法也不区分大小写
		fmt.Println(ctx.Request.Header.Get("Token"))
		fmt.Println(ctx.Request.Header.Get("token"))

		ctx.JSON(200, gin.H{"msg": "成功"})
	})
	//2、爬虫和用户区别对待
	router.GET("/index", func(ctx *gin.Context) {
		userAgent := ctx.GetHeader("user-agent")
		//1、用正则去匹配

		//2、字符串的包含匹配Contains
		if strings.Contains(userAgent, "python") {
			//爬虫来了
			ctx.JSON(0, gin.H{"data": "你是爬虫"})
			return
		}
		ctx.JSON(0, gin.H{"data": "你是用户"})
	})

	router.Run(":8080")
}
