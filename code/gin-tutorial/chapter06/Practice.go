package chapter06

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Practice() *gin.Engine {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "用户名和密码不能为空",
			})
			return 
		}

		if username == "admin" && password == "123455" {
			c.JSON(http.StatusOK, gin.H{"msg": "登陆成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名和密码错误"})
		}
	})

	r.POST("/webhook", func(c *gin.Context) {
		body, _ := c.GetRawData()
		contentType := c.ContentType()

		fmt.Printf("Content-Type: %s\n", contentType)
		fmt.Printf("Body: %s\n", string(body))

		c.JSON(http.StatusOK, gin.H{"msg": "收到"})
	})

	return r
}