package chapter03

import (
	// "fmt"
	// "net/http"

	// "github.com/gin-gonic/gin"
)

func LearnCode() {
	// r := gin.Default()

	// 这里的路径是相对于运行时的目录来匹配的，不是按照相对目录来的。
	// r.LoadHTMLGlob("./templates/*.html")
	// r.GET("/welcome", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"name" : "张三"})
	// })

	// r.POST("/users", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusCreated, gin.H{
	// 		"message": "用户创建成功",
	// 		"id": 1,
	// 	})
	// })

	// r.PUT("/users/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("用户 %s 已完整更新", id),
	// 	})
	// })

	// r.PATCH("/users/:id",func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("用户 %s 已部分更新", id),
	// 	})
	// })

	// r.DELETE("/users/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("用户 %s 已删除", id),
	// 	})
	// })

	// if err := r.Run(":8080"); err != nil {
	// 	panic(err)
	// }

	Practice()
}