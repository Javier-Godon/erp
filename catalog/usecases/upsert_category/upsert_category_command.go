package upsert_category

import "github.com/jackc/pgx/v5/pgtype"

type UpsertCategoryCommand struct {
	Id                  pgtype.UUID
	CategoryName        string
	CategoryDescription string
}
