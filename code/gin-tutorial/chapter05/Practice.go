package chapter05

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Practice() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id/orders", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil || userID < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "用户 id 必须是正整数"})
			return
		}

		status := c.DefaultQuery("status", "all")
		if status != "all" && status != "paid" && status != "unpaid" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "status 只能是 paid 或 unpaid"})
			return
		}
		orders := []gin.H{
			{"order_id": 1001, "user_id": userID, "status": "paid", "amount": 99.0},
			{"order_id": 1002, "user_id": userID, "status": "unpaid", "amount": 45.5},
		}
		if status != "all" {
			filtered := make([]gin.H, 0, len(orders))
			for _, o := range orders {
				if o["status"] == status {
					filtered = append(filtered, o)
				}
			}
			orders = filtered
		}
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"status":  status,
			"count":   len(orders),
			"orders":  orders,
		})
	})

	return r
}