package model

import (
	"daveslist/internal/core/domain/dto"
	"daveslist/pkg/utils"
	"time"
)

type Message struct {
	ID          string    `bson:"id"`
	SenderID    string    `bson:"sender_id"`
	SenderName  string    `bson:"sender_name"`
	ReciverID   string    `bson:"reciver_id"`
	ReciverName string    `bson:"reciver_name"`
	Message     string    `bson:"message"`
	CreatedAt   time.Time `bson:"created_at"`
}

func (m *Message) FromDTO(dto *dto.CreateMessageRequest) *Message {
	return &Message{
		ID:          utils.GetUUID(),
		SenderID:    dto.SenderID,
		SenderName:  dto.SenderName,
		ReciverID:   dto.ReciverID,
		ReciverName: dto.ReciverName,
		Message:     dto.Message,
		CreatedAt:   time.Now(),
	}
}

func (m *Message) ToDTO() *dto.MessageResponse {
	return &dto.MessageResponse{
		SenderID:    m.SenderID,
		SenderName:  m.SenderName,
		ReciverID:   m.ReciverID,
		ReciverName: m.ReciverName,
		Message:     m.Message,
		CreatedAt:   m.CreatedAt,
	}
}

type MessageList []Message

func (m MessageList) ToDTO() dto.MessageListResponse {
	res := make(dto.MessageListResponse, 0)
	for _, v := range m {
		res = append(res, *v.ToDTO())
	}
	return res
}
