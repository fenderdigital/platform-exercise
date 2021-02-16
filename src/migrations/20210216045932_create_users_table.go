package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
CREATE TABLE users (
	id serial PRIMARY KEY,
	first_name VARCHAR ( 50 ) NOT NULL,
	last_name VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 100 ) UNIQUE NOT NULL,
	password VARCHAR ( 100 ) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS users`)
	if err != nil {
		return err
	}

	return nil
}
