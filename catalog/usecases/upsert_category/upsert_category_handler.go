package upsert_category

import (
	"fmt"
	"github.com/google/uuid"
)

type UpsertCategoryHandler struct {
}

func (handler UpsertCategoryHandler) Handle(command UpsertCategoryCommand) (UpsertCategoryResult, error) {
	result := UpsertCategoryResult{
		CategoryId: uuid.New(),
	}
	fmt.Printf("in upsertCategoryHandler, generated uuid: %v", result.CategoryId.String())
	return result, nil
}
