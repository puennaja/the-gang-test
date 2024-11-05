package model

import (
	"daveslist/internal/core/domain/dto"
	"daveslist/pkg/utils"
	"time"
)

type Listing struct {
	ID           string    `bson:"id"`
	UserID       string    `bson:"user_id"`
	UserName     string    `bson:"user_name"`
	CategoryID   string    `bson:"category_id"`
	CategoryName string    `bson:"category_name"`
	Title        string    `bson:"title"`
	Content      string    `bson:"content"`
	Picture      []string  `bson:"picture"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	IsPrivate    bool      `bson:"is_private"`
	IsDeleted    bool      `bson:"is_deleted"`
	IsHide       bool      `bson:"is_hide"`
}

func (l *Listing) FromDTO(dto *dto.CreateListingRequest) *Listing {
	if dto == nil {
		return nil
	}
	return &Listing{
		ID:           utils.GetUUID(),
		UserID:       dto.UserID,
		UserName:     dto.UserName,
		CategoryID:   dto.CategoryID,
		CategoryName: dto.CategoryName,
		Title:        dto.Title,
		Content:      dto.Content,
		Picture:      dto.Picture,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		IsPrivate:    *dto.IsPrivate,
		IsDeleted:    false,
		IsHide:       false,
	}
}

func (l *Listing) ToDTO() *dto.ListingResponse {
	return &dto.ListingResponse{
		ID:           l.ID,
		UserID:       l.UserID,
		UserName:     l.UserName,
		CategoryID:   l.CategoryID,
		CategoryName: l.CategoryName,
		Title:        l.Title,
		Content:      l.Content,
		Picture:      &l.Picture,
		CreatedAt:    l.CreatedAt,
		UpdatedAt:    l.UpdatedAt,
		IsPrivate:    l.IsPrivate,
		IsHide:       l.IsHide,
		IsDeleted:    l.IsDeleted,
	}
}

type ListingList []Listing

func (l ListingList) ToDTO() dto.ListingListResponse {
	res := make(dto.ListingListResponse, 0)
	for _, v := range l {
		res = append(res, *v.ToDTO())
	}
	return res
}

type UpdateListing struct {
	CategoryID   string    `bson:"category_id,omitempty"`
	CategoryName string    `bson:"category_name,omitempty"`
	Title        string    `bson:"title,omitempty"`
	Content      string    `bson:"content,omitempty"`
	Picture      *[]string `bson:"picture,omitempty"`
	IsPrivate    *bool     `bson:"is_private,omitempty"`
	IsDeleted    *bool     `bson:"is_deleted,omitempty"`
	IsHide       *bool     `bson:"is_hide,omitempty"`
	UpdatedAt    time.Time `bson:"updated_at"`
}

func (l *UpdateListing) FromDTO(dto *dto.UpdateListingRequest) *UpdateListing {
	if dto == nil {
		return nil
	}
	return &UpdateListing{
		CategoryID:   dto.CategoryID,
		CategoryName: dto.CategoryName,
		Title:        dto.Title,
		Content:      dto.Content,
		Picture:      dto.Picture,
		IsPrivate:    dto.IsPrivate,
		IsDeleted:    nil,
		IsHide:       nil,
		UpdatedAt:    time.Now().UTC(),
	}
}

func (l *UpdateListing) FromHideDTO(dto *dto.HideListingRequest) *UpdateListing {
	if dto == nil {
		return nil
	}
	t := true
	return &UpdateListing{
		IsHide:    &t,
		UpdatedAt: time.Now().UTC(),
	}
}

func (l *UpdateListing) FromDeleteDTO(dto *dto.DeleteListingRequest) *UpdateListing {
	if dto == nil {
		return nil
	}
	t := true
	return &UpdateListing{
		IsDeleted: &t,
		UpdatedAt: time.Now().UTC(),
	}
}
