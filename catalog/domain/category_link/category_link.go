package category_link

import "github.com/google/uuid"

type CategoryLink struct {
	id               uuid.UUID
	mainCategoryId   Category
	linkedCategoryId Category
}
