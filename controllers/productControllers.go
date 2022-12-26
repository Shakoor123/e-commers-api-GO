package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
)

func CreateProduct(c *gin.Context) {
	type Data struct {
		Title string
		// Image    string
		Price    int
		Category string
		Color    string
		Size     string
	}
	var data Data

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "product data not found	",
		})
		return
	}
	product := models.Product{Title: data.Title, Price: data.Price, Category: data.Category, Color: data.Color, Size: data.Size}
	result := inititalizers.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Not inserted",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func SelectAllProducts(c *gin.Context) {
	var products []models.Product
	result := inititalizers.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Products not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}
func SelectOneProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	result := inititalizers.DB.First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func DeleteOneProduct(c *gin.Context) {
	id := c.Param("id")

	result := inititalizers.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Product delete successfully",
		"id":   id,
	})
}
func UpdateOneProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	results := inititalizers.DB.First(&product, id)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found",
		})
		return
	}

	if c.Bind(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "product data not found	",
		})
		return
	}

	result := inititalizers.DB.Save(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
