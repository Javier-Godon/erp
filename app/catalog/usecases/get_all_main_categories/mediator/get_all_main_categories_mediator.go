package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/get_all_main_categories"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[get_all_main_categories.GetAllMainCategoriesCommand, []get_all_main_categories.GetAllMainCategoriesResult](get_all_main_categories.NewGetAllMainCategoriesHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command get_all_main_categories.GetAllMainCategoriesCommand) []get_all_main_categories.GetAllMainCategoriesResult {
	GetAllMainCategoriesResult, err := framework.Send[get_all_main_categories.GetAllMainCategoriesCommand, []get_all_main_categories.GetAllMainCategoriesResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return GetAllMainCategoriesResult
}
