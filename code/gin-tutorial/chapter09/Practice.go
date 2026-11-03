package chapter09

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Practice() *gin.Engine {
	r := gin.Default()

	r.GET("/old", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/new")
	})

	r.GET("/s/abc", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/articles/123"
		r.HandleContext(ctx)
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/dashboard")
	})

	return r

}