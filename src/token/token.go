package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mrsmuneton/platform-test/src/error"
	"github.com/mrsmuneton/platform-test/src/user"
)

func CreateUserJWT(u user.User) (error.Error, string) {
	err := error.Error{Code: ""}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{"read": true})
	tokenstring, e := token.SignedString([]byte("2hZ2cpjxFSz8sR2MbKqo7XLz4HS6Nx4tuBWlLpvIrXQPR5O36syvcefGZAdbZisog9LWPvDCYEJajl9X"))

	if e != nil {
		err.Code = "Failed to generate jwt"
	}

	return err, tokenstring
}

func ParseJWT() {
	return
}
