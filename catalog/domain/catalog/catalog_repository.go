package catalog

import (
	"erp-back/catalog/domain/category"
	"github.com/google/uuid"
)

type Repository interface {
	CreateCategory(category *category.Category) (uuid.UUID, error)
}
