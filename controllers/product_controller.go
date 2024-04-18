package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-crud-app/database"
	"golang-crud-app/models"
	"net/http"
	"strconv"
)

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, product)
}

func ReadProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := database.DB.Scopes(database.ProductByID(uint(id))).First(&product).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := database.DB.Scopes(database.ProductByID(uint(id))).First(&product).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := database.DB.Scopes(database.ProductByID(uint(id))).Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "Product deleted")
}
