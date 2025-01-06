package delete_category_link_by_id

import "github.com/jackc/pgx/v5/pgtype"

type DeleteCategoryLinkByIdResult struct {
	CategoryLinkId pgtype.UUID
}
