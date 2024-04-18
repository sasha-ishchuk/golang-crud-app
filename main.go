package main

import (
	"github.com/labstack/echo/v4"
	"golang-crud-app/controllers"
	"golang-crud-app/database"
)

func main() {
	e := echo.New()

	database.InitDatabase()

	e.POST("/products", controllers.CreateProduct)
	e.GET("/products/:id", controllers.ReadProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.POST("/carts/:cartId/products/:productId", controllers.AddProductToCart)
	e.DELETE("/carts/:cartId/products/:productId", controllers.RemoveProductFromCart)

	e.Logger.Fatal(e.Start(":8080"))
}
