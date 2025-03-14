package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger is a middleware that logs the incoming requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log the request details
		duration := time.Since(start)
		log.Printf("Request: %s %s | Duration: %v | Status: %d",
			c.Request.Method, c.Request.URL.Path, duration, c.Writer.Status())
	}
}