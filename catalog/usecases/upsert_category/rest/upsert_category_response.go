package rest

import "github.com/google/uuid"

type UpsertCategoryResponse struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
