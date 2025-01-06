package create_category

import "github.com/jackc/pgx/v5/pgtype"

type CreateCategoryCommand struct {
	Id                  pgtype.UUID
	CategoryName        string
	CategoryDescription string
}
