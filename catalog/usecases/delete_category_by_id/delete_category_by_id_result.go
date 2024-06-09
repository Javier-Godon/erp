package delete_category_by_id

import "github.com/jackc/pgx/v5/pgtype"

type DeleteCategoryByIdResult struct {
	CategoryId pgtype.UUID
}
