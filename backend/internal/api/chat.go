package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"likemind-backend/internal/services"
)

// RegisterChatRoutes provides chat session and messaging endpoints
func RegisterChatRoutes(rg *gin.RouterGroup, chat *services.ChatService) {
	rg.GET("/sessions", func(c *gin.Context) {
		uid, _ := c.Get("user_id")
		userID := uint(uid.(float64))
		sessions, err := chat.GetUserSessions(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sessions)
	})

	rg.POST("/sessions", func(c *gin.Context) {
		uid, _ := c.Get("user_id")
		userID := uint(uid.(float64))
		var payload struct {
			Title string `json:"title"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		session, err := chat.CreateSession(c.Request.Context(), userID, payload.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, session)
	})

	rg.GET("/sessions/:id/messages", func(c *gin.Context) {
		sid, _ := strconv.Atoi(c.Param("id"))
		msgs, err := chat.GetSessionMessages(c.Request.Context(), uint(sid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, msgs)
	})

	rg.POST("/sessions/:id/messages", func(c *gin.Context) {
		sid, _ := strconv.Atoi(c.Param("id"))
		var payload struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg, err := chat.SendMessage(c.Request.Context(), uint(sid), payload.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, msg)
	})
}
