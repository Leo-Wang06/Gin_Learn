package chapter21

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func HaveSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessions := sessions.Default(c)
		if sessions.Get("user_id") == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
			return

		}
		c.Next()
	}
}

func Practice() *gin.Engine {
	r := gin.Default()

	r.GET("/set-cookie", func(c *gin.Context) {
		c.SetCookie(
			"username",
			"zhangsan",
			3600,
			"/",
			"localhost",
			false,
			true,
		)
		c.String(200, "Cookie 已设置")
	})

	store := cookie.NewStore([]byte("your-secret-key-32-bytes!!!"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(HaveSession())

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "admin" && password == "123456" {
			sessions := sessions.Default(c)
			sessions.Set("user_id", 1)
			sessions.Set("username", username)
			sessions.Save()
			c.JSON(200, gin.H{"msg": "登录成功"})
			return 
		}
		c.JSON(401, gin.H{"error": "用户名或密码错误"})
	})

	r.GET("/profile", func(c *gin.Context) {
		sessions := sessions.Default(c)
		userID := sessions.Get("user_id")
		if userID == nil {
			c.JSON(401, gin.H{"error": "请先登录"})
			return 
		}
		c.JSON(200, gin.H{
			"user_id": userID,
			"username": sessions.Get("username"),
		})
	})

	r.POST("/logout", func(c *gin.Context) {
		sessions := sessions.Default(c)
		sessions.Clear()
		sessions.Save()
		c.JSON(200, gin.H{"msg": "已登出"})
	})

	return r
}