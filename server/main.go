package main

import (
	"h23s_15/api"
	"h23s_15/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	instance := echo.New()
	server := handler.Server{}
	api.RegisterHandlers(instance, server)

	// model.SetUp()

	instance.Logger.Fatal(instance.Start(":8080"))
}
