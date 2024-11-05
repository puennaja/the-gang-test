package message

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"

	errors "daveslist/pkg/go-errors"
)

func (s *Service) CreateMessage(ctx context.Context, data *dto.CreateMessageRequest) (*dto.MessageResponse, error) {
	res, err := s.messageRepo.Insert(ctx, new(model.Message).FromDTO(data))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetMessageList(ctx context.Context, query *dto.MessageQuery) (*dto.PaginationResponse, error) {
	count, err := s.messageRepo.CountByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrMessageNotFound
	}

	res, err := s.messageRepo.FindByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.NewPaginationResponse(query.Page, query.Limit, count, res), nil
}
