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

	r.POST("/register", handlers.Register())
	r.POST("/login", handlers.Login())
	r.POST("/logout", handlers.Logout())

	r.GET("/user", handlers.GetUser(s.User))
	r.PUT("/user", handlers.UpdateUser(s.User))

	return r, nil
}
