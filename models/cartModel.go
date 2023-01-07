package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId int
}
