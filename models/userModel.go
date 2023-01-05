package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Admin    bool
	PhoneNo  int
	Address  string
	State    string
	District string
}
