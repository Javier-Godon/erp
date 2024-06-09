package update_category

import "github.com/jackc/pgx/v5/pgtype"

type UpdateCategoryCommand struct {
	Id                  pgtype.UUID
	CategoryName        string
	CategoryDescription string
}
