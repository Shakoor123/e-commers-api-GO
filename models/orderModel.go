package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId int
	Total  int
	Status int
}
