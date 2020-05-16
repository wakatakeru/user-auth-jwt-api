package infrastructure

import (
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	
	"github.com/wakatakeru/user-auth-jwt-api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	// Config for CORS (AllowOrigins for development environment)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) {userController.Create(c)})

	Router = router
}
