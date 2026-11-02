package chapter04

import (
	"gin-tutorial/chapter04/routes"
)

func LearnCode() {
	// r := gin.Default()
	
	// listUsers := func (c *gin.Context)  {
	// 	c.JSON(http.StatusOK, gin.H{"vsersion": "v1"})
	// }

	// adminListUsers := func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"role": "admin"}) }
	// listUsersV2 := func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"version": "v2"}) }

	// api := r.Group("/api")
	// {
	// 	api.GET("/ping", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
	// 	})
	// 	api.GET("/status", func(ctx *gin.Context) {
	// 		ctx.JSON(http.StatusOK, gin.H{"status" : "ok"})
	// 	})

	// 	if err := r.Run(":8080"); err != nil {
	// 		panic(err)
	// 	}

	// 	v1 := api.Group("/v1")
	// 	{
	// 		v1.GET("/users", listUsers) // /api/v1/users
	// 		admin := v1.Group("/admin")
	// 		{
	// 			admin.GET("/users", adminListUsers) // /api/v1/admin/users
	// 		}
	// 	}

	// 	v2 := api.Group("/v2")
	// 	{
	// 		v2.GET("/users", listUsersV2) // /api/v2/users
	// 	}
	// }

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}