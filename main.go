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

	r.Use(middlewares.CORSMiddleware())

	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.SignIn)
	r.POST("/api/admin/login", controllers.AdminLogin)
	r.GET("/api/validate", middlewares.IsAdmin, controllers.Validate)

	r.POST("/api/products", middlewares.IsAdmin, controllers.CreateProduct)
	r.GET("/api/products", middlewares.RequireAuth, controllers.SelectAllProducts)
	r.GET("/api/products/*key&value", middlewares.RequireAuth, controllers.SelectCategoryProducts)
	r.GET("/api/product/:id", middlewares.RequireAuth, controllers.SelectOneProduct)
	r.DELETE("/api/products/:id", middlewares.IsAdmin, controllers.DeleteOneProduct)
	r.PUT("/api/products/:id", middlewares.IsAdmin, controllers.UpdateOneProduct)

	r.POST("/api/watchlist", middlewares.RequireAuth, controllers.CreateWatchList)
	r.GET("/api/watchlist/:id", middlewares.RequireAuth, controllers.SelectWatchlistOfUser)

	r.POST("/api/cart/:id", middlewares.RequireAuth, controllers.CreateUserCart)
	r.DELETE("/api/cart/:id", middlewares.RequireAuth, controllers.RemoveCartItem)
	r.DELETE("/api/cart/all/:id", middlewares.RequireAuth, controllers.RemoveOneUserCart)
	r.GET("/api/cart/:id", middlewares.RequireAuth, controllers.SelectCartOfUser)

	r.POST("/api/order/:cartid", middlewares.RequireAuth, controllers.CreateUserOrder)
	r.PUT("/api/order/:orderid", middlewares.IsAdmin, controllers.ChangeOrderStatus)
	r.GET("/api/order/:id", middlewares.RequireAuth, controllers.SelectUserOrder)

	// r.Use(func(c *gin.Context) {
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
	// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type")
	// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 		c.AbortWithStatus(204)
	// 	}
	// })

	r.Run()
}

// https://github.com/gin-gonic/gin
