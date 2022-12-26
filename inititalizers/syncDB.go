package inititalizers

import "github.com/shakoor123/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
