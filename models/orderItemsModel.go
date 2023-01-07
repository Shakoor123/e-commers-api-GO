package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	UserId    int
	ProductId int
	Count     int
}
