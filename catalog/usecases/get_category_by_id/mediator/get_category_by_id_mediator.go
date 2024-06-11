package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/get_category_by_id"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[get_category_by_id.GetCategoryByIdCommand, get_category_by_id.GetCategoryByIdResult](get_category_by_id.NewGetCategoryByIdHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command get_category_by_id.GetCategoryByIdCommand) get_category_by_id.GetCategoryByIdResult {
	GetCategoryByIdResult, err := framework.Send[get_category_by_id.GetCategoryByIdCommand, get_category_by_id.GetCategoryByIdResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return GetCategoryByIdResult
}
