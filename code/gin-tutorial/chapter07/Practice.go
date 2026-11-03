package chapter07

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Practice() *gin.Engine {
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

	r.POST("/api", func(c *gin.Context) {
		ct := c.ContentType()

		ct2 := c.GetHeader("content-Type")

		if ct == "application/json" {

		} else if ct == "application/x-www-form-urlencoded" {

		}

		c.JSON(200, gin.H{
			"detected_type": ct,
			"raw_header":    ct2,
		})
	})

	return r
}