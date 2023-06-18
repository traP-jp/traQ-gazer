package main

import (
	"h23s_15/api"
	"h23s_15/handler"
	"h23s_15/model"
	"h23s_15/traqmessage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	instance := echo.New()
	instance.Use(middleware.Logger())

	server := handler.Server{}

	apiInstance := instance.Group("/api")
	api.RegisterHandlers(apiInstance, server)

	// まとめて賢くルーティングするのは厳しそうなので
	instance.Static("/", "dist")
	instance.File("/words", "dist/index.html")
	instance.File("/words/add", "dist/index.html")

	err := model.SetUp()
	if err != nil {
		slog.Info("Error setting up: %v", err)
	}

	go func() {
		traqmessage.PollingMessages()
	}()

	instance.Logger.Fatal(instance.Start(":8080"))
}
