package chapter06

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		remember := ctx.DefaultPostForm("remember", "false")

		ctx.JSON(200, gin.H{
			"username": username,
			"password": password,
			"remember": remember,
		})
	})

	r.POST("/webhook", func(ctx *gin.Context) {
		body, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(400, gin.H{"error": "读取失败"})
			return
		}

		fmt.Printf("收到 webhook: %s\n", string(body))

		ctx.JSON(200, gin.H{"msg": "收到"})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}