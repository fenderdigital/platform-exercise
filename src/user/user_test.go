package user

import (
	"testing"
)

func getUserStub() User {
	return User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"}
}

func TestValidateMinimumFieldsPresentReturnsUser(t *testing.T) {
	var userStub = getUserStub()
	e, u := ValidateUserMinimumFields(userStub)
	if e.Code != "" {
		t.Error("Validating minimum fields failed with an error")
	}
	if userStub != u {
		t.Error("Unexpected User Mutation")
	}
}

func TestFailValidateCurrentPasswordReturnsError(t *testing.T) {
	var userStub = getUserStub()
	userStub.Name = ""
	e, _ := ValidateUserMinimumFields(userStub)
	if e.Code != "Please provide all fields, including: Name" {
		t.Error("Name is a required field error not thrown")
	}
}
