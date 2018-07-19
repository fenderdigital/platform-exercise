package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mrsmuneton/platform-test/src/secret"
	"github.com/mrsmuneton/platform-test/src/user"
)

type Error struct {
	Code string `code`
}

func CreateUserJWT(u user.User) (string, Error) {
	err := Error{Code: ""}
	var signing_secret string = fetchSecret()

	jwToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwToken.Claims.(jwt.MapClaims)

	week := time.Now().AddDate(0, 0, 7).Unix()
	claims["expire"] = week
	claims["userId"] = u.Id

	tokenstring, e := jwToken.SignedString([]byte(signing_secret))
	if e != nil {
		err.Code = "Failed to generate jwt"
	}

	return tokenstring, err
}

func ParseJWT(tokenString string) bool {
	var signing_secret string = fetchSecret()
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signing_secret), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println(claims["userId"])
	} else {
		fmt.Println(err)
	}
	return true
}

func fetchSecret() string {
	secret := secret.Fetch()
	return secret
}
