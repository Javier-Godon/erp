package get_all_main_categories

import "github.com/jackc/pgx/v5/pgtype"

type GetAllMainCategoriesResult struct {
	Id          pgtype.UUID
	Name        string
	Description string
}
