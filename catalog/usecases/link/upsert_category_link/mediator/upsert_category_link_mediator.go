package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/link/upsert_category_link"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[upsert_category_link.UpsertCategoryLinkCommand, upsert_category_link.UpsertCategoryLinkResult](upsert_category_link.NewUpsertCategoryLinkHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command upsert_category_link.UpsertCategoryLinkCommand) upsert_category_link.UpsertCategoryLinkResult {
	UpsertCategoryResult, err := framework.Send[upsert_category_link.UpsertCategoryLinkCommand, upsert_category_link.UpsertCategoryLinkResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return UpsertCategoryResult
}
