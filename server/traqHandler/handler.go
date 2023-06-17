package traqHandler

import (
	"log"
	"os"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
	"golang.org/x/exp/slog"
)

type traqServerInterface interface {
	SetPingHandler() error
	SetMessageCreatedHandler(p *traqbot.MessageCreatedPayload) error
	SetChannelCreatedHandler() error
}

func Start() *traqbot.BotServer {
	slog.Info("traQ http server started on 8000")

	client := traq.NewAPIClient(traq.NewConfiguration())
	// TOKEN := os.Getenv("TRAQ_CLIENT_TOKEN")
	// auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

	traQBotVrToken := os.Getenv("TRAQ_BOT_VERIFICATION_TOKEN")

	var traqServer traqServerInterface
	traqServer = TraqServer{
		client: client,
	}

	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(func(payload *traqbot.PingPayload) {
		err := traqServer.SetPingHandler()
		if err != nil {
			log.Println(err)
		}
	})
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		err := traqServer.SetMessageCreatedHandler(payload)
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

	server := traqbot.NewBotServer(traQBotVrToken, handlers)
	return server
}
