package model

import (
	"daveslist/internal/core/domain/dto"
	"daveslist/pkg/utils"
	"time"
)

type ReplyListing struct {
	ID        string    `bson:"id"`
	ListingID string    `bson:"listing_id"`
	UserID    string    `bson:"user_id"`
	UserName  string    `bson:"user_name"`
	Message   string    `bson:"message"`
	CreatedAt time.Time `bson:"created_at"`
}

func (r *ReplyListing) FromDTO(dto *dto.CreateReplyListingRequest) *ReplyListing {
	if dto == nil {
		return nil
	}
	return &ReplyListing{
		ID:        utils.GetUUID(),
		ListingID: dto.ListingID,
		UserID:    dto.UserID,
		UserName:  dto.UserName,
		Message:   dto.Message,
		CreatedAt: time.Now(),
	}
}

func (r *ReplyListing) ToDTO() *dto.ReplyListingResponse {
	return &dto.ReplyListingResponse{
		UserID:    r.UserID,
		UserName:  r.UserName,
		Message:   r.Message,
		CreatedAt: r.CreatedAt,
	}
}

type ReplyListingList []ReplyListing

func (r ReplyListingList) ToDTO() dto.ReplyListingListResponse {
	res := make(dto.ReplyListingListResponse, 0)
	for i, v := range r {
		res[i] = *v.ToDTO()
	}
	return res
}
