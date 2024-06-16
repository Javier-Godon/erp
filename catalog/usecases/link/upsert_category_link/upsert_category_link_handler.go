package upsert_category_link

import (
	"erp-back/catalog/domain/catalog"
	"erp-back/catalog/domain/category_link"
	"erp-back/catalog/persistence/adapter"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
)

type UpsertCategoryLinkHandler struct {
	Repository catalog.Repository
}

func NewUpsertCategoryLinkHandler(repository catalog.Repository) *UpsertCategoryLinkHandler {
	return &UpsertCategoryLinkHandler{
		Repository: repository}
}

func (handler UpsertCategoryLinkHandler) Handle(command UpsertCategoryLinkCommand) (UpsertCategoryLinkResult, error) {

	catalogRepository := adapter.New()
	categoryFound, err := catalogRepository.FindCategoryLinkById(command.Id)
	var categoryId pgtype.UUID
	if errors.Is(err, pgx.ErrNoRows) {
		categoryId, err = catalogRepository.CreateCategoryLink(fromCategoryLinkCommandToCategoryLink(command))
	}
	if categoryFound.Id.Valid {
		categoryId, err = catalogRepository.UpdateCategoryLink(fromCategoryLinkCommandToCategoryLink(command))
	}

	if err != nil {
		log.Fatalf("Could not upsert category link with id: %v", command.Id)
	}
	result := UpsertCategoryLinkResult{
		categoryId,
	}
	return result, err
}

func fromCategoryLinkCommandToCategoryLink(command UpsertCategoryLinkCommand) *category_link.CategoryLink {
	return &category_link.CategoryLink{
		Id:               command.Id,
		MainCategoryId:   category_link.MainCategoryId{Value: command.MainCategoryId},
		LinkedCategoryId: category_link.LinkedCategoryId{Value: command.LinkedCategoryId},
	}
}
