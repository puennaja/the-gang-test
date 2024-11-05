package message_test

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/mocks"
	"daveslist/internal/core/port"
	"daveslist/internal/core/service/message"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx             context.Context
	service         port.MessageService
	mockMessageRepo *mocks.MockMessageRepository
}

func (s *TestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.ctx = context.Background()
	s.mockMessageRepo = mocks.NewMockMessageRepository(ctrl)
	s.service = message.New(&message.Config{
		MessageRepo: s.mockMessageRepo,
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreateMessage() {

	s.T().Run("success", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(&dto.MessageResponse{}, nil)
		_, err := s.service.CreateMessage(s.ctx, &dto.CreateMessageRequest{})

		s.Assert().NoError(err)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.CreateMessage(s.ctx, &dto.CreateMessageRequest{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestGetMessageList() {

	s.T().Run("success", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockMessageRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(dto.MessageListResponse{}, nil)
		_, err := s.service.GetMessageList(s.ctx, &dto.MessageQuery{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: count documents mongo error", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), errors.New("error"))
		_, err := s.service.GetMessageList(s.ctx, &dto.MessageQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: message not found", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), nil)
		_, err := s.service.GetMessageList(s.ctx, &dto.MessageQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: mongo error", func(t *testing.T) {
		s.mockMessageRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockMessageRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.GetMessageList(s.ctx, &dto.MessageQuery{})

		s.Assert().Error(err)
	})
}
