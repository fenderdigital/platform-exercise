package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"platform-exer/src/repos"
)

func GetUser(repo repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := repo.Get()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error retrieving user"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(repo repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		//if err := repo.Update(); err != nil {
		//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error updating user"})
		//	return
		//}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}