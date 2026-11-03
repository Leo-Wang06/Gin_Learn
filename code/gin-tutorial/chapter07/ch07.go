package chapter07

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.GET("/headers", func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		contentType := c.GetHeader("Content-Type")

		userAgent := c.Request.Header.Get("User-Agent")

		fmt.Println("=== 所有请求头 ===")
		for key, values := range c.Request.Header {
			for _, v := range values {
				fmt.Printf("%s: %s\n", key, v)
			}
		}

		c.JSON(200, gin.H{
			"authorization": auth,
			"content_type": contentType,
			"user_agent": userAgent,
		})
	})

	r.GET("/ip", func(c *gin.Context) {
		clientIP := c.ClientIP()

		remoteAddr := c.Request.RemoteAddr

		c.JSON(200, gin.H{
			"client_ip": clientIP,
			"remote_addr": remoteAddr,
		})
	})

	r.GET("/download", func(c *gin.Context) {
		c.Header("Content-Disposition", "attachment; filename=report.pdf")
		c.Header("X-Custom-Header", "hello-word")

		c.Header("X-Custom-Header", "new-value")

		c.File("./report.pdf")
	})

	r.POST("/api", func(c *gin.Context) {
		ct := c.ContentType()

		ct2 := c.GetHeader("Content-Type")

		if ct == "application/json" {
			// 用 ShodleBindSJON
		} else if ct == "application/x-www-from-urlencoded" {
			// 用 PostForm
		}

		c.JSON(200, gin.H{
			"detected_type": ct,
			"raw_header": ct2,
		})

	})

	// 预检请求处理
	r.OPTIONS("/api/data", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(http.StatusNoContent)
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}