package get_category_by_id

import "github.com/jackc/pgx/v5/pgtype"

type GetCategoryByIdResult struct {
	Id          pgtype.UUID
	Name        string
	Description string
}
