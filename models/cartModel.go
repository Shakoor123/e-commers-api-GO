package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Uid   int
	Pid   int
	Count int
}
