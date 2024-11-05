package dto

type CategoryQuery struct {
	PaginationQuery
	Level string
}

type CreateCategoryRequest struct {
	Name      string `json:"name" validate:"required"`
	IsPrivate bool   `json:"is_private"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id" validate:"required"`
}

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
}

type CategoryListResponse []CategoryResponse
