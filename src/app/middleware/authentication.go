package middleware

import (
	"github.com/gin-gonic/gin"

	"platform-exer/src/repos"
)

func UserFromJWT(u repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve token
		// validate token is good
		//  - true - fetch user from db
		//  - true - put user on context
		//  - false
	}
}
