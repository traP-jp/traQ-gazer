package main

import (
	"h23s_15/api"
	"h23s_15/handler"
	"h23s_15/traqHandler"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func init() {
	traqHandler.MessageRequestToken = os.Getenv("MESSAGE_REQUEST_TOKEN")
}

func main() {
	c := cron.New()
	c.AddFunc("* * * * *", traqHandler.MessageAPI)
	c.Start()

	instance := echo.New()
	server := handler.Server{}
	api.RegisterHandlers(instance, server)

	// model.SetUp()

	instance.Logger.Fatal(instance.Start(":8080"))
}
