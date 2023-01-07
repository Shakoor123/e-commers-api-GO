package models

import "gorm.io/gorm"

type CartItems struct {
	gorm.Model
	UserId    int
	ProductId int
	Count     int
}
