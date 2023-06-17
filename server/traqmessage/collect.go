package traqmessage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

var ACCESS_TOKEN = os.Getenv("BOT_ACCESS_TOKEN")

// go routineの中で呼ぶこと
func PollingMessages() {
	pollingInterval := time.Minute * 5

	lastCheckpoint := time.Now()
	ticker := time.Tick(pollingInterval)

	for range ticker {
		now := time.Now()
		messages, err := collectMessages(lastCheckpoint, now)
		if err != nil {
			slog.Error(fmt.Sprintf("Failled to polling messages: %v", err))
			continue
		}

		lastCheckpoint = now

		slog.Info(fmt.Sprintf("Collect %d messages", len(messages.Hits)))
		// TODO: 取得したメッセージを使っての処理の呼び出し
	}
}

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
