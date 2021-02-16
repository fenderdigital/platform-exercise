package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"platform-exer/src/repos"
)

func UpdateUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := r.Update(nil); err != nil {
			log.Printf("error updating user: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error updating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}

func DeleteUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := r.Delete(nil); err != nil {
			log.Printf("error deleting user: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error deleting user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	}
}
