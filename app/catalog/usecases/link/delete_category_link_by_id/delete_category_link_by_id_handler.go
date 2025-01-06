package delete_category_link_by_id

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/persistence/adapter"
	"erp-back/framework"
)

type DeleteCategoryLinkByIdHandler struct {
	Repository catalog.Repository
}

func NewDeleteCategoryLinkByIdHandler(repository catalog.Repository) *DeleteCategoryLinkByIdHandler {
	return &DeleteCategoryLinkByIdHandler{
		Repository: repository,
	}
}

func (handler DeleteCategoryLinkByIdHandler) Handle(command DeleteCategoryLinkByIdCommand) (DeleteCategoryLinkByIdResult, error) {
	catalogRepository := adapter.New()
	uuid := framework.StringToUUID(command.Id)
	categoryLinkId, err := catalogRepository.DeleteCategoryLink(uuid)
	result := DeleteCategoryLinkByIdResult{
		CategoryLinkId: categoryLinkId,
	}
	return result, err
}
