package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"platform-exer/src/config"
	"platform-exer/src/models"
)

func TestUsersRepo(t *testing.T) {
	config := config.EnvPostgresConfig()
	db, err := gorm.Open(postgres.Open(config.FormatDSN()), &gorm.Config{})
	if err != nil {
		t.Fatalf("error: %v\n", err.Error())
	}

	// hacky solution for development testing.
	// proper way would be to create a new
	// database with a random string and
	// drop that database after test run

	// if you don't truncate - you run into race conditions on
	// subsequent runs of the tests
	if err = db.Exec(`TRUNCATE users`).Error; err != nil {
		t.Fatalf("error truncating users table: %v\n", err.Error())
	}

	userRepo := NewUsersRepo(db)

	t.Run("Create", func(t *testing.T) {
		testUser := models.User{
			FirstName: "TestFirst",
			LastName:  "TestLast",
			Email:     "test@test.com",
			Password:  "hashed-password",
		}
		err := userRepo.Create(&testUser)
		require.NoError(t, err, "error creating user")
		assert.NotEqual(t, 0, testUser.ID, "testUser.ID should not have a default value of zero")
	})

	t.Run("Update", func(t *testing.T) {
		testUser := models.User{
			FirstName: "TestName",
			LastName:  "TestLastName",
			Email:     "test@update.com",
			Password:  "hashed-password",
		}

		err := db.Create(&testUser).Error
		require.NoError(t, err, "error creating user")

		testUser.FirstName = "NewFirst"
		testUser.LastName = "NewLast"
		err = userRepo.Update(&testUser)
		require.NoError(t, err, "error updating user")

		var found models.User
		err = db.Where("id = ?", testUser.ID).Find(&found).Error
		require.NoError(t, err, "error updating user")
		assert.Equal(t, testUser.FirstName, found.FirstName)
		assert.Equal(t, testUser.LastName, found.LastName)
	})
}