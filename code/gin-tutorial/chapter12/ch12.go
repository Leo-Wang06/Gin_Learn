package chapter12

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		// 告诉客户端/代理这是流式内容，别缓存
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.Header("Cache-Control", "no-cache")
		c.Header("X-Accel-Buffering", "no") // 让 Nginx 不缓冲（部署时有用）

		count := 0
		c.Stream(func(w io.Writer) bool {
			count++
			fmt.Fprintf(w, "第 %d 条消息（%s）\n", count, time.Now().Format("15:04:05"))

			if count >= 5 {
				return false // 发完 5 条，结束
			}

			time.Sleep(time.Second) // 停顿 1 秒，再发下一条
			return true             // 继续下一次 step
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
