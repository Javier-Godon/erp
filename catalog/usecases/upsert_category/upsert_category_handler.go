package upsert_category

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/persistence/adapter"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

type UpsertCategoryHandler struct {
	Repository catalog.Repository
}

func NewUpsertCategoryHandler(repository catalog.Repository) *UpsertCategoryHandler {
	return &UpsertCategoryHandler{
		Repository: repository}
}

func (handler UpsertCategoryHandler) Handle(command UpsertCategoryCommand) (UpsertCategoryResult, error) {

	catalogRepository := adapter.New()
	categoryFound, err := catalogRepository.FindCategoryById(command.Id)
	var categoryId pgtype.UUID
	if errors.Is(err, pgx.ErrNoRows) {
		categoryId, err = catalogRepository.CreateCategory(fromCategoryCommandToCategory(command))
	}
	if categoryFound.Id.Valid {
		categoryId, err = catalogRepository.UpdateCategory(fromCategoryCommandToCategory(command))
	}
	//categoryId, err := catalogRepository.UpdateCategory(fromCategoryCommandToCategory(command))
	if err != nil {
		log.Fatalf("Could not upsert category with id: %v", command.Id)
	}
	result := UpsertCategoryResult{
		categoryId,
	}
	return result, err
}

func fromCategoryCommandToCategory(command UpsertCategoryCommand) *category.Category {
	return &category.Category{
		Id:          command.Id,
		Name:        category.Name{Value: command.CategoryName},
		Description: category.Description{Value: command.CategoryDescription},
	}
}
