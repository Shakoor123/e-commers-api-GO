package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
	// "github.com/shakoor123/inititalizers"
)

// Create WatchList for a user
func CreateUserCart(c *gin.Context) {
	uid := c.Param("id")
	id, _ := strconv.Atoi(uid)
	var cart models.Cart
	//checking user have cart
	result := inititalizers.DB.Where("user_id=?", uid).Find(&cart)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Cart is Not found",
		})
		return
	}
	if cart.UserId == 0 {
		cart.UserId = id
		result := inititalizers.DB.Create(&cart)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "cart is not created",
			})
			return
		}
	}
	//cart item binding
	var cartItem models.CartItems
	if c.Bind(&cartItem) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request watchList",
		})
		return
	}
	cartItem.UserId = id
	//check item if exist
	var oldCartItem models.CartItems
	result = inititalizers.DB.Where("product_id=? AND user_id=?", cartItem.ProductId, cartItem.UserId).Find(&oldCartItem)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cartitem error",
		})
		return
	}
	if oldCartItem.UserId == id {
		c.JSON(http.StatusBadRequest, gin.H{
			"dat": oldCartItem,
		})
		return
	}
	// creating new cart item
	result = inititalizers.DB.Create(&cartItem)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cartitem not created",
		})
		return
	}
	//response after creating
	c.JSON(http.StatusOK, gin.H{
		"data": cartItem,
	})
}
