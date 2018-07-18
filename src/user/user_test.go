package user

import (
	"strconv"
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

func TestCreateUserReturnsFailure(t *testing.T) {
	var userStub = getUserStub()
	userStub.Email = ""
	var error = CreateUser(userStub)
	if error == false {
		t.Error("CreateUser should return error	")
	}
}

func TestGetUserRecordReturnsSuccess(t *testing.T) {
	var userStub = getUserStub()
	var error = CreateUser(userStub)
	if error != false {
		t.Error("CreateUser returned unexpected error	")
	}
}

func TestUpdateUserReturnsSuccess(t *testing.T) {
	var userStub = getUserStub()
	var userRecord, err = GetUserRecordByEmail(userStub.Email)
	id := strconv.Itoa(userRecord.Id)
	var updatedUser, err1 = UpdateUserFields(id, userStub)
	t.Log(updatedUser)
	t.Log(err)
	if err1 != false {
		t.Error("Update user returned unexpected error	")
	}
}

func TestDeleteUserReturnsSuccess(t *testing.T) {
	var userStub = getUserStub()
	var userRecord, err = GetUserRecordByEmail(userStub.Email)
	id := strconv.Itoa(userRecord.Id)
	err1 = DeleteUser(userRecord)
	if err1 != false {
		t.Error("Delete user returned unexpected error")
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
