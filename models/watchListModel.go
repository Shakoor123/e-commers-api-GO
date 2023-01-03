package models

import "gorm.io/gorm"

type WatchList struct {
	gorm.Model
	Uid int
	Pid int
}
