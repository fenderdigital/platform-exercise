package types

import (
	"github.com/dgrijalva/jwt-go"
)

// this file prevents circular dependency error by
// creating a side layer for the Claims type
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

