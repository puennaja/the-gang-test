package port

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"
)

//go:generate mockgen -source=repository.go -destination=../mocks/repository_mock.go -package=mocks

type CategoryRepository interface {
	Insert(ctx context.Context, data *model.Category) (*dto.CategoryResponse, error)
	FindByQuery(ctx context.Context, query *dto.CategoryQuery) (dto.CategoryListResponse, error)
	CountByQuery(ctx context.Context, query *dto.CategoryQuery) (int64, error)
	UpdateOneByID(ctx context.Context, id string, update *model.UpdateCategory) (*dto.CategoryResponse, error)
}

type ListingRepository interface {
	Insert(ctx context.Context, data *model.Listing) (*dto.ListingResponse, error)
	FindByQuery(ctx context.Context, query *dto.ListingQuery) (dto.ListingListResponse, error)
	CountByQuery(ctx context.Context, query *dto.ListingQuery) (int64, error)
	UpdateByQuery(ctx context.Context, query *dto.ListingQuery, update *model.UpdateListing) (int64, error)
	FindOneByID(ctx context.Context, id string) (*dto.ListingResponse, error)
	UpdateOneByID(ctx context.Context, id string, update *model.UpdateListing) (*dto.ListingResponse, error)
}

type ReplyListingRepository interface {
	Insert(ctx context.Context, data *model.ReplyListing) (*dto.ReplyListingResponse, error)
	FindByQuery(ctx context.Context, query *dto.ReplyListingQuery) (dto.ReplyListingListResponse, error)
	CountByQuery(ctx context.Context, query *dto.ReplyListingQuery) (int64, error)
}

type MessageRepository interface {
	Insert(ctx context.Context, data *model.Message) (*dto.MessageResponse, error)
	FindByQuery(ctx context.Context, query *dto.MessageQuery) (dto.MessageListResponse, error)
	CountByQuery(ctx context.Context, query *dto.MessageQuery) (int64, error)
}
