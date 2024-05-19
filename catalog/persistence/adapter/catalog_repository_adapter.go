package adapter

import (
	"database/sql"
	"erp-back/catalog/domain/category"
	db "erp-back/catalog/persistence/sqlc"
	"github.com/gin-gonic/gin"
)

import (
	"context"
	"github.com/google/uuid"
)

type CatalogRepositoryAdapter struct {
}

func CreateCategory(category category.Category) (uuid.UUID, error) {

	params := db.CreateCategoryParams{
		CategoryID:          category.Id,
		CategoryName:        sql.NullString{String: category.Name.Value, Valid: true},
		CategoryDescription: sql.NullString{String: category.Description.Value, Valid: true},
	}

	ctx := context.Background()

	db.New(db.)

	categoryCreated, err := db.Queries.CreateCategory(ctx, params)
	if err != nil {
		return [16]byte{}, nil
	}
	return categoryCreated.CategoryID, err

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
