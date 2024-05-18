package upsert_category

import "github.com/google/uuid"

type UpsertCategoryCommand struct {
	Id                  uuid.UUID
	categoryName        string
	categoryDescription string
}
