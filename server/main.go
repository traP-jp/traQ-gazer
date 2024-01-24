package main

import (
	"traQ-gazer/api"
	"traQ-gazer/handler"
	"traQ-gazer/model"
	"traQ-gazer/traqmessage"

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

	messagePoller := traqmessage.NewMessagePoller()
	go messagePoller.Run()

	instance.Logger.Fatal(instance.Start(":8080"))
}
