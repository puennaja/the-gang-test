package dto

import "time"

type MessageQuery struct {
	PaginationQuery
	SenderID string `validate:"required"`
}

type CreateMessageRequest struct {
	SenderID    string `json:"-" validate:"required"`
	SenderName  string `json:"-" validate:"required"`
	ReciverID   string `json:"reciver_id" validate:"required"`
	ReciverName string `json:"reciver_name" validate:"required"`
	Message     string `json:"message" validate:"required"`
}

type MessageResponse struct {
	SenderID    string    `json:"sender_id"`
	SenderName  string    `json:"sender_name"`
	ReciverID   string    `json:"reciver_id"`
	ReciverName string    `json:"reciver_name"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

type MessageListResponse []MessageResponse
