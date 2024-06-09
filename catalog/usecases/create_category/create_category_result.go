package create_category

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateCategoryResult struct {
	CategoryId pgtype.UUID
}
