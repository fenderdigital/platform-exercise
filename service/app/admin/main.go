package main

import (
	"fmt"
	"log"
	"os"

	"platform-exercise/service/app/admin/commands"
	"platform-exercise/service/foundation/database"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

func main() {
	log := log.New(os.Stdout, "ADMIN : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(log); err != nil {
		if errors.Cause(err) != commands.ErrHelp {
			log.Printf("error: %s", err)
		}
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	var cfg struct {
		Args conf.Args
		DB   struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:0.0.0.0"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:false"`
		}
	}

	const prefix = "TEST"
	if err := conf.Parse(os.Args[1:], prefix, &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage(prefix, &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main: Config :\n%v\n", out)

	// =========================================================================
	// Commands

	dbConfig := database.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := commands.Migrate(dbConfig); err != nil {
			return errors.Wrap(err, "migrating database")
		}

	case "seed":
		if err := commands.Seed(dbConfig); err != nil {
			return errors.Wrap(err, "seeding database")
		}

	case "genkey":
		if err := commands.GenKey(); err != nil {
			return errors.Wrap(err, "key generation")
		}

	case "gentoken":
		id := cfg.Args.Num(1)
		privateKeyFile := cfg.Args.Num(2)
		algorithm := cfg.Args.Num(3)
		if err := commands.GenToken("000", log, dbConfig, id, privateKeyFile, algorithm); err != nil {
			return errors.Wrap(err, "generating token")
		}

	default:
		fmt.Println("migrate: create the schema in the database")
		fmt.Println("seed: add data to the database")
		fmt.Println("provide a command to get more help.")
		fmt.Println("genkey: generate a set of private/public key files")
		fmt.Println("gentoken: generate a JWT for a user with claims")
		return commands.ErrHelp
	}

	return nil
}
