package adapter

import (
	"database/sql"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/domain/category_link"
	db "erp-back/catalog/persistence/sqlc"
	"erp-back/framework"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

import (
	"context"
)

type CatalogRepositoryAdapter struct {
	db *pgxpool.Pool
}

func New() *CatalogRepositoryAdapter {
	return &CatalogRepositoryAdapter{}
}

func (c CatalogRepositoryAdapter) CreateCategory(category *category.Category) (pgtype.UUID, error) {

	params := db.CreateCategoryParams{
		CategoryID:          category.Id,
		CategoryName:        pgtype.Text(sql.NullString{String: category.Name.Value, Valid: true}),
		CategoryDescription: pgtype.Text(sql.NullString{String: category.Description.Value, Valid: true}),
	}

	ctx := context.Background()
	queries := db.New(framework.DB)

	categoryCreated, err := queries.CreateCategory(ctx, params)

	return categoryCreated.CategoryID, err

}

func (c CatalogRepositoryAdapter) CreateCategoryLink(categoryLink *category_link.CategoryLink) (pgtype.UUID, error) {

	params := db.CreateCategoryLinkParams{
		CategoryLinkID:   categoryLink.Id,
		MainCategoryID:   categoryLink.MainCategoryId.Value,
		LinkedCategoryID: categoryLink.LinkedCategoryId.Value,
	}

	ctx := context.Background()
	queries := db.New(framework.DB)

	categoryCreated, err := queries.CreateCategoryLink(ctx, params)

	return categoryCreated.CategoryLinkID, err

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (c CatalogRepositoryAdapter) UpdateCategory(category *category.Category) (pgtype.UUID, error) {
	params := db.UpdateCategoryParams{
		CategoryID:          category.Id,
		CategoryName:        pgtype.Text(sql.NullString{String: category.Name.Value, Valid: true}),
		CategoryDescription: pgtype.Text(sql.NullString{String: category.Description.Value, Valid: true}),
	}

	ctx := context.Background()
	queries := db.New(framework.DB)

	categoryCreated, err := queries.UpdateCategory(ctx, params)

	return categoryCreated.CategoryID, err

}

func (c CatalogRepositoryAdapter) UpdateCategoryLink(categoryLink *category_link.CategoryLink) (pgtype.UUID, error) {
	params := db.UpdateCategoryLinkParams{
		CategoryLinkID:   categoryLink.Id,
		MainCategoryID:   categoryLink.MainCategoryId.Value,
		LinkedCategoryID: categoryLink.LinkedCategoryId.Value,
	}

	ctx := context.Background()
	queries := db.New(framework.DB)

	categoryCreated, err := queries.UpdateCategoryLink(ctx, params)

	return categoryCreated.CategoryLinkID, err

}

func (c CatalogRepositoryAdapter) DeleteCategory(categoryId pgtype.UUID) (pgtype.UUID, error) {

	ctx := context.Background()
	queries := db.New(framework.DB)

	err := queries.DeleteCategoryById(ctx, categoryId)

	return categoryId, err

}

func (c CatalogRepositoryAdapter) DeleteCategoryLink(categoryLinkId pgtype.UUID) (pgtype.UUID, error) {

	ctx := context.Background()
	queries := db.New(framework.DB)

	err := queries.DeleteCategoryLinkById(ctx, categoryLinkId)

	return categoryLinkId, err

}

func (c CatalogRepositoryAdapter) FindCategoryById(categoryId pgtype.UUID) (category.Category, error) {
	ctx := context.Background()
	queries := db.New(framework.DB)

	catalogCategory, err := queries.FindCategoryById(ctx, categoryId)

	return fromCatalogEntityToDomain(catalogCategory), err
}

func fromCatalogEntityToDomain(catalogCategory db.CatalogCategory) (catalog category.Category) {
	return category.Category{
		Id:          catalogCategory.CategoryID,
		Name:        *category.NewName(catalogCategory.CategoryName.String),
		Description: *category.NewDescription(catalogCategory.CategoryDescription.String),
	}
}

func (c CatalogRepositoryAdapter) FindCategoryLinkById(categoryLinkId pgtype.UUID) (category_link.CategoryLink, error) {
	ctx := context.Background()
	queries := db.New(framework.DB)

	catalogCategoryLink, err := queries.FindCategoryLinkById(ctx, categoryLinkId)

	return fromCatalogCategoryLinkEntityToDomain(catalogCategoryLink), err
}

func fromCatalogCategoryLinkEntityToDomain(catalogCategoryLInk db.CatalogCategoryLink) (categoryLink category_link.CategoryLink) {
	return category_link.CategoryLink{
		Id:               catalogCategoryLInk.CategoryLinkID,
		MainCategoryId:   *category_link.NewMainCategoryIdOfLink(catalogCategoryLInk.MainCategoryID),
		LinkedCategoryId: *category_link.NewLinkedCategoryIdOfLink(catalogCategoryLInk.LinkedCategoryID),
	}
}

func (c CatalogRepositoryAdapter) FindAllMainCategories() ([]category.Category, error) {
	ctx := context.Background()
	queries := db.New(framework.DB)

	catalogCategories, err := queries.FindAllMainCategories(ctx)

	return mapCatalogCategoryToCategory(catalogCategories), err
}

func mapCatalogCategoryToCategory(catalogCategories []db.CatalogCategory) (categories []category.Category) {
	for _, catalogCategory := range catalogCategories {
		categories = append(categories, fromCatalogEntityToDomain(catalogCategory))
	}
	return categories
}

//func (c CatalogRepositoryAdapter) FindParentCategory(categoryId pgtype.UUID) (category.Category, error) {
//	ctx := context.Background()
//	queries := db.New(framework.DB)
//
//	categoryLink, err := queries.FindCategoryLinkByLinkedCategoryId(ctx, categoryId)
//	categoryFound, err := queries.FindCategoryById(ctx, categoryLink.MainCategoryID)
//
//	return fromCatalogEntityToDomain(categoryFound), err
//}
