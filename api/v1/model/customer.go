package model

import "gorm.io/gorm"

type Customer struct {
	*gorm.Model
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}
