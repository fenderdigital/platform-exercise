package user

import (
	"testing"
)

func getUserStub() User {
	return User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"}
}

func TestCreateUserReturnsSuccess(t *testing.T) {
	var userStub = getUserStub()
	var error = CreateUser(userStub)
	if error != false {
		t.Error("CreateUser returned unexpected error	")
	}
}

func TestValidateMinimumFieldsFailReturnsError(t *testing.T) {
	var userStub = getUserStub()
	userStub.Name = ""
	_, e := ValidateUserMinimumFields(userStub)
	if e.Code != "Please provide all fields, including: Name" {
		t.Error("Name is a required field error not thrown")
	}
}

func TestValidateMinimumFieldsPresentReturnsUser(t *testing.T) {
	var userStub = getUserStub()
	u, e := ValidateUserMinimumFields(userStub)
	if e.Code != "" {
		t.Error("Validating minimum fields failed with an error")
	}
	if userStub != u {
		t.Error("Unexpected User Mutation")
	}
}
