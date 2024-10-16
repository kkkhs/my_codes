package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/download", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/octet-stream")                //设置响应头为文件流
		ctx.Header("Content-Disposition", "attachment; filename="+"good.png") //指定下载的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")                     //传输过程中的编码形式
		ctx.File("../uploads/12.png")
	})

	router.Run(":8080")
}
