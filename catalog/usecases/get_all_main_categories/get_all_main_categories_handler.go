package get_all_main_categories

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category"
	"erp-back/catalog/persistence/adapter"
)

type GetAllMainCategoriesHandler struct {
	Repository catalog.Repository
}

func NewGetAllMainCategoriesHandler(repository catalog.Repository) *GetAllMainCategoriesHandler {
	return &GetAllMainCategoriesHandler{
		Repository: repository,
	}
}

func (handler GetAllMainCategoriesHandler) Handle(command GetAllMainCategoriesCommand) ([]GetAllMainCategoriesResult, error) {
	catalogRepository := adapter.New()
	categories, err := catalogRepository.FindAllMainCategories()

	return mapCatalogCategoryToCategory(categories), err
}

func mapCatalogCategoryToCategory(categories []category.Category) (categoryResults []GetAllMainCategoriesResult) {
	for _, catalogCategory := range categories {
		categoryResults = append(categoryResults, fromCategoryDomainToResult(catalogCategory))
	}
	return categoryResults
}

func fromCategoryDomainToResult(category category.Category) GetAllMainCategoriesResult {
	return GetAllMainCategoriesResult{
		Id:          category.Id,
		Name:        category.Name.Value,
		Description: category.Description.Value,
	}
}
