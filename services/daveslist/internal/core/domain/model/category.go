package model

import (
	"daveslist/internal/core/domain/dto"
	"daveslist/pkg/utils"
	"time"
)

type Category struct {
	ID        string    `bson:"id"`
	Name      string    `bson:"name"`
	IsPrivate bool      `bson:"is_private"`
	IsDeleted bool      `bson:"is_deleted"`
	CreatedAt time.Time `bson:"created_at"`
}

func (c *Category) FromDTO(dto *dto.CreateCategoryRequest) *Category {
	if dto == nil {
		return nil
	}
	return &Category{
		ID:        utils.GetUUID(),
		Name:      dto.Name,
		IsPrivate: dto.IsPrivate,
		IsDeleted: false,
		CreatedAt: time.Now().UTC(),
	}
}

func (c *Category) ToDTO() *dto.CategoryResponse {
	if c == nil {
		return nil
	}
	return &dto.CategoryResponse{
		ID:        c.ID,
		Name:      c.Name,
		IsPrivate: c.IsPrivate,
		IsDeleted: c.IsDeleted,
		CreatedAt: c.CreatedAt.String(),
	}
}

type CategoryList []Category

func (c CategoryList) ToDTO() dto.CategoryListResponse {
	out := make(dto.CategoryListResponse, 0)
	for _, v := range c {
		out = append(out, *v.ToDTO())
	}
	return out
}

type UpdateCategory struct {
	Name      string `bson:"name,omitempty"`
	IsPrivate *bool  `bson:"is_private,omitempty"`
	IsDeleted *bool  `bson:"is_deleted,omitempty"`
}

func (u *UpdateCategory) FromDeleteDTO(dto *dto.DeleteCategoryRequest) *UpdateCategory {
	if dto == nil {
		return nil
	}
	t := true
	return &UpdateCategory{
		IsDeleted: &t,
	}
}
