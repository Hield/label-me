package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "holahola",
		})
	})
	router.Run(":" + port)
}
