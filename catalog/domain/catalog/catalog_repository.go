package catalog

import (
	"erp-back/catalog/domain/category"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	CreateCategory(category *category.Category) (pgtype.UUID, error)
	UpdateCategory(category *category.Category) (pgtype.UUID, error)
}
