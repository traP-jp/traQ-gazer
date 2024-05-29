package main

import (
	"fmt"
	"traQ-gazer/db"
	"traQ-gazer/message"
	"traQ-gazer/oapi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/slog"
)

func main() {
	instance := echo.New()
	instance.Use(middleware.Logger())

	server := oapi.Server{}

	apiInstance := instance.Group("/api")
	oapi.RegisterHandlers(apiInstance, server)

	// まとめて賢くルーティングするのは厳しそうなので
	instance.Static("/", "dist")
	instance.File("/words", "dist/index.html")
	instance.File("/words/add", "dist/index.html")

	err := db.SetUp()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set up: %v", err))
	}

	messagePoller := message.NewMessagePoller()
	go messagePoller.Run()

	instance.Logger.Fatal(instance.Start(":8080"))
}
