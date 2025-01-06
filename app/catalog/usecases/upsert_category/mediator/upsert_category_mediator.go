package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/upsert_category"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[upsert_category.UpsertCategoryCommand, upsert_category.UpsertCategoryResult](upsert_category.NewUpsertCategoryHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command upsert_category.UpsertCategoryCommand) upsert_category.UpsertCategoryResult {
	UpsertCategoryResult, err := framework.Send[upsert_category.UpsertCategoryCommand, upsert_category.UpsertCategoryResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return UpsertCategoryResult
}
