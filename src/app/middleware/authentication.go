package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"platform-exer/src/app/types"
	"platform-exer/src/repos"
)

func UserFromJWT(u repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Print("error parsing token")
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token did not exist in cookie"})
				return
			}

			log.Print("error retrieving token from cookie")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error retrieving token from cookie"})
			return
		}

		token, err := jwt.ParseWithClaims(cookieToken, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})
		if err != nil {
			log.Print("error parsing token")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error parsing token"})
			return
		}

		claims, ok := token.Claims.(*types.Claims)
		if !ok {
			log.Print("couldn't parse claims")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error parsing token"})
			return
		}

		if claims.ExpiresAt < time.Now().UTC().Unix() {
			log.Print("jwt is expired")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "jwt is expired"})
			return
		}

		user, err := u.Get(claims.Email)
		if err != nil {
			log.Print("error retrieving user from claims")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error retrieving user from claims"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
