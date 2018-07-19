package token

import (
	"testing"

	"github.com/mrsmuneton/platform-test/src/user"
)

type token_string string

func getUserStub() user.User {
	return user.User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"}
}

func TestCreateJWTToReturnTokenString(t *testing.T) {
	var e = Error{}
	var userstub user.User = getUserStub()
	_, e = CreateUserJWT(userstub)
	if e.Code != "" {
		t.Error("Unexpected token generation error")
	}
}
