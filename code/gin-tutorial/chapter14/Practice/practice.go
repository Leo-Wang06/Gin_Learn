package practice

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("[%s] %s %s -> %d (%v)\n",
			time.Now().Format("15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Token") == "" {
			c.AbortWithStatusPureJSON(http.StatusUnauthorized, gin.H{
				"error": "缺少 X-Token",
			})
			return 
		}
		c.Next()
	}
}

func Practice() *gin.Engine {
	r := gin.Default()

	r.Use(TraceMiddleware())

	return  r
}