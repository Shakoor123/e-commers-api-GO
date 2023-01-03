package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shakoor123/inititalizers"
	"github.com/shakoor123/models"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *gin.Context) {
	type Body struct {
		Email    string
		Password string
	}
	var body Body
	//get data from user
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}
	//select user from db
	var user models.User
	inititalizers.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	//check password hash and body password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password is incorrect",
		})
		return
	}

	// check is admin
	if user.Admin == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your not Authorized",
		})
		return
	}

	//create an jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "jwt token error",
		})
		return
	}

	//return response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 36000*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
