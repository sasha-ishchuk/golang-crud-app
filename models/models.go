package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint
}

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Cart struct {
	gorm.Model
	Products []Product `gorm:"many2many:cart_products;"`
}

type Card struct {
	gorm.Model
	Number       string
	ExpireDate   string
	SecurityCode string
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Token    string
	Service  string //`gorm:"not null"`
	GoToken  string //`gorm:"not null"`
}
