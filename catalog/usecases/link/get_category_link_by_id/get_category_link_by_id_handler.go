package get_category_link_by_id

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

func (handler GetCategoryByIdHandler) Handle(command GetCategoryLinkByIdQuery) (GetCategoryLinkByIdResult, error) {
	catalogRepository := adapter.New()
	uuid := framework.StringToUUID(command.Id)
	category, err := catalogRepository.FindCategoryLinkById(uuid)
	result := GetCategoryLinkByIdResult{
		Id:               category.Id,
		MainCategoryId:   category.MainCategoryId.Value,
		LinkedCategoryId: category.LinkedCategoryId.Value,
	}
	return result, err
}
