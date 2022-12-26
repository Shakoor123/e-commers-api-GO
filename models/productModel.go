package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title    string `gorm:"unique"`
	Image    string
	Price    int
	Category string
	Color    string
	Size     string
}
