package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang-crud-app/database"
	"golang-crud-app/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AddProductToCart(c echo.Context) error {
	cartID, _ := strconv.ParseUint(c.Param("cartId"), 10, 32)
	productId, _ := strconv.ParseUint(c.Param("productId"), 10, 32)
	var cart models.Cart
	if err := database.DB.Scopes(database.CartByID(uint(cartID)), database.PreloadProducts).First(&cart).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	var product models.Product
	if err := database.DB.Scopes(database.ProductByID(uint(productId))).First(&product).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	cart.Products = append(cart.Products, product)
	if err := database.DB.Save(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update cart")
	}
	return c.JSON(http.StatusOK, cart)
}

func RemoveProductFromCart(c echo.Context) error {
	cartID, err := strconv.ParseUint(c.Param("cartId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart ID format")
	}
	productId, err := strconv.ParseUint(c.Param("productId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID format")
	}
	var cart models.Cart
	if err := database.DB.Preload("Products").Scopes(database.CartByID(uint(cartID))).First(&cart).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	var product models.Product
	if err := database.DB.Scopes(database.ProductByID(uint(productId))).First(&product).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	if err := database.DB.Model(&cart).Association("Products").Delete(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to remove product from cart")
	}
	return c.JSON(http.StatusOK, "Product removed from cart successfully")
}

func ReadCart(c echo.Context) error {
	cartID, err := strconv.ParseUint(c.Param("cartId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid cart ID format")
	}
	var cart models.Cart
	if err := database.DB.Scopes(database.CartByID(uint(cartID)), database.PreloadProducts).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Cart not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cart)
}
