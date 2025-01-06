package get_category_by_id

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/persistence/adapter"
	"erp-back/framework"
)

type GetCategoryByIdHandler struct {
	Repository catalog.Repository
}

func NewGetCategoryByIdHandler(repository catalog.Repository) *GetCategoryByIdHandler {
	return &GetCategoryByIdHandler{
		Repository: repository,
	}
}

func (handler GetCategoryByIdHandler) Handle(command GetCategoryByIdCommand) (GetCategoryByIdResult, error) {
	catalogRepository := adapter.New()
	uuid := framework.StringToUUID(command.Id)
	category, err := catalogRepository.FindCategoryById(uuid)
	result := GetCategoryByIdResult{
		Id:          category.Id,
		Name:        category.Name.Value,
		Description: category.Description.Value,
	}
	return result, err
}
