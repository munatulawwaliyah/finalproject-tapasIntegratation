package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"final/controller"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", controller.UploadHandler)

	log.Println("Server started at http://localhost:8080")
	r.Run(":8080")
}
