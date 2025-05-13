package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"unique"`
	Password  string `gorm:"type:varchar(255)"`
	Confirmed bool
}
