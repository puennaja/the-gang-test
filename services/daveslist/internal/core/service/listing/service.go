package listing

import (
	"context"
	"daveslist/internal/core/domain/constant"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"
	errors "daveslist/pkg/go-errors"
)

func (s *Service) CreateListing(ctx context.Context, data *dto.CreateListingRequest) (*dto.ListingResponse, error) {
	out, err := s.listingRepo.Insert(ctx, new(model.Listing).FromDTO(data))
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) GetListingList(ctx context.Context, query *dto.ListingQuery) (*dto.PaginationResponse, error) {
	count, err := s.listingRepo.CountByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrCategoryNotFound
	}

	res, err := s.listingRepo.FindByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.NewPaginationResponse(query.Page, query.Limit, count, res), nil
}

func (s *Service) UpdateListing(ctx context.Context, data *dto.UpdateListingRequest) (*dto.ListingResponse, error) {
	listing, err := s.listingRepo.FindOneByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if data.Role != constant.AdminRole {
		if listing.UserID != data.UserID {
			return nil, errors.ErrPermissionDenied
		}
	}

	update := new(model.UpdateListing).FromDTO(data)
	out, err := s.listingRepo.UpdateOneByID(ctx, data.ID, update)
	if err != nil {
		return nil, err
	}
	return out, nil

}

func (s *Service) HideListing(ctx context.Context, data *dto.HideListingRequest) (*dto.ListingResponse, error) {
	update := new(model.UpdateListing).FromHideDTO(data)
	out, err := s.listingRepo.UpdateOneByID(ctx, data.ID, update)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) DeleteListing(ctx context.Context, data *dto.DeleteListingRequest) (*dto.ListingResponse, error) {
	update := new(model.UpdateListing).FromDeleteDTO(data)
	out, err := s.listingRepo.UpdateOneByID(ctx, data.ID, update)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) CreateReplyListing(ctx context.Context, data *dto.CreateReplyListingRequest) (*dto.ReplyListingResponse, error) {
	out, err := s.replyListingRepo.Insert(ctx, new(model.ReplyListing).FromDTO(data))
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) GetReplyListingList(ctx context.Context, query *dto.ReplyListingQuery) (*dto.PaginationResponse, error) {
	count, err := s.replyListingRepo.CountByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.ErrCategoryNotFound
	}

	res, err := s.replyListingRepo.FindByQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return dto.NewPaginationResponse(query.Page, query.Limit, count, res), nil
}
