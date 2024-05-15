package rest

import "github.com/google/uuid"

type UpsertCategoryRequest struct {
	Id                  uuid.UUID `json:"id" binding:"required"`
	CategoryName        string    `json:"categoryName" binding:"required"`
	CategoryDescription string    `json:"categoryDescription" binding:"required"`
}
