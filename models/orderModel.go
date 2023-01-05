package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Uid    int
	Total  int
	Status int
}
