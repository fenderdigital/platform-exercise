package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mrsmuneton/platform-test/src/error"
	"github.com/mrsmuneton/platform-test/src/secret"
	"github.com/mrsmuneton/platform-test/src/user"
)

func CreateUserJWT(u user.User) (error.Error, string) {
	err := error.Error{Code: ""}
	secret := secret.Fetch()
	fmt.Println(secret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	week := time.Now().AddDate(0, 0, 7).Unix()
	claims["expire"] = week
	claims["userId"] = u.Id

	tokenstring, e := token.SignedString([]byte(secret))
	if e != nil {
		err.Code = "Failed to generate jwt"
	}

	return err, tokenstring
}

func ParseJWT() {
	return
}
