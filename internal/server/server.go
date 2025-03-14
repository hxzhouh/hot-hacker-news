package main

import (
	"hot-hacker-new/internal/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("internal/views/**/*")

	// Set up middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
