package token

import (
	"testing"

	"github.com/mrsmuneton/platform-test/src/user"
)

func getUserStub() user.User {
	return user.User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"}
}

func TestCreateJWTToReturnTokenString(t *testing.T) {
	expected_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZWFkIjp0cnVlfQ.TiOLN041F3V2MSAORNU0CvPOBNjdMqMoUmRTTjnRd6w"
	_, tokenCreated := CreateUserJWT(getUserStub())
	if tokenCreated != expected_token {
		t.Error("Token does not match expected token.")
	}
}

func TestParseJWTToReturnUserReadRight(t *testing.T) {}

func TestParseJWTToReturnUserReadWriteRight(t *testing.T) {}
