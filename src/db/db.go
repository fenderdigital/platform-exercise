package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {
	connStr := "user=postgres password=postgres dbname=platform sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
