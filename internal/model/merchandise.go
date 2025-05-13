package model

import "gorm.io/gorm"

type Merchandise struct {
	gorm.Model

	Name       string `gorm:"unique"`
	VisitCount uint64
}
