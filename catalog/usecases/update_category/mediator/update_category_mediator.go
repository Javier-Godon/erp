package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/update_category"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[update_category.UpdateCategoryCommand, update_category.UpdateCategoryResult](update_category.NewUpdateCategoryHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command update_category.UpdateCategoryCommand) update_category.UpdateCategoryResult {
	UpdateCategoryResult, err := framework.Send[update_category.UpdateCategoryCommand, update_category.UpdateCategoryResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return UpdateCategoryResult
}
