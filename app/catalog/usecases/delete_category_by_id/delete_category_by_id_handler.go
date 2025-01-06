package delete_category_by_id

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/persistence/adapter"
	"erp-back/framework"
)

type DeleteCategoryByIdHandler struct {
	Repository catalog.Repository
}

func NewDeleteCategoryByIdHandler(repository catalog.Repository) *DeleteCategoryByIdHandler {
	return &DeleteCategoryByIdHandler{
		Repository: repository,
	}
}

func (handler DeleteCategoryByIdHandler) Handle(command DeleteCategoryByIdQuery) (DeleteCategoryByIdResult, error) {
	catalogRepository := adapter.New()
	uuid := framework.StringToUUID(command.Id)
	categoryId, err := catalogRepository.DeleteCategory(uuid)
	result := DeleteCategoryByIdResult{
		CategoryId: categoryId,
	}
	return result, err
}
