package routes

import (
	"gin-tutorial/chapter04/handlers"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/:userID", func(ctx *gin.Context) {
			userID := ctx.Param("userID")
			handlers.GetUsers(userID, ctx)

		})
		user := users.Group("/:userID")
		{
			user.GET("/articles", handlers.GetArticles)
			article := user.Group("/:articleID")
			{
				article.GET("/comment",func(ctx *gin.Context) {
					articleID := ctx.Param("articleID")
					handlers.GetComments(articleID, ctx)
				})
			}
		}
	}
	return r
}