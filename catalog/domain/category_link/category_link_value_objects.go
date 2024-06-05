package category_link

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gopkg.in/validator.v2"
)

type Category struct {
	Value pgtype.UUID `validator:"nonzero"`
}

func NewCategory(id pgtype.UUID) *Category {
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
