package dto

import "time"

type CreateReplyListingRequest struct {
	UserID   string `header:"x-user-id" validate:"required"`
	UserName string `header:"x-user-name" validate:"required"`

	ListingID string `json:"listing_id" validate:"required"`
	Message   string `json:"message" validate:"required,max=255"`
}

type ReplyListingQuery struct {
	PaginationQuery
	ListingID string `param:"listing_id" validate:"required"`
}

type ReplyListingResponse struct {
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ReplyListingListResponse []ReplyListingResponse
