package chapter08

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "张三",
		})
	})

	// 上面都没匹配到 → 到这里
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": fmt.Sprintf("路径 %s 不存在", c.Request.URL.Path),
		})
	})
}