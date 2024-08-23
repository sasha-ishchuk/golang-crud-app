package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-crud-app/controllers"
	"golang-crud-app/database"
)

const ProductIDRoute = "/products/:id"

func main() {
	e := echo.New()

	database.InitDatabase()

	// Enable CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.GET("/logout", controllers.Logout)

	e.GET("/get_user", controllers.IsAuthenticated(controllers.GetUser))

	e.POST("/products", controllers.IsAuthenticated(controllers.CreateProduct))
	e.GET(ProductIDRoute, controllers.IsAuthenticated(controllers.ReadProduct))
	e.PUT(ProductIDRoute, controllers.IsAuthenticated(controllers.UpdateProduct))
	e.DELETE(ProductIDRoute, controllers.IsAuthenticated(controllers.DeleteProduct))

	e.POST("/carts/:cartId/products/:productId", controllers.IsAuthenticated(controllers.AddProductToCart))
	e.DELETE("/carts/:cartId/products/:productId", controllers.IsAuthenticated(controllers.RemoveProductFromCart))
	e.GET("/carts/:cartId", controllers.IsAuthenticated(controllers.ReadCart))

	e.GET("/categories", controllers.IsAuthenticated(controllers.ReadCategoriesWithProducts))

	e.POST("/payment", controllers.IsAuthenticated(controllers.ValidateCard))

	e.Logger.Fatal(e.Start(":8080"))
}
