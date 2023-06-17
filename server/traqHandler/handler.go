package traqHandler

import (
	"log"
	"os"

	traqbot "github.com/traPtitech/traq-bot"
)

type traqServerInterface interface {
	SetPingHandler() error
	SetMessageCreatedHandler() error
	SetChannelCreatedHandler() error
}

func Start() *traqbot.BotServer {
	log.Println("traQ http server started on 8000")
	vt := os.Getenv("VERIFICATION_TOKEN")

	var traqServer traqServerInterface

	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(func(payload *traqbot.PingPayload) {
		err := traqServer.SetPingHandler()
		if err != nil {
			log.Println(err)
		}
	})
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		err := traqServer.SetMessageCreatedHandler()
		if err != nil {
			log.Println(err)
		}
	})
	handlers.SetChannelCreatedHandler(func(payload *traqbot.ChannelCreatedPayload) {
		err := traqServer.SetChannelCreatedHandler()
		if err != nil {
			log.Println(err)
		}
	})

	server := traqbot.NewBotServer(vt, handlers)
	return server
}
