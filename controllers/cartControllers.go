package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
)

// Create cart for a user
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

//Remove one Cart Item

func RemoveCartItem(c *gin.Context) {
	cartItemId := c.Param("id")
	result := inititalizers.DB.Delete(&models.CartItems{}, cartItemId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cartItem  not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "cart item deleted successfully",
	})
}

// clear a user Cart

func RemoveOneUserCart(c *gin.Context) {
	userId := c.Param("id")
	result := inititalizers.DB.Delete(&models.Cart{}, "user_id LIKE ?", userId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cart  not found",
		})
		return
	}

	result = inititalizers.DB.Where("user_id LIKE ?", userId).Delete(&models.CartItems{})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cartItems  not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "cart deleted successfully",
	})
}

// Select One user cart

func SelectCartOfUser(c *gin.Context) {
	uid := c.Param("id")
	var cart models.Cart
	// selcting User cart
	result := inititalizers.DB.Where("user_id=?", uid).Find(&cart)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Cart is Not found",
		})
		return
	}
	if cart.UserId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cart is not found",
		})
		return
	}
	products := []models.Product{}
	type Ids struct {
		ProductId int
		Count     int
	}
	ids := []Ids{}
	//selecting cart items ids
	sub := inititalizers.DB.Model(&models.CartItems{}).Select("product_id,count").Where("user_id = ?", uid).Find(&ids)
	if sub.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "products1 not found",
			"data":  ids,
		})
		return
	}
	//assigning only ids for selecting products details
	var arrLeng int = len(ids)
	var productIds = make([]int, arrLeng)
	for i := 0; i < arrLeng; i++ {
		productIds[i] = ids[i].ProductId
	}
	//selecting products using ids
	result = inititalizers.DB.Where("id IN ?", productIds).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "products2 not found",
			"data":  ids,
		})
		return
	}
	// Respond giving products and its Counts
	c.JSON(http.StatusOK, gin.H{
		"data":          products,
		"productCounts": ids,
	})
}
