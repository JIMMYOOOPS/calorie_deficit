package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/", func(context *gin.Context) {
		// Respond with a JSON message
		context.JSON(200, Response{Message: "Hello, World!"})
	})

	// Start the server on port 8080
	router.Run("127.0.0.1:3000")
}
