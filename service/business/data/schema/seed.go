package schema

import (
	"github.com/jmoiron/sqlx"
)

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.
//('452bb454-7a23-4dc4-800b-d384fc7c2814', 'Admin', 'admin@admin.com', '{ADMIN,USER}', '$2a$10$iMaLYkXKdKybF0Zl15JtUOx0B66lU71aV6Ryziu1JnGCxHrPF4Xfa', '2020-09-01 00:00:00', '2020-09-01 00:00:00'),

const seeds = `
-- Create admin and regular User with password "gophers"
INSERT INTO users (user_id, name, email, roles, password_hash, date_created, date_updated) VALUES
	('9c223318-deb4-416a-9778-fad0df4edf98', 'admin', 'admin@admin.com', '{ADMIN,USER}', '$2a$10$1E9MeZMU8Nox9dJXW/dB1unw4naWoAUsE/WWZBqgtGcvxazF57o86', '2020-09-01 00:00:00', '2020-09-01 00:00:00')
	 ON CONFLICT DO NOTHING;
	`

// DeleteAll runs the set of Drop-table queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func DeleteAll(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(deleteAll); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// deleteAll is used to clean the database between tests.
const deleteAll = `
DELETE FROM users;`
