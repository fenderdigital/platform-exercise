package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"platform-exer/src/app/handlers"
)

func InitRouter(s Services, ginMode string) (*gin.Engine, error) {
	gin.SetMode(ginMode)
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/register", handlers.Register(s.User))
	r.POST("/login", handlers.Login(s.User))
	r.POST("/logout", handlers.Logout())

	r.PUT("/user", handlers.UpdateUser(s.User))
	r.DELETE("/user", handlers.DeleteUser(s.User))

	return r, nil
}
