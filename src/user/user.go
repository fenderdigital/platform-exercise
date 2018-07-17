package user

import (
	"fmt"

	"github.com/mrsmuneton/platform-test/src/error"
)

//using timestamsp ca improve sorting efficiency in queries
type User struct {
	Id              int    `id`
	CreatedDate     string `createdDate` //cheating by a string, this must be a timestamp
	CurrentPassword string `currentPassword`
	Email           string `email`
	Name            string `name`
	UpdatedDate     string `updatedDate` //cheating by a string, this must be a timestamp
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
