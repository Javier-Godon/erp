package upsert_category

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/persistence/adapter"
	"erp-back/framework"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type UpsertCategoryHandler struct {
	Db         *pgxpool.Pool
	Repository catalog.Repository
}

func NewUpsertCategoryHandler(repository catalog.Repository) *UpsertCategoryHandler {
	return &UpsertCategoryHandler{
		Db:         framework.DB,
		Repository: repository}
}

func (handler UpsertCategoryHandler) Handle(command UpsertCategoryCommand) (UpsertCategoryResult, error) {

	catalogRepository := adapter.New(framework.DB)
	categoryCreatedId, err := catalogRepository.CreateCategory(fromCategoryCommandToCategory(command))
	if err != nil {
		log.Fatalf("Could not create category with id: %v", command.Id)
	}
	result := UpsertCategoryResult{
		categoryCreatedId,
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
