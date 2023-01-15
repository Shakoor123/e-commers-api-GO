package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
)

// create user Order
func CreateUserOrder(c *gin.Context) {
	cartid := c.Param("cartid")
	var order models.Order
	if c.Bind(&order) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "order data not found",
		})
		return
	}
	//create a user order
	result := inititalizers.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "order is not created",
		})
		return
	}
	var cart models.Cart
	result = inititalizers.DB.Where("id=?", cartid).Find(&cart)
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
	ids := []models.CartItems{}
	sub := inititalizers.DB.Where("user_id = ?", cart.UserId).Find(&ids)
	if sub.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ids not found in create order",
			"data":  ids,
		})
		return
	}
	var arrLeng int = len(ids)
	for i := 0; i < arrLeng; i++ {
		orderItem := models.OrderItems{UserId: ids[i].UserId, ProductId: ids[i].ProductId, Count: ids[i].Count, OderId: int(order.ID)}
		result = inititalizers.DB.Create(&orderItem)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "orderItem not created",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

//change order status

func ChangeOrderStatus(c *gin.Context) {
	orderid := c.Param("orderid")
	var order models.Order
	result := inititalizers.DB.Where("id=?", orderid).Find(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "order is Not found",
		})
		return
	}
	if order.Status > 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "order is max statused",
		})
		return
	}
	order.Status = order.Status + 1
	result = inititalizers.DB.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "order status is Not Changed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

// SElect User orders

func SelectUserOrder(c *gin.Context) {
	uid := c.Param("id")
	orders := []models.Order{}
	result := inititalizers.DB.Where("user_id=?", uid).Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "order is Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}
