package chapter22

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func LearnCode() {
	r := gin.Default()

	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filepath := "./uploads/" + filename

		// 检查文件是否存在
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			c.JSON(404, gin.H{"error": "文件不存在"})
			return
		}

		// 方式一：自动识别 Content-Type
		c.File(filepath)

		// 方式二：指定下载文件名
		// c.FileAttachment(filepath, filename)
	})

	r.GET("/file/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		action := c.DefaultQuery("action", "inline")

		if action == "download" {
			c.Header("Content-Disposition", "attachment; filename=" + filename)
		} else {
			c.Header("Content-Disposition", "inline")
		}
		c.File("./uploads/" + filename)
	})

	r.GET("/video/:filename", func(c *gin.Context) {
		filepath := "./uploads/" + c.Param("filename")

		// rangeHeader := c.GetHeader("Range")

		http.ServeFile(c.Writer, c.Request, filepath)
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}