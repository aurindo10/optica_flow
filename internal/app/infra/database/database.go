package database

import (
	"database/sql"
	database "optica_flow/internal/app/infra/database/queries"

	_ "github.com/lib/pq"
)



func NewPostgresDatabase() (*database.Queries, error ){
	db, err := sql.Open("postgres", "user=docker dbname=otica password=otica123 sslmode=disable host=localhost port=5432")
	if err != nil {
		return nil , err
	}
	queries := database.New(db)
	return queries, nil
}