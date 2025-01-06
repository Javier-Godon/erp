package category_link

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gopkg.in/validator.v2"
)

type MainCategoryId struct {
	Value pgtype.UUID `validator:"nonzero"`
}

func NewMainCategoryIdOfLink(id pgtype.UUID) *MainCategoryId {
	return &MainCategoryId{
		Value: id,
	}
}

func (n *MainCategoryId) validate() error {
	err := validator.Validate(n)
	if err != nil {
		return err
	}
	return nil
}

type LinkedCategoryId struct {
	Value pgtype.UUID `validator:"nonzero"`
}

func NewLinkedCategoryIdOfLink(id pgtype.UUID) *LinkedCategoryId {
	return &LinkedCategoryId{
		Value: id,
	}
}

func (n *LinkedCategoryId) validate() error {
	err := validator.Validate(n)
	if err != nil {
		return err
	}
	return nil
}
