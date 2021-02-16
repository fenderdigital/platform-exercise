package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"platform-exer/src/app/types"
	"platform-exer/src/models"
	"platform-exer/src/repos"

	"github.com/dgrijalva/jwt-go"
)

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Register user
func Register(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("error retrieving request: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("error generating password from request: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		u := models.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Password:  string(pass),
		}

		if err = r.Create(&u); err != nil {
			log.Printf("error generating password from request: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		// TODO make this value configurable
		expirationTime := time.Now().Add(time.Minute * 5)
		claims := &types.Claims{
			Email: u.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(os.Getenv("TOKEN_SECRET"))
		if err != nil {
			log.Printf("error creating jwt: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		// TODO - potentially redirect user
		c.JSON(http.StatusCreated, gin.H{"message": "user logged in successfully"})
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login user
func Login(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lr LoginRequest
		if err := c.ShouldBindJSON(&lr); err != nil {
			log.Printf("error generating password from request: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		user, err := r.Get(lr.Email)
		if err != nil {
			log.Printf("error retrieving user", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lr.Password))
		if err != nil {
			log.Printf("error retrieving user", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "incorrect password"})
			return
		}

		// TODO make this value configurable
		expirationTime := time.Now().Add(time.Minute * 5)
		claims := &types.Claims{
			Email: lr.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(os.Getenv("TOKEN_SECRET"))
		if err != nil {
			log.Printf("error creating jwt: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}
}

// Logout user - invalidate token
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
