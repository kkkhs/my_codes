package main

import (
	"khs/Study-3/01-Gin/09-Logrus/6-gin_logrus/log"
	"khs/Study-3/01-Gin/09-Logrus/6-gin_logrus/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	log.InitFile("./logs", "server")
	router := gin.New()
	router.Use(middleware.LogMiddleware())

	router.GET("/", func(c *gin.Context) {
		logrus.Info("来了")
		c.JSON(200, gin.H{"msg": "你好"})
	})
	router.Run(":8080")

}
