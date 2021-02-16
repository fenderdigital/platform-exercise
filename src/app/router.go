package app

import (
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
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/user", handlers.GetUser(s.User))
	r.PUT("/user", handlers.UpdateUser(s.User))

	return r, nil
}
