package message

import "daveslist/internal/core/port"

type Config struct {
	MessageRepo port.MessageRepository
}

type Service struct {
	messageRepo port.MessageRepository
}

func New(cfg *Config) port.MessageService {
	return &Service{
		messageRepo: cfg.MessageRepo,
	}
}
