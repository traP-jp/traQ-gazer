package main

import (
	"h23s_15/api"
	"h23s_15/handler"
	"h23s_15/traqHandler"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	instance := echo.New()
	server := handler.Server{}
	api.RegisterHandlers(instance, server)

	// model.SetUp()

	go func() {
		instance.Logger.Fatal(instance.Start(":8080"))
		wg.Done()
	}()

	traqServer := traqHandler.Start()
	go func() {
		log.Fatal(traqServer.ListenAndServe(":8000"))
	}()

	wg.Wait()
}
