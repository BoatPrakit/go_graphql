package resolvers

import "database/sql"

type QueryResolver struct {
	db *sql.DB
}
