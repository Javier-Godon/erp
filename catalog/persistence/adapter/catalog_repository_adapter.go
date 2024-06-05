package adapter

import (
	"database/sql"
	"erp-back/catalog/domain/category"
	db "erp-back/catalog/persistence/sqlc"
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

func New(db *pgxpool.Pool) *CatalogRepositoryAdapter {
	return &CatalogRepositoryAdapter{
		db: db,
	}
}

func (c CatalogRepositoryAdapter) CreateCategory(category *category.Category) (pgtype.UUID, error) {

	params := db.CreateCategoryParams{
		CategoryID:          category.Id,
		CategoryName:        pgtype.Text(sql.NullString{String: category.Name.Value, Valid: true}),
		CategoryDescription: pgtype.Text(sql.NullString{String: category.Description.Value, Valid: true}),
	}

	ctx := context.Background()
	queries := db.New(c.db)

	categoryCreated, err := queries.CreateCategory(ctx, params)

	return categoryCreated.CategoryID, err

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
