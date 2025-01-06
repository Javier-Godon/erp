package rest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UpsertCategoryLinkRequest struct {
	Id               pgtype.UUID `json:"id" binding:"required"`
	MainCategoryId   pgtype.UUID `json:"mainCategoryId" binding:"required"`
	LinkedCategoryId pgtype.UUID `json:"linkedCategoryId" binding:"required"`
}
