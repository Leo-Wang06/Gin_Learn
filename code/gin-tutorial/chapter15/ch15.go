package chapter15

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams)string {
		
		return fmt.Sprintf(
			"%s | %3d | %10v | %-15s | %-6s %-20s | UA: %s\n",
			params.TimeStamp.Format("2006-01-02 15:04:05"),
			params.StatusCode,
			params.Latency,
			params.ClientIP,                    // 客户端 IP
			params.Method,
			params.Path,
			params.Request.UserAgent(),         // User-Agent
		)
	}))
}