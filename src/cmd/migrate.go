package cmd

import (
	"github.com/spf13/cobra"
	"github.com/pressly/goose"

	"platform-exer/src/app"

	_ "platform-exer/src/migrations"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "goose migrations (go run main.go migrate up)",
	RunE:  migrate,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	MigrateCmd.SetUsageFunc(func(c *cobra.Command) error {
		c.Println(`
Usage: pe migrate [OPTIONS] COMMAND

Drivers:
    postgres
    mysql
    sqlite3
    redshift

Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp

Examples:
    pe migrate status
    pe migrate create init sql
    pe migrate create add_some_column sql
    pe migrate create fetch_user_data go
    pe migrate up

    pe migrate status`)
		return nil
	})
}

func migrate(_ *cobra.Command, args []string) error {
	db, err := app.InitDB()
	if err != nil {
		return err
	}

	var arguments []string
	if len(args) > 1 {
		arguments = args[1:]
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return goose.Run(args[0], sqlDB, ".", arguments...)
}
