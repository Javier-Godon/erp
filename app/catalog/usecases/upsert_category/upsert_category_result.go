package upsert_category

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UpsertCategoryResult struct {
	CategoryId pgtype.UUID
}
