package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerWithFormatter(params gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	statusColor = params.StatusCodeColor()
	methodColor = params.MethodColor()
	resetColor = params.ResetColor()
	return fmt.Sprintf(
		"[ feng ] %s  | %s %d  %s | \t %s | %s | %s %-7s %s \t  %s\n",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor,
		params.ClientIP,
		params.Latency,
		methodColor, params.Method, resetColor,
		params.Path,
	)
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(LoggerWithFormatter))
	router.GET("/index", func(ctx *gin.Context) {})

	router.Run()
}
