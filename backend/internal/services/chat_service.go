package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"likemind-backend/internal/models"
)

type ChatService struct {
	aiService   *AIService
	redisClient *redis.Client
	db          *gorm.DB
}

func NewChatService(aiService *AIService, redisClient *redis.Client) *ChatService {
	return &ChatService{
		aiService:   aiService,
		redisClient: redisClient,
	}
}

func (s *ChatService) CreateSession(ctx context.Context, userID uint, title string) (*models.ChatSession, error) {
	session := &models.ChatSession{
		UserID:   userID,
		Title:    title,
		IsActive: true,
	}

	if err := s.db.Create(session).Error; err != nil {
		return nil, fmt.Errorf("failed to create chat session: %w", err)
	}

	return session, nil
}

func (s *ChatService) GetUserSessions(ctx context.Context, userID uint) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	if err := s.db.Where("user_id = ? AND is_active = ?", userID, true).
		Order("updated_at DESC").
		Find(&sessions).Error; err != nil {
		return nil, fmt.Errorf("failed to get user sessions: %w", err)
	}

	return sessions, nil
}

func (s *ChatService) GetSessionMessages(ctx context.Context, sessionID uint) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	if err := s.db.Where("session_id = ?", sessionID).
		Order("created_at ASC").
		Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("failed to get session messages: %w", err)
	}

	return messages, nil
}

func (s *ChatService) SendMessage(ctx context.Context, sessionID uint, userMessage string) (*models.ChatMessage, error) {
	// Create user message
	userMsg := &models.ChatMessage{
		SessionID: sessionID,
		Role:      "user",
		Content:   userMessage,
	}

	if err := s.db.Create(userMsg).Error; err != nil {
		return nil, fmt.Errorf("failed to save user message: %w", err)
	}

	// Get conversation history
	messages, err := s.GetSessionMessages(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation history: %w", err)
	}

	// Generate AI response
	aiResponse, err := s.aiService.GenerateResponse(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("failed to generate AI response: %w", err)
	}

	// Save AI response
	aiResponse.SessionID = sessionID
	if err := s.db.Create(aiResponse).Error; err != nil {
		return nil, fmt.Errorf("failed to save AI response: %w", err)
	}

	// Cache recent conversation in Redis
	s.cacheConversation(ctx, sessionID, messages)

	return aiResponse, nil
}

func (s *ChatService) cacheConversation(ctx context.Context, sessionID uint, messages []models.ChatMessage) {
	cacheKey := fmt.Sprintf("chat:session:%d", sessionID)
	
	// Keep only the last 20 messages
	if len(messages) > 20 {
		messages = messages[len(messages)-20:]
	}

	data, err := json.Marshal(messages)
	if err != nil {
		return
	}

	s.redisClient.Set(ctx, cacheKey, data, time.Hour)
}

func (s *ChatService) GetCachedConversation(ctx context.Context, sessionID uint) ([]models.ChatMessage, error) {
	cacheKey := fmt.Sprintf("chat:session:%d", sessionID)
	
	data, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}

	var messages []models.ChatMessage
	if err := json.Unmarshal([]byte(data), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *ChatService) DeleteSession(ctx context.Context, sessionID uint, userID uint) error {
	result := s.db.Where("id = ? AND user_id = ?", sessionID, userID).
		Update("is_active", false)
	
	if result.Error != nil {
		return fmt.Errorf("failed to delete session: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("session not found or access denied")
	}

	// Clear cache
	cacheKey := fmt.Sprintf("chat:session:%d", sessionID)
	s.redisClient.Del(ctx, cacheKey)

	return nil
}
