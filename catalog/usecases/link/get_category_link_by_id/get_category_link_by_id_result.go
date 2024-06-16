package get_category_link_by_id

import "github.com/jackc/pgx/v5/pgtype"

type GetCategoryLinkByIdResult struct {
	Id               pgtype.UUID
	MainCategoryId   pgtype.UUID
	LinkedCategoryId pgtype.UUID
}
