package token

import (
	"fmt"
	"testing"

	"github.com/mrsmuneton/platform-test/src/user"
)

// type token_string string

func getUserStub() user.User {
	return user.User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"}
}

func TestCreateJWTToReturnTokenString(t *testing.T) {
	var userstub user.User = getUserStub()
	var ts, e = CreateUserJWT(userstub)
	fmt.Println(ts)
	if e.Code != "" {
		t.Error("Unexpected token generation error")
	}
}

// func TestParseJWTToReturnUserReadRight(t *testing.T) {}
//
// func TestParseJWTToReturnUserReadWriteRight(t *testing.T) {}
