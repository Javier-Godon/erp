package category

import "github.com/jackc/pgx/v5/pgtype"

type Category struct {
	Id          pgtype.UUID
	Name        Name
	Description Description
}
