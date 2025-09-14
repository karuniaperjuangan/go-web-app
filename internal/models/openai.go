package models

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

type Message struct {
	gorm.Model
	ID             uint
	ConversationID uint
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Role           string
	Content        string
}

type OpenAIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type OpenAIResponse struct {
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type SimpleChatRequest struct {
	Message string `json:"message"`
}
