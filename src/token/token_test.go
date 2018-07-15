package token

import (
	"testing"
)

func TestCreateJWTToReturnTokenString(t *testing.T) {
	t.Log("shine sun shine")
	// email := "buffalo@shuffle.com"
	// password := "and4llthatJ4ZZ"
	expected_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
	tokenCreated := CreateUserJWT()
	// tokenCreated := token.CreateJWT(email, password)
	if tokenCreated != expected_token {
		t.Error("Token does not match expected token.")
	}
}

func TestParseJWTToReturnUserReadRight(t *testing.T) {}

func TestParseJWTToReturnUserReadWriteRight(t *testing.T) {}
