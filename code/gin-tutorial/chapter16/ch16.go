package chapter16

import (
	"fmt"
	"uuid"

	"github.com/gin-gonic/gin"
)

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先从请求头获取，没有就生成
		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		
		// 存到 Context（后续 Handler 和中间件都能取到）
		c.Set("trace_id", traceID)
		
		// 返回给客户端
		c.Header("X-Trace-ID", traceID)
		
		c.Next()
	}
}

func handler(c *gin.Context) {
	traceID, _ := c.Get("trace_id")
	fmt.Println("TraceID:", traceID)
}

func LearnCode() {
	r := gin.Default()



	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}