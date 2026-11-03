package chapter10

import "github.com/gin-gonic/gin"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email"    binding:"required,email"`
	Age      int    `json:"age"`
}

func LearnCode() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(201, gin.H{
			"username": req.Username,
			"email": req.Email,
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}