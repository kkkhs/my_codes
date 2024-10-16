package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//单文件上传
	router.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		readerFile, _ := file.Open()
		writerFile, _ := os.Create("../uploads/1.png")
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)
		fmt.Println(n)

		ctx.JSON(200, gin.H{"msg": "上传成功"})
	})

	//多文件上传
	router.POST("/uploads", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			ctx.SaveUploadedFile(file, "../uploads/"+file.Filename)
		}

		ctx.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})

	})

	router.Run(":8080")
}
