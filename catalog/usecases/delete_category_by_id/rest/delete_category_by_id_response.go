package rest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DeleteCategoryByIdResponse struct {
	Id pgtype.UUID `json:"id" binding:"required"`
}
