package catalog

import "github.com/jackc/pgx/v5/pgtype"

type CategoryLinkResponse struct {
	Id               pgtype.UUID `json:"id" binding:"required"`
	MainCategoryId   pgtype.UUID `json:"mainCategoryId" binding:"required"`
	LinkedCategoryId pgtype.UUID `json:"linkedCategoryId" binding:"required"`
}
