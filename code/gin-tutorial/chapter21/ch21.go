package chapter21

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func LearnCode() {
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

	r.GET("/get-cookie", func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.String(200, "未找到 Cookie")
			return
		}
		c.String(200, "用户名：%s", username)
	})

	store := cookie.NewStore([]byte("your-secret-key-32-bytes!!!"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "admin" && password == "123456" {
			session := sessions.Default(c)
			session.Set("user_id", 1)
			session.Set("username", username)
			session.Save()
			c.JSON(200, gin.H{"msg": "登录成功"})
			return
		}
		c.JSON(401, gin.H{"error": "用户名或密码错误"})
	})

	r.GET("/profile", func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			c.JSON(401, gin.H{"error": "请先登录"})
			return
		}
		c.JSON(200, gin.H{
			"user_id": userID,
			"username": session.Get("username"),
		})
	})

	r.POST("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{"msg": "已登出"})
	})	

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}