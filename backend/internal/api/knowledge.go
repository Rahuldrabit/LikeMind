package api

import "github.com/gin-gonic/gin"
import "likemind-backend/internal/services"

// RegisterKnowledgeRoutes is a placeholder
func RegisterKnowledgeRoutes(rg *gin.RouterGroup, _ *services.SearchService) {
	rg.GET("/documents", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "list documents"})
	})
}
