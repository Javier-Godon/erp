package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/link/get_category_link_by_id"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[get_category_link_by_id.GetCategoryLinkByIdQuery, get_category_link_by_id.GetCategoryLinkByIdResult](get_category_link_by_id.NewGetCategoryByIdHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command get_category_link_by_id.GetCategoryLinkByIdQuery) get_category_link_by_id.GetCategoryLinkByIdResult {
	GetCategoryByIdResult, err := framework.Send[get_category_link_by_id.GetCategoryLinkByIdQuery, get_category_link_by_id.GetCategoryLinkByIdResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return GetCategoryByIdResult
}
