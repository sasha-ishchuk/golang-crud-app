package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-crud-app/database"
	"golang-crud-app/models"
	"net/http"
)

func ReadCategoriesWithProducts(c echo.Context) error {
	var categories []models.Category
	if err := database.DB.Scopes(database.PreloadProducts).Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve categories")
	}
	return c.JSON(http.StatusOK, categories)
}
