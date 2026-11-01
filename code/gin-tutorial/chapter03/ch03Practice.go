package chapter03

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Practice() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"users": []gin.H{
				{"id":1, "name": "张三"},
				{"id":2, "name": "李四"},
			},
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
			"name": fmt.Sprintf("用户%s", id),
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{
			"id": 3,
			"message": "用户创建成功",
		})
	})

	r.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s完整更新", id),
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s部分更新", id),
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s已删除", id),
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}