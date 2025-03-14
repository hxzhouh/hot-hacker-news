package routes

import (
	"hot-hacker-new/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.NewNewsHandler.Index)
	r.GET("/news/:id", handlers.NewsDetailHandler)
}
