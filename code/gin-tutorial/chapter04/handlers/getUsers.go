package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(userID string,c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": userID,
		"name": fmt.Sprintf("用户%s", userID),
	})
}