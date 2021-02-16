package config

import (
	"fmt"
	"os"
)

func EnvPostgresConfig() PostgresConfig {
	return PostgresConfig{
		User:     os.Getenv("DATABASE_USER"),
		Database: os.Getenv("DATABASE_NAME"),
		SSLMode:  os.Getenv("DATABASE_SSL_MODE"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     os.Getenv("DATABASE_PORT"),
		Host:     os.Getenv("DATABASE_HOST"),
	}
}

type PostgresConfig struct {
	Host     string
	User     string
	Port     string
	Password string
	Database string
	SSLMode  string
}

func (c *PostgresConfig) FormatDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Database)
}
