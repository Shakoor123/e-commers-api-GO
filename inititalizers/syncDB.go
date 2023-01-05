package inititalizers

import "github.com/shakoor123/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.CartItems{})
	DB.AutoMigrate(&models.OrderItems{})
	DB.AutoMigrate(&models.WatchList{})

}
