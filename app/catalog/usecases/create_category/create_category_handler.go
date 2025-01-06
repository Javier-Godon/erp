package create_category

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/persistence/adapter"
	"log"
)

type CreateCategoryHandler struct {
	Repository catalog.Repository
}

func NewCreateCategoryHandler(repository catalog.Repository) *CreateCategoryHandler {
	return &CreateCategoryHandler{
		Repository: repository}
}

func (handler CreateCategoryHandler) Handle(command CreateCategoryCommand) (CreateCategoryResult, error) {

	catalogRepository := adapter.New()
	categoryCreatedId, err := catalogRepository.CreateCategory(fromCategoryCommandToCategory(command))
	if err != nil {
		log.Fatalf("Could not create category with id: %v", command.Id)
	}
	result := CreateCategoryResult{
		categoryCreatedId,
	}
	return result, err
}

func fromCategoryCommandToCategory(command CreateCategoryCommand) *category.Category {
	return &category.Category{
		Id:          command.Id,
		Name:        category.Name{Value: command.CategoryName},
		Description: category.Description{Value: command.CategoryDescription},
	}
}
