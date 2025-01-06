package upsert_category_link

import "github.com/jackc/pgx/v5/pgtype"

type UpsertCategoryLinkCommand struct {
	Id               pgtype.UUID
	MainCategoryId   pgtype.UUID
	LinkedCategoryId pgtype.UUID
}
