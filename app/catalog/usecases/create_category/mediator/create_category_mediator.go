package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/create_category"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[create_category.CreateCategoryCommand, create_category.CreateCategoryResult](create_category.NewCreateCategoryHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command create_category.CreateCategoryCommand) create_category.CreateCategoryResult {
	CreateCategoryResult, err := framework.Send[create_category.CreateCategoryCommand, create_category.CreateCategoryResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return CreateCategoryResult
}
