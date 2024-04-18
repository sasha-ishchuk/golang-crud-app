package database

import (
	"gorm.io/gorm"
)

func PreloadProducts(db *gorm.DB) *gorm.DB {
	return db.Preload("Products")
}

func ProductByID(id uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func CartByID(id uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
