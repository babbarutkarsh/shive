package main

import (
	"os"
	"shive/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()

	// run database
	database.StartDB()

	// log events
	router.Use(gin.Logger())

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "Welcome to shive api"})
	})
	router.Run(":" + port)
}
