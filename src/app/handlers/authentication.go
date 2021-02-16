package handlers

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
