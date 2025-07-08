package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"likemind-backend/internal/services"
)

// RegisterUserRoutes exposes user related endpoints
func RegisterUserRoutes(rg *gin.RouterGroup, users *services.UserService) {
	rg.GET("/me", func(c *gin.Context) {
		idVal, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		idFloat, ok := idVal.(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
			return
		}
		user, err := users.GetByID(uint(idFloat))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})
}
