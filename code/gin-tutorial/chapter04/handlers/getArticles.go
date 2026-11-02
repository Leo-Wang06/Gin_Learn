package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"article": "内容"})
}