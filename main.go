package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-crud-app/controllers"
	"golang-crud-app/database"
)

func main() {
	e := echo.New()

	database.InitDatabase()

	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.POST("/products", controllers.CreateProduct)
	e.GET("/products/:id", controllers.ReadProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.POST("/carts/:cartId/products/:productId", controllers.AddProductToCart)
	e.DELETE("/carts/:cartId/products/:productId", controllers.RemoveProductFromCart)
	e.GET("/carts/:cartId", controllers.ReadCart)

	e.GET("/categories", controllers.ReadCategoriesWithProducts)

	e.POST("/payment", controllers.ValidateCard)

	e.Logger.Fatal(e.Start(":8080"))
}
