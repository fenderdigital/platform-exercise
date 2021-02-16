package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"platform-exer/src/models"
	"platform-exer/src/repos"
)

func GetUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := r.Get()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error retrieving user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User
		// TODO finish logic
		if err := r.Create(&u); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error creating user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
	}
}

func UpdateUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := r.Update(nil); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error updating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}