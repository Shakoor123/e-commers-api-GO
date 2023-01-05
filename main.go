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
	r.POST("/api/admin/login", controllers.AdminLogin)
	r.GET("/api/validate", middlewares.IsAdmin, controllers.Validate)

	r.POST("/api/products", middlewares.IsAdmin, controllers.CreateProduct)
	r.GET("/api/products", middlewares.RequireAuth, controllers.SelectAllProducts)
	r.GET("/api/products/:id", middlewares.RequireAuth, controllers.SelectOneProduct)
	r.DELETE("/api/products/:id", middlewares.IsAdmin, controllers.DeleteOneProduct)
	r.PUT("/api/products/:id", middlewares.IsAdmin, controllers.UpdateOneProduct)

	r.POST("/api/watchlist", middlewares.RequireAuth, controllers.CreateWatchList)
	r.GET("/api/watchlist/:id", middlewares.RequireAuth, controllers.SelectWatchlistOfUser)

	r.Run()
}

// https://github.com/gin-gonic/gin
