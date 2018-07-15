package user

import (
	"fmt"
	"time"

	"github.com/mrsmuneton/platform-test/src/error"
)

type User struct {
	Id              int       `id`
	CreatedDate     time.Time `createdDate`
	CurrentPassword string    `currentPassword`
	Email           string    `email`
	Name            string    `name`
	UpdatedDate     time.Time `updatedDate`
}

func ValidateUserMinimumFields(u User) (error.Error, User) {
	var requiredFields string
	e := error.Error{Code: ""}
	fmt.Print(e)

	if u.CurrentPassword == "" {
		requiredFields = requiredFields + string(" CurrentPassword")
	}

	if u.Email == "" {
		requiredFields = requiredFields + string(" Email")
	}

	if u.Name == "" {
		requiredFields = requiredFields + string(" Name")
	}

	if len(requiredFields) > 0 {
		e.Code = "Please provide all fields, including:" + requiredFields
	}

	return e, u
}
