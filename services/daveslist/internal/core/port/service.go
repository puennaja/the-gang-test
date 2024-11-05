package port

import (
	"context"
	"daveslist/internal/core/domain/dto"
)

//go:generate mockgen -source=service.go -destination=../mocks/service_mock.go -package=mocks
type AuthService interface {
	Authorize(ctx context.Context, req *dto.AuthorizeRequest) (bool, error)
}

type CategoryService interface {
	CreateCategory(ctx context.Context, data *dto.CreateCategoryRequest) (*dto.CategoryResponse, error)
	GetCategoryList(ctx context.Context, query *dto.CategoryQuery) (*dto.PaginationResponse, error)
	DeleteCategory(ctx context.Context, id string) (*dto.CategoryResponse, error)
}

type ListingService interface {
	CreateListing(ctx context.Context, data *dto.CreateListingRequest) (*dto.ListingResponse, error)
	GetListingList(ctx context.Context, query *dto.ListingQuery) (*dto.PaginationResponse, error)
	UpdateListing(ctx context.Context, data *dto.UpdateListingRequest) (*dto.ListingResponse, error)
	HideListing(ctx context.Context, data *dto.HideListingRequest) (*dto.ListingResponse, error)
	DeleteListing(ctx context.Context, data *dto.DeleteListingRequest) (*dto.ListingResponse, error)
	CreateReplyListing(ctx context.Context, data *dto.CreateReplyListingRequest) (*dto.ReplyListingResponse, error)
	GetReplyListingList(ctx context.Context, query *dto.ReplyListingQuery) (*dto.PaginationResponse, error)
}

type MessageService interface {
	CreateMessage(ctx context.Context, data *dto.CreateMessageRequest) (*dto.MessageResponse, error)
	GetMessageList(ctx context.Context, query *dto.MessageQuery) (*dto.PaginationResponse, error)
}
