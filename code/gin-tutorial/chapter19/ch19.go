package chapter19

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/layout/*")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"message": "欢迎！",
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}