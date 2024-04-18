package database

import (
	"golang-crud-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := database.AutoMigrate(&models.Product{}, &models.Category{}, &models.Cart{}); err != nil {
		panic("failed to auto migrate")
	}
	DB = database

	InitializeCategories()
	InitializeCart()
}

func InitializeCategories() {
	var count int64
	DB.Model(&models.Category{}).Count(&count)
	if count == 0 {
		categories := []models.Category{
			{Name: "Electronics"},
			{Name: "Books"},
			{Name: "Clothing"},
		}
		DB.Create(&categories)
	}
}

func InitializeCart() {
	var count int64
	DB.Model(&models.Cart{}).Count(&count)
	if count == 0 {
		cart := models.Cart{}
		DB.Create(&cart)
	}
}
