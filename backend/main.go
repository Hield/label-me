package backend

import (
	"log"
	"os"

	"github.com/Hield/label-me/backend/global"
	"github.com/Hield/label-me/backend/handlers"
	"github.com/gin-gonic/gin"
)

// Main function of the application
func Main() {
	port := os.Getenv("PORT")
	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	// Setup the database
	if err := global.SetupDB(); err != nil {
		log.Println(err)
		log.Fatal("Error database")
	}

	// Create the router
	router := gin.New()
	router.Use(gin.Logger())
	router.StaticFile("/", "index.html")
	router.StaticFile("/admin", "admin.html")
	router.Static("/app", "frontend")

	router.GET("/api/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "holahola",
		})
	})

	router.POST("/api/sentences", handlers.AddSentence)

	router.Run(":" + port)
}
