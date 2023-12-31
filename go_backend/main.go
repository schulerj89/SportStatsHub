package main

import (
	"github.com/gin-gonic/gin"
	// Other imports...
)

func main() {
	router := gin.Default()

	// Set up routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	router.Run(":8000")
}
