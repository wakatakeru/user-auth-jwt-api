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

	// Endpoint for Users
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users/:name", func(c *gin.Context) { userController.Show(c) })
	router.PATCH("/users/:name", func(c *gin.Context) { userController.Update(c) })

	// Endpoint for Login
	router.POST("/login", func(c *gin.Context) { userController.Login(c) })

	Router = router
}
