package category_link

import (
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Category struct {
	Value uuid.UUID `validator:"nonzero"`
}

func NewCategory(id uuid.UUID) *Category {
	return &Category{
		Value: id,
	}
}

func (n *Category) validate() error {
	err := validator.Validate(n)
	if err != nil {
		return err
	}
	return nil
}
