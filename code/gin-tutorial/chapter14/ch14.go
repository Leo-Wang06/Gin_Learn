package chapter14

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)  {
		fmt.Println("① 中间件前置：请求到达")

		c.Next()

		fmt.Println("③ 中间件后置：响应返回")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			fmt.Println("认证失败，中止请求")
			c.AbortWithStatusPureJSON(401, gin.H{"error": "未认证"})
			return 
		}
		c.Next()
	}
}

func LearnCode() {
	r := gin.Default()

	r.GET("/test", MyMiddleware(), func(c *gin.Context) {
		fmt.Println("② 处理函数执行")
		c.String(200, "done")
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}