package main

import (
	"h23s_15/api"
	"h23s_15/handler"
	"h23s_15/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	instance := echo.New()
	instance.Use(middleware.Logger())

	server := handler.Server{}

	apiInstance := instance.Group("/api")
	api.RegisterHandlers(apiInstance, server)

	instance.Static("/", "dist")
	instance.File("/*", "dist/index.html")

	model.SetUp()

	instance.Logger.Fatal(instance.Start(":8080"))
}
