package main

import (
	"log"
	"os"

	"likemind-backend/internal/api"
	"likemind-backend/internal/config"
	"likemind-backend/internal/database"
	"likemind-backend/internal/middleware"
	"likemind-backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize database connections
	db, err := database.InitPostgreSQL(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	redisClient, err := database.InitRedis(cfg.RedisURL)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// Initialize services
	userService := services.NewUserService(db)
	authService := services.NewAuthService(db, cfg.JWTSecret)
	aiService := services.NewAIService(cfg.OpenAIAPIKey)
	searchService := services.NewSearchService(cfg.VectorDBURL)
	chatService := services.NewChatService(aiService, redisClient)

	// Initialize Gin router
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// API routes
	apiV1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := apiV1.Group("/auth")
		api.RegisterAuthRoutes(auth, authService, userService)

		// Protected routes
		protected := apiV1.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			// User routes
			api.RegisterUserRoutes(protected.Group("/users"), userService)

			// AI routes
			api.RegisterAIRoutes(protected.Group("/ai"), aiService)

			// Chat routes
			api.RegisterChatRoutes(protected.Group("/chat"), chatService)

			// Search routes
			api.RegisterSearchRoutes(protected.Group("/search"), searchService)

			// Knowledge routes
			api.RegisterKnowledgeRoutes(protected.Group("/knowledge"), searchService)
		}

		// WebSocket routes
		api.RegisterWebSocketRoutes(apiV1.Group("/ws"), chatService)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
