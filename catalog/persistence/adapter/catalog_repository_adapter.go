package adapter

import (
	"database/sql"
	"erp-back/catalog/domain/category"
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

func (c CatalogRepositoryAdapter) DeleteCategory(categoryId pgtype.UUID) (pgtype.UUID, error) {

	ctx := context.Background()
	queries := db.New(framework.DB)

	err := queries.DeleteCategoryById(ctx, categoryId)

	return categoryId, err

}
