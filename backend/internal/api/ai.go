package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"likemind-backend/internal/models"
	"likemind-backend/internal/services"
)

// RegisterAIRoutes exposes a simple AI generation endpoint
func RegisterAIRoutes(rg *gin.RouterGroup, ai *services.AIService) {
	rg.POST("/generate", func(c *gin.Context) {
		var payload struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg := []models.ChatMessage{{Role: "user", Content: payload.Message}}
		resp, err := ai.GenerateResponse(c.Request.Context(), msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
