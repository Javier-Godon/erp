package category_link

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryLink struct {
	Id               pgtype.UUID
	MainCategoryId   MainCategoryId
	LinkedCategoryId LinkedCategoryId
}
