package traqmessage

import (
	"context"
	"os"
	"time"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

var ACCESS_TOKEN = os.Getenv("BOT_ACCESS_TOKEN")

func collectMessages(from time.Time, to time.Time) (*traq.MessageSearchResult, error) {
	if ACCESS_TOKEN == "" {
		slog.Info("Skip collectMessage")
		return &traq.MessageSearchResult{}, nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, ACCESS_TOKEN)

	result, _, err := client.MessageApi.SearchMessages(auth).After(from).Before(to).Execute()
	if err != nil {
		return nil, err
	}

	return result, nil
}
