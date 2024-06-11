package catalog

import "github.com/jackc/pgx/v5/pgtype"

type CategoryResponse struct {
	Id                  pgtype.UUID `json:"id" binding:"required"`
	CategoryName        string      `json:"categoryName" binding:"required"`
	CategoryDescription string      `json:"categoryDescription" binding:"required"`
}
