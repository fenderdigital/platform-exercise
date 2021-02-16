package repos

import (
	"fmt"
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
		//assert.Equal

		fmt.Println("---------------")
		fmt.Println("---------------")
		fmt.Println("---------------")
		fmt.Printf("testUser:%v\n", testUser)
		fmt.Println("---------------")
	})

}