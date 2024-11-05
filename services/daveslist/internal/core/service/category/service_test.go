package category_test

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/mocks"
	"daveslist/internal/core/port"
	"daveslist/internal/core/service/category"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ctx                context.Context
	service            port.CategoryService
	mockCategoryRepo   *mocks.MockCategoryRepository
	mockListingService *mocks.MockListingService
}

func (s *TestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.ctx = context.Background()
	s.mockCategoryRepo = mocks.NewMockCategoryRepository(ctrl)
	s.mockListingService = mocks.NewMockListingService(ctrl)
	s.service = category.New(&category.Config{
		CategoryRepo:   s.mockCategoryRepo,
		ListingService: s.mockListingService,
	})
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreateCategory() {

	s.T().Run("success", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(&dto.CategoryResponse{}, nil)
		_, err := s.service.CreateCategory(s.ctx, &dto.CreateCategoryRequest{})

		s.Assert().NoError(err)
	})

	s.T().Run("error", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().Insert(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.CreateCategory(s.ctx, &dto.CreateCategoryRequest{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestGetCategoryList() {

	s.T().Run("success", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockCategoryRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(dto.CategoryListResponse{}, nil)
		_, err := s.service.GetCategoryList(s.ctx, &dto.CategoryQuery{})

		s.Assert().NoError(err)
	})

	s.T().Run("error: count documents mongo error", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), errors.New("error"))
		_, err := s.service.GetCategoryList(s.ctx, &dto.CategoryQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: category not found", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(0), nil)
		_, err := s.service.GetCategoryList(s.ctx, &dto.CategoryQuery{})

		s.Assert().Error(err)
	})

	s.T().Run("error: mongo error", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().CountByQuery(s.ctx, gomock.Any()).Return(int64(1), nil)
		s.mockCategoryRepo.EXPECT().FindByQuery(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.GetCategoryList(s.ctx, &dto.CategoryQuery{})

		s.Assert().Error(err)
	})
}

func (s *TestSuite) TestDeleteCategory() {

	s.T().Run("success", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.CategoryResponse{}, nil)
		s.mockListingService.EXPECT().DeleteListing(s.ctx, gomock.Any()).Return(&dto.ListingResponse{}, nil)
		_, err := s.service.DeleteCategory(s.ctx, "1")

		s.Assert().NoError(err)
	})

	s.T().Run("error: update category mongo error", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.DeleteCategory(s.ctx, "1")

		s.Assert().Error(err)
	})

	s.T().Run("error: delete listing service error", func(t *testing.T) {
		s.mockCategoryRepo.EXPECT().UpdateOneByID(s.ctx, gomock.Any(), gomock.Any()).Return(&dto.CategoryResponse{}, nil)
		s.mockListingService.EXPECT().DeleteListing(s.ctx, gomock.Any()).Return(nil, errors.New("error"))
		_, err := s.service.DeleteCategory(s.ctx, "1")
		s.Assert().Error(err)
	})
}
