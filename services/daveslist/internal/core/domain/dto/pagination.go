package dto

import "math"

const (
	QueryCreateAt     = "created_at"
	QueryUpdateAt     = "updated_at"
	SortDirectionDesc = "desc"
	SortDirectionAsc  = "asc"
)

type PaginationQuery struct {
	Page  int64 `json:"page" query:"page" validate:"required,gt=0"`
	Limit int64 `json:"limit" query:"limit" validate:"required,gte=-1,ne=0,lte=100"`
}

type PaginationResponse struct {
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int64       `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func NewPaginationResponse(page, limit, totalRows int64, rows interface{}) *PaginationResponse {
	return &PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: int64(math.Ceil(float64(totalRows) / float64(limit))),
		Rows:       rows,
	}
}
