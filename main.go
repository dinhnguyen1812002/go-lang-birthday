package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	msg := template.HTML(
		"Gửi tôi, người vẫn luôn nỗ lực để mỉm cười mỗi ngày. <br />" +
			"Chúc bản thân một sinh nhật tràn đầy tình yêu, niềm vui và tất cả những hạnh phúc mà tôi xứng đáng có được. <br />" +
			"Happy new year to me!",
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
