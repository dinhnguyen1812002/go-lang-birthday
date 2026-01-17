package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	msg := template.HTML(
		"Gửi chính tôi, người đã luôn nỗ lực mỉm cười suốt 365 ngày qua. <br />" +
			"Mong bản thân ở tuổi này không chỉ già đi theo năm tháng, mà sẽ lớn lên cùng những trải nghiệm. Chúc tôi luôn giữ được sự tỉnh táo để chọn con đường khiến mình tự hào. <br />" +
			"Happy birthday and a brilliant new chapter to me!",
	)

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "birthday-cake.html", gin.H{
			"message": msg,
		})
	})

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
