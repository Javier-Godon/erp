package category_link

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryLink struct {
	id               pgtype.UUID
	mainCategoryId   Category
	linkedCategoryId Category
}
