package dto

import "time"

type ListingQuery struct {
	PaginationQuery
	CategoryID string `query:"category_id" validate:"omitempty"`
	Level      string
}

type CreateListingRequest struct {
	UserID   string `header:"x-user-id" validate:"required"`
	UserName string `header:"x-user-name" validate:"required"`

	CategoryID   string   `json:"category_id" validate:"required"`
	CategoryName string   `json:"category_name" validate:"required"`
	Title        string   `json:"title" validate:"required"`
	Content      string   `json:"content" validate:"required,omitempty,max=5000"`
	Picture      []string `json:"picture" validate:"omitempty,max=10"`
	IsPrivate    *bool    `json:"is_private" validate:"required"`
}

type UpdateListingRequest struct {
	UserID       string    `header:"x-user-id" validate:"required"`
	Role         string    `header:"x-user-role" validate:"required"`
	ID           string    `json:"listing_id" validate:"required"`
	CategoryID   string    `json:"category_id" validate:"omitempty"`
	CategoryName string    `json:"category_name" validate:"omitempty"`
	Title        string    `json:"title" validate:"omitempty"`
	Content      string    `json:"content" validate:"omitempty,max=5000"`
	Picture      *[]string `json:"picture" validate:"omitempty,max=10"`
	IsPrivate    *bool     `json:"is_private" validate:"omitempty"`
}

type HideListingRequest struct {
	ID string `param:"listing_id" validate:"required"`
}

type DeleteListingRequest struct {
	UserID     string `header:"x-user-id" validate:"required"`
	Role       string `header:"x-user-role" validate:"required"`
	CategoryID string `json:"category_id" validate:"require_with=ID"`
	ID         string `json:"listing_id" validate:"require_with=CategoryID"`
}

type ListingResponse struct {
	ID           string    `json:"listing_id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Picture      *[]string `json:"picture"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsPrivate    bool      `json:"is_private"`
	IsHide       bool      `json:"is_hide"`
	IsDeleted    bool      `json:"is_deleted"`
}

type ListingListResponse []ListingResponse
