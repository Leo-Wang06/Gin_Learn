package chapter09

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	// r.GET("/old-page", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/new-page")
	// })
	
	// r.GET("/new-page", func(c *gin.Context) {
	// 	c.String(200, "这是新页面")
	// })

	// r.GET("/short", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/long/detailed/path"
	// 	r.HandleContext(c)
	// })

	// r.GET("/long/detailed/path", func(c *gin.Context) {
	// 	c.String(200, "这是长路径的处理结果")
	// })

	r.GET("/", func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			redirectURL := "/login?redirect=" + c.Request.URL.Path
			c.Redirect(http.StatusFound, redirectURL)
			return
		}
		c.String(200, "欢迎回来！")
	})

	r.GET("/login", func(c *gin.Context) {
		redirectTo := c.DefaultQuery("redirect", "/")
		c.String(200, "登录页面（成功后跳转到: %s）", redirectTo)
	})

	r.GET("/dashboard", func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.Redirect(http.StatusFound, "/login?redirect=/dashboard")
			return
		}
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}