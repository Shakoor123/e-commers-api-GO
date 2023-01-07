package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
)

// Create WatchList for a user
func CreateWatchList(c *gin.Context) {
	type DataModel struct {
		Uid       int
		ProductId int
	}
	var data DataModel

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request watchList",
		})
		return
	}

	watchList := models.WatchList{Uid: data.Uid, ProductId: data.ProductId}
	result := inititalizers.DB.Create(&watchList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Not inserted watchList",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": watchList,
	})
}

//Select Watchlists of A User

func SelectWatchlistOfUser(c *gin.Context) {
	uid := c.Param("id")

	watchLists := []models.WatchList{}
	result := inititalizers.DB.Where("uid=?", uid).Find(&watchLists)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "watchLists not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": watchLists,
	})
}

//Remove One WatchlistItem from user

func RemoveOneWatchlistItem(c *gin.Context) {
	wid := c.Param("id")
	result := inititalizers.DB.Delete(&models.WatchList{}, wid)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Watchlist item not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Watchlist item delete successfully",
	})
}

//Remove All Watchlist of a User

func RemoveAllWatchlist(c *gin.Context) {
	uid := c.Param("id")
	result := inititalizers.DB.Where("uid = ?", uid).Delete(&models.WatchList{})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Watchlists not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Watchlists deleted successfully",
	})
}
