package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shakoor123/controllers"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/middlewares"
)

func main() {
	inititalizers.LoadEnvVariables()
	inititalizers.ConnectToDB()
	inititalizers.SyncDatabase()
	r := gin.Default()
	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.SignIn)
	r.GET("/api/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run()
}
