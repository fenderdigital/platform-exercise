package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"platform-exer/src/repos"
	"platform-exer/src/config"
)

type Services struct {
	DB *gorm.DB

	User repos.UsersRepo
}

// Initialize application services
func InitServices() (Services, error) {
	var s Services
	var err error

	s.DB, err = InitDB()
	if err != nil {
		return s, err
	}

	s.User = repos.NewUsersRepo(s.DB)

	return s, nil
}

// Initialize the gorm with config values
func InitDB() (*gorm.DB, error) {
	config := config.EnvPostgresConfig()

	db, err := gorm.Open(postgres.Open(config.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
