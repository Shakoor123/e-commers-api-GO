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

	r.POST("/api/products", middlewares.RequireAuth, controllers.CreateProduct)
	r.GET("/api/products", middlewares.RequireAuth, controllers.SelectAllProducts)
	r.GET("/api/products/:id", middlewares.RequireAuth, controllers.SelectOneProduct)
	r.DELETE("/api/products/:id", middlewares.RequireAuth, controllers.DeleteOneProduct)
	r.PUT("/api/products/:id", middlewares.RequireAuth, controllers.UpdateOneProduct)

	r.Run()
}

// https://github.com/gin-gonic/gin
