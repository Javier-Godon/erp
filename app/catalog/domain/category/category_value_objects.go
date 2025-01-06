package category

import (
	"gopkg.in/validator.v2"
)

type Name struct {
	Value string `validator:"nonzero"`
}

func NewName(name string) *Name {
	return &Name{
		Value: name,
	}
}

func (n *Name) validate() error {
	err := validator.Validate(n)
	if err != nil {
		return err
	}
	return nil
}

type Description struct {
	Value string
}

func NewDescription(description string) *Description {
	return &Description{
		Value: description,
	}
}
