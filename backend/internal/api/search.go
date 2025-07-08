package api

import "github.com/gin-gonic/gin"
import "likemind-backend/internal/services"

// RegisterSearchRoutes provides placeholder search endpoints
func RegisterSearchRoutes(rg *gin.RouterGroup, _ *services.SearchService) {
	rg.GET("/history", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "search history"})
	})
}
