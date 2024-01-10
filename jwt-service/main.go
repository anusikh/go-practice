package main

import (
	"jwt-service/controllers"
	"jwt-service/initializers"
	"jwt-service/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}
