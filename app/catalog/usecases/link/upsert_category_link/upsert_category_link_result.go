package upsert_category_link

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UpsertCategoryLinkResult struct {
	CategoryLinkId pgtype.UUID
}
