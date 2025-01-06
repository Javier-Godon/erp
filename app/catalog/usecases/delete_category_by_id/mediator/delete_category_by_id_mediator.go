package mediator

import (
	"erp-back/catalog/persistence/adapter"
	"erp-back/catalog/usecases/delete_category_by_id"
	"erp-back/framework"
	"log"
)

func init() {
	err := framework.Register[delete_category_by_id.DeleteCategoryByIdQuery, delete_category_by_id.DeleteCategoryByIdResult](delete_category_by_id.NewDeleteCategoryByIdHandler(adapter.CatalogRepositoryAdapter{}))
	if err != nil {
		return
	}
}

func Send(command delete_category_by_id.DeleteCategoryByIdQuery) delete_category_by_id.DeleteCategoryByIdResult {
	DeleteCategoryByIdResult, err := framework.Send[delete_category_by_id.DeleteCategoryByIdQuery, delete_category_by_id.DeleteCategoryByIdResult](command)
	if err != nil {
		log.Fatalf("Could not execute: %v", command)
	}
	return DeleteCategoryByIdResult
}
