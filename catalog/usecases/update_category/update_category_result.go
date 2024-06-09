package update_category

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateCategoryResult struct {
	CategoryId pgtype.UUID
}
