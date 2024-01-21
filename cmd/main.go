package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"line-man.com/tin-dpj/pkg/api"
)

func main() {
	// Initialize a new Gin router.
	router := gin.Default()

	// Define routes and handlers.
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}).GET("/covid/summary", api.CovidSummaryHandler)

	// Start the server.
	port := 8080 // You can change this to your desired port.
	fmt.Printf("Server is running on :%d...\n", port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
