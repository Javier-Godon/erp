package catalog

import (
	"erp-back/catalog/domain/category"
	"erp-back/catalog/domain/category_link"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository interface {
	CreateCategory(category *category.Category) (pgtype.UUID, error)
	UpdateCategory(category *category.Category) (pgtype.UUID, error)
	DeleteCategory(categoryId pgtype.UUID) (pgtype.UUID, error)
	FindCategoryById(categoryId pgtype.UUID) (category.Category, error)
	CreateCategoryLink(categoryLink *category_link.CategoryLink) (pgtype.UUID, error)
	UpdateCategoryLink(categoryLink *category_link.CategoryLink) (pgtype.UUID, error)
	DeleteCategoryLink(categoryLinkId pgtype.UUID) (pgtype.UUID, error)
	FindCategoryLinkById(categoryLinkId pgtype.UUID) (category_link.CategoryLink, error)
	FindAllMainCategories() ([]category.Category, error)
	//FindParentCategory(categoryId pgtype.UUID) (category.Category, error)
	//FindParentCategoriesChain(categoryId pgtype.UUID) ([]category.Category, error)
	//FindChildCategoriesChain(categoryId pgtype.UUID) ([]category.Category, error)
}
