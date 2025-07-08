package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"likemind-backend/internal/services"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

// RegisterWebSocketRoutes registers basic websocket chat endpoint
func RegisterWebSocketRoutes(rg *gin.RouterGroup, chat *services.ChatService) {
	rg.GET("/chat", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			chat.SendMessage(c.Request.Context(), 0, string(msg))
			conn.WriteMessage(websocket.TextMessage, msg)
		}
	})
}
