package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"platform-exer/src/models"
	"platform-exer/src/repos"
)

type UserUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Update specified user handler
func UpdateUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := userFromContext(c)
		if err != nil {
			log.Printf("error retrieving user from context: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error retrieving user from context"})
			return
		}

		var uur UserUpdateRequest
		if err := c.ShouldBindJSON(&uur); err != nil {
			log.Printf("error binding json from user update request: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error binding json from user update request"})
			return
		}

		user.FirstName = uur.FirstName
		user.LastName = uur.LastName

		if err := r.Update(user); err != nil {
			log.Printf("error updating user: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error updating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
	}
}

// Delete specified user handler
func DeleteUser(r repos.UsersRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := userFromContext(c)
		if err != nil {
			log.Printf("error retrieving user from context: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error retrieving user from context"})
			return
		}

		if err := r.Delete(user); err != nil {
			log.Printf("error deleting user: %v\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "error deleting user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	}
}

// Retrieve the user from context which
// is placed from the JWT middleware
func userFromContext(c *gin.Context) (*models.User, error) {
	contextUser, ok := c.Get("user")
	if !ok {
		return nil, errors.New("error retreiving user from request")
	}
	user, ok := contextUser.(*models.User)
	if !ok {
		return nil, errors.New("error retreiving user from request")
	}

	return user, nil
}
