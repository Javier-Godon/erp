package upsert_category

import (
	"github.com/google/uuid"
)

type UpsertCategoryHandler interface {
	Handle() (UpsertCategoryResult, error)
}

type UpsertCategoryHandlerImpl struct {
}

func (handler UpsertCategoryHandlerImpl) Handle(command UpsertCategoryCommand) (UpsertCategoryResult, error) {
	result := UpsertCategoryResult{
		CategoryId: uuid.New(),
	}
	return result, nil
}
