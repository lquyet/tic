package main

import (
	"demo/apis"
	"demo/database"
	"demo/entity"
	"demo/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	// First init users data
	database.Users = utils.GenerateMockData[entity.User](10)
	utils.SetUsersAvatarURL()
	database.Products = utils.GenerateMockData[entity.Product](10)
	utils.SetProductsImageURL()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Routes
	// User routes
	e.GET("/users", apis.GetUsers)
	e.POST("/users", apis.UpsertUser)
	e.DELETE("/users/:id", apis.DeleteUser)
	e.GET("/users/reset", apis.ResetUserData)

	// Product routes
	e.GET("/products", apis.GetProducts)
	e.POST("/products", apis.UpsertProduct)
	e.DELETE("/products/:id", apis.DeleteProduct)
	e.GET("/products/reset", apis.ResetProductData)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
