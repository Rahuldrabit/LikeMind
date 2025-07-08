package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"likemind-backend/internal/models"
)

type AIService struct {
	apiKey      string
	httpClient  *http.Client
	baseURL     string
}

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
	Delta   Message `json:"delta,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func NewAIService(apiKey string) *AIService {
	return &AIService{
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    "https://api.openai.com/v1",
	}
}

func (s *AIService) GenerateResponse(ctx context.Context, messages []models.ChatMessage) (*models.ChatMessage, error) {
	if s.apiKey == "" {
		return s.generateMockResponse(messages)
	}

	// Convert messages to OpenAI format
	openAIMessages := make([]Message, len(messages))
	for i, msg := range messages {
		openAIMessages[i] = Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	request := OpenAIRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    openAIMessages,
		Temperature: 0.7,
		MaxTokens:   1000,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.baseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned from OpenAI")
	}

	response := &models.ChatMessage{
		Role:      "assistant",
		Content:   openAIResp.Choices[0].Message.Content,
		CreatedAt: time.Now(),
	}

	return response, nil
}

func (s *AIService) generateMockResponse(messages []models.ChatMessage) (*models.ChatMessage, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages provided")
	}

	lastMessage := messages[len(messages)-1]
	
	// Generate a mock response based on the last message
	mockResponses := []string{
		"I understand your question about " + lastMessage.Content + ". Let me provide you with a comprehensive answer.",
		"That's an interesting point. Based on my knowledge, I can help you with that topic.",
		"I can assist you with that. Here's what I know about " + lastMessage.Content,
		"Thank you for your question. Let me break this down for you.",
	}

	response := &models.ChatMessage{
		Role:      "assistant",
		Content:   mockResponses[len(messages)%len(mockResponses)],
		CreatedAt: time.Now(),
	}

	return response, nil
}

func (s *AIService) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	// This would integrate with OpenAI's embedding API
	// For now, return a mock embedding
	return make([]float32, 1536), nil
}

func (s *AIService) AnalyzeText(ctx context.Context, text string) (map[string]interface{}, error) {
	// Text analysis functionality
	return map[string]interface{}{
		"sentiment": "positive",
		"topics":    []string{"technology", "AI"},
		"entities":  []string{"OpenAI", "GPT"},
	}, nil
}
