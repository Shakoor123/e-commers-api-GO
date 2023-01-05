package models

import "gorm.io/gorm"

type WatchList struct {
	gorm.Model
	Uid       int
	ProductId int
}
