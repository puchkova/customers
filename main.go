package main

import (
	"customers/controller"
	"customers/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	storage.NewDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/customers", controller.GetCustomers)
	e.POST("/customers", controller.CreateCustomer)
	e.PUT("/customers", controller.UpdateCustomer)

	e.Logger.Fatal(e.Start(":1323"))
}
