package chapter20

import "github.com/gin-gonic/gin"

func LearnCode() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "请选择文件"})
			return
		}

		dst := "./uploads" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{"error": "保存失败"})
			return
		}

		c.JSON(200, gin.H{
			"message": "上传成功",
			"filename": file.Filename,
			"size": file.Size,
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}