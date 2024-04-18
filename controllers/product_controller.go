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
		return err
	}
	database.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func ReadProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	database.DB.First(&product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	database.DB.First(product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := c.Bind(product); err != nil {
		return err
	}
	database.DB.Save(product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	database.DB.First(product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	database.DB.Delete(product)
	return c.JSON(http.StatusOK, "Product deleted")
}

func AddProductToCart(c echo.Context) error {
	cartID := c.Param("cartId")
	productId := c.Param("productId")

	var cart models.Cart
	if result := database.DB.First(&cart, cartID); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	var product models.Product
	if result := database.DB.First(&product, productId); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	cart.Products = append(cart.Products, product)
	database.DB.Save(&cart)

	return c.JSON(http.StatusOK, cart)
}

func RemoveProductFromCart(c echo.Context) error {
	cartID := c.Param("cartId")
	productIdParam := c.Param("productId")

	productId, err := strconv.ParseUint(productIdParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	var cart models.Cart
	if err := database.DB.Preload("Products").First(&cart, cartID).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}

	var updatedProducts []models.Product
	for _, product := range cart.Products {
		if product.ID != uint(productId) {
			updatedProducts = append(updatedProducts, product)
		}
	}

	cart.Products = updatedProducts
	if err := database.DB.Save(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update cart")
	}

	return c.JSON(http.StatusOK, cart)
}
