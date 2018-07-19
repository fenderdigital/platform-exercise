package token

import (
	"fmt"
	"strconv"
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

	fmt.Println(u)
	fmt.Println(u.Id)

	week := time.Now().AddDate(0, 0, 7).Unix()
	claims["expire"] = week
	claims["userId"] = u.Id

	tokenstring, e := jwToken.SignedString([]byte(signing_secret))
	if e != nil {
		err.Code = "Failed to generate jwt"
	}

	return tokenstring, err
}

func ParseJWT(tokenString string) (string, bool) {
	var signing_secret string = fetchSecret()
	parsedToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signing_secret), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userid := strconv.FormatFloat(claims["userId"].(float64), 'f', -1, 64)
		return userid, true
	} else {
		return "0", false
	}
}

func fetchSecret() string {
	secret := secret.Fetch()
	return secret
}
