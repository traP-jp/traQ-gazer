package main

import (
	"fmt"
	"traQ-gazer/api"
	"traQ-gazer/client"
	"traQ-gazer/db"
	"traQ-gazer/traqmessage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	instance := echo.New()
	instance.Use(middleware.Logger())

	server := client.Server{}

	apiInstance := instance.Group("/api")
	api.RegisterHandlers(apiInstance, server)

	// まとめて賢くルーティングするのは厳しそうなので
	instance.Static("/", "dist")
	instance.File("/words", "dist/index.html")
	instance.File("/words/add", "dist/index.html")

	err := db.SetUp()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set up: %v", err))
	}

	messagePoller := traqmessage.NewMessagePoller()
	go messagePoller.Run()

	instance.Logger.Fatal(instance.Start(":8080"))
}
