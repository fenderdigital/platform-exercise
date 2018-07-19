package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DBConnect() (*sql.DB, error) {
	connStr := "host=db user=postgres password=postgres dbname=platform sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
