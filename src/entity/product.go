package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          int
	ProductName string
	ProductDesc string
}
