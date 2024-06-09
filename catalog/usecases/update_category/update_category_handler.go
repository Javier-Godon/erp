package update_category

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/persistence/adapter"
	"log"
)

type UpdateCategoryHandler struct {
	Repository catalog.Repository
}

func NewUpdateCategoryHandler(repository catalog.Repository) *UpdateCategoryHandler {
	return &UpdateCategoryHandler{
		Repository: repository}
}

func (handler UpdateCategoryHandler) Handle(command UpdateCategoryCommand) (UpdateCategoryResult, error) {

	catalogRepository := adapter.New()
	categoryUpdatedId, err := catalogRepository.UpdateCategory(fromCategoryCommandToCategory(command))
	if err != nil {
		log.Fatalf("Could not update category with id: %v", command.Id)
	}
	result := UpdateCategoryResult{
		categoryUpdatedId,
	}
	return result, err
}

func fromCategoryCommandToCategory(command UpdateCategoryCommand) *category.Category {
	return &category.Category{
		Id:          command.Id,
		Name:        category.Name{Value: command.CategoryName},
		Description: category.Description{Value: command.CategoryDescription},
	}
}
