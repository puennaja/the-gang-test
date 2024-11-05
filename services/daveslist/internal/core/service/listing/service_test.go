package listing_test

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/mocks"
	"daveslist/internal/core/port"
	"daveslist/internal/core/service/listing"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx                  context.Context
	service              port.ListingService
	mockAuthService      *mocks.MockAuthService
	mockListingRepo      *mocks.MockListingRepository
	mockReplyListingRepo *mocks.MockReplyListingRepository
}

func (s *TestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.ctx = context.Background()
	s.mockAuthService = mocks.NewMockAuthService(ctrl)
	s.mockListingRepo = mocks.NewMockListingRepository(ctrl)
	s.mockReplyListingRepo = mocks.NewMockReplyListingRepository(ctrl)
	s.service = listing.New(&listing.Config{
		AuthService:      s.mockAuthService,
		ListingRepo:      s.mockListingRepo,
		ReplyListingRepo: s.mockReplyListingRepo,
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreateListing() {

	var mockTrue = true

	s.T().Run("success", func(t *testing.T) {
		s.mockListingRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.CreateListing(s.ctx, &dto.CreateListingRequest{
			IsPrivate: &mockTrue,
		})

		s.Assert().NoError(err)
	})

	s.T().Run("error: insert error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.CreateListing(s.ctx, &dto.CreateListingRequest{
			IsPrivate: &mockTrue,
		})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestGetListingList() {

	s.T().Run("success", func(t *testing.T) {
		s.mockListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockListingRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(dto.ListingListResponse{}, nil)
		_, err := s.service.GetListingList(s.ctx, &dto.ListingQuery{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: count documents mongo error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), errors.New("error"))
		_, err := s.service.GetListingList(s.ctx, &dto.ListingQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: listing not found", func(t *testing.T) {
		s.mockListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), nil)
		_, err := s.service.GetListingList(s.ctx, &dto.ListingQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: mongo error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockListingRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.GetListingList(s.ctx, &dto.ListingQuery{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestUpdateListing() {

	var (
		mockTrue        = true
		mockSliceString = []string{}
	)

	s.T().Run("success: admin", func(t *testing.T) {
		s.mockListingRepo.EXPECT().FindOneByID(s.ctx, gomock.Any()).Return(&dto.ListingResponse{}, nil)
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.UpdateListing(s.ctx, &dto.UpdateListingRequest{
			Role:      "admin",
			Picture:   &mockSliceString,
			IsPrivate: &mockTrue,
		})

		s.Assert().NoError(err)
	})

	s.T().Run("success: user", func(t *testing.T) {
		s.mockListingRepo.EXPECT().FindOneByID(s.ctx, gomock.Any()).Return(&dto.ListingResponse{
			UserID: "1",
		}, nil)
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.UpdateListing(s.ctx, &dto.UpdateListingRequest{
			UserID:    "1",
			Role:      "user",
			Picture:   &mockSliceString,
			IsPrivate: &mockTrue,
		})

		s.Assert().NoError(err)
	})

	s.T().Run("error: listing not found", func(t *testing.T) {
		s.mockListingRepo.EXPECT().FindOneByID(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.UpdateListing(s.ctx, &dto.UpdateListingRequest{
			Role:      "admin",
			Picture:   &mockSliceString,
			IsPrivate: &mockTrue,
		})

		s.Assert().Error(err)
	})

	s.T().Run("error: permission denied", func(t *testing.T) {
		s.mockListingRepo.EXPECT().FindOneByID(s.ctx, gomock.Any()).Return(&dto.ListingResponse{
			UserID: "1",
		}, nil)
		_, err := s.service.UpdateListing(s.ctx, &dto.UpdateListingRequest{
			UserID:    "2",
			Role:      "user",
			Picture:   &mockSliceString,
			IsPrivate: &mockTrue,
		})

		s.Assert().Error(err)
	})

	s.T().Run("error: update listing mongo error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().FindOneByID(s.ctx, gomock.Any()).Return(&dto.ListingResponse{}, nil)
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.UpdateListing(s.ctx, &dto.UpdateListingRequest{
			Role:      "admin",
			Picture:   &mockSliceString,
			IsPrivate: &mockTrue,
		})

		fmt.Println(err)
		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestHideListing() {

	s.T().Run("success", func(t *testing.T) {
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.HideListing(s.ctx, &dto.HideListingRequest{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: update listing mongo error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.HideListing(s.ctx, &dto.HideListingRequest{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestDeleteListing() {

	s.T().Run("success", func(t *testing.T) {
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.DeleteListing(s.ctx, &dto.DeleteListingRequest{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: update listing mongo error", func(t *testing.T) {
		s.mockListingRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.DeleteListing(s.ctx, &dto.DeleteListingRequest{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestCreateReplyListing() {

	s.T().Run("success", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(&dto.ReplyListingResponse{}, nil)
		_, err := s.service.CreateReplyListing(s.ctx, &dto.CreateReplyListingRequest{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: insert error", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.CreateReplyListing(s.ctx, &dto.CreateReplyListingRequest{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestGetReplyListingList() {
	s.T().Run("success", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockReplyListingRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(dto.ReplyListingListResponse{}, nil)
		_, err := s.service.GetReplyListingList(s.ctx, &dto.ReplyListingQuery{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: count documents mongo error", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), errors.New("error"))
		_, err := s.service.GetReplyListingList(s.ctx, &dto.ReplyListingQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: listing not found", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), nil)
		_, err := s.service.GetReplyListingList(s.ctx, &dto.ReplyListingQuery{})

		s.Assert().Error(err)

	})

	s.T().Run("error: mongo error", func(t *testing.T) {
		s.mockReplyListingRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockReplyListingRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.GetReplyListingList(s.ctx, &dto.ReplyListingQuery{})

		s.Assert().Error(err)
	})
}
