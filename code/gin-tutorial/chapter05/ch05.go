package chapter05

import (
	// "fmt"
	// "net/http"
	// "strconv"

	// "github.com/gin-gonic/gin"
)

func LearnCode() {
	// r := gin.Default()

	// // GET /api/v1/users?keyword=张&status=active&page=1&page_size=10
	// r.GET("/api/v1/users", func(c *gin.Context) {
	// 	keyword := c.DefaultQuery("keyword", "")
	// 	status := c.DefaultQuery("status", "all")
	// 	// strconv.Atoi 是将字符串转为整数
	// 	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	// 	if err != nil || page < 1 {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "page 必须是大于等于 1 的整数"})
	// 		return
	// 	}
	// 	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	// 	if err != nil || pageSize < 1 || pageSize > 100 {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size 必须是 1 到 100 之间的整数"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"keyword": keyword,
	// 		"status": status,
	// 		"page": page,
	// 		"page_size": pageSize,
	// 		"message": fmt.Sprintf("搜索关键词=%s, 状态=%s, 第%d页, 每页%d条", keyword, status, page, pageSize),
	// 	})
	// })

	r := Practice()

	if err := r.Run(); err != nil {
		panic(err)
	}
}