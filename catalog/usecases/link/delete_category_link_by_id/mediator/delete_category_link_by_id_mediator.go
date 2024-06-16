package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/link/delete_category_link_by_id"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[delete_category_link_by_id.DeleteCategoryLinkByIdCommand, delete_category_link_by_id.DeleteCategoryLinkByIdResult](delete_category_link_by_id.NewDeleteCategoryLinkByIdHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command delete_category_link_by_id.DeleteCategoryLinkByIdCommand) delete_category_link_by_id.DeleteCategoryLinkByIdResult {
	DeleteCategoryByIdResult, err := framework.Send[delete_category_link_by_id.DeleteCategoryLinkByIdCommand, delete_category_link_by_id.DeleteCategoryLinkByIdResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return DeleteCategoryByIdResult
}
