package category

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"
	errors "daveslist/pkg/go-errors"
)

func (s *Service) CreateCategory(ctx context.Context, data *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	res, err := s.categoryRepo.Insert(ctx, new(model.Category).FromDTO(data))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetCategoryList(ctx context.Context, query *dto.CategoryQuery) (*dto.PaginationResponse, error) {
	count, err := s.categoryRepo.CountByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrCategoryNotFound
	}

	res, err := s.categoryRepo.FindByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.NewPaginationResponse(query.Page, query.Limit, count, res), nil
}

func (s *Service) DeleteCategory(ctx context.Context, id string) (*dto.CategoryResponse, error) {
	t := true
	out, err := s.categoryRepo.UpdateOneByID(ctx, id, &model.UpdateCategory{IsDeleted: &t})
	if err != nil {
		return nil, err
	}

	_, err = s.listingService.DeleteListing(ctx, &dto.DeleteListingRequest{CategoryID: out.ID})
	if err != nil {
		return nil, err
	}
	return out, nil
}
