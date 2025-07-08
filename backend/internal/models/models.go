package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	Role      string         `json:"role" gorm:"default:user"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// ChatSession represents a chat session
type ChatSession struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Title     string         `json:"title"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Messages  []ChatMessage  `json:"messages,omitempty" gorm:"foreignKey:SessionID"`
}

// ChatMessage represents a message in a chat session
type ChatMessage struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	SessionID uint           `json:"session_id" gorm:"not null"`
	Session   ChatSession    `json:"-" gorm:"foreignKey:SessionID"`
	Role      string         `json:"role" gorm:"not null"` // user, assistant, system
	Content   string         `json:"content" gorm:"type:text;not null"`
	Metadata  string         `json:"metadata,omitempty" gorm:"type:jsonb"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// KnowledgeDocument represents a document in the knowledge base
type KnowledgeDocument struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Title       string         `json:"title" gorm:"not null"`
	Content     string         `json:"content" gorm:"type:text;not null"`
	Source      string         `json:"source"`
	DocumentType string        `json:"document_type"`
	Metadata    string         `json:"metadata,omitempty" gorm:"type:jsonb"`
	EmbeddingID string         `json:"embedding_id,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// SearchQuery represents a search query log
type SearchQuery struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
	Query     string         `json:"query" gorm:"not null"`
	Results   string         `json:"results,omitempty" gorm:"type:jsonb"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Agent represents an AI agent configuration
type Agent struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Config      string         `json:"config" gorm:"type:jsonb"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
