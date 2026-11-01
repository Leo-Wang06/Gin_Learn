package chapter01

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Practice() {
	r := gin.Default()

	r.GET("/hello", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"msg": "你好，世界",
		})
	})

	if err := r.Run(":9090"); err != nil {
		panic(err)
	}
}