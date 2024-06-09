package catalog

import (
	"erp-back/catalog/domain/category"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	CreateCategory(category *category.Category) (pgtype.UUID, error)
	UpdateCategory(category *category.Category) (pgtype.UUID, error)
	DeleteCategory(categoryId pgtype.UUID) (pgtype.UUID, error)
	FindCategoryById(categoryId pgtype.UUID) (category.Category, error)
}
