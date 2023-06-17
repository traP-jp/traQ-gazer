package traqmessage

import (
	"context"
	"fmt"
	"h23s_15/model"
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
		messageList, err := ConvertMessageHits(messages.Hits)
		if err != nil {
			slog.Error(fmt.Sprintf("Failled to convert messages: %v", err))
			continue
		}
		sendList, err := model.TraqMessageProcessor(messageList)
		if err != nil {
			slog.Error(fmt.Sprintf("Failled to process messages: %v", err))
			continue
		}
		for _, message := range sendList {
			err := sendMessage(message)
			if err != nil {
				slog.Error(fmt.Sprintf("Failled to send message: %v", err))
				continue
			}
		}
	}
}

func sendMessage(message model.Send) error {
	// TODO: 送信処理
	// 送信先: message.userId
	// 送信内容: "ワード:"+message.word+"\n https://q.trap.jp/messages/"+message.messageId
	v, _, err := t.client.MessageApi.PostDirectMessage(context.Background(), "userId").
		PostMessageRequest(traq.PostMessageRequest{
			// メッセージ本文
			Content: "",
			// // メンション・チャンネルリンクを自動埋め込みするか
			// Embed: check,
		}).
		Execute()
	slog.Info("%#v", v)
	if err != nil {
		slog.Info("%s", err)
	}
	return nil
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

func ConvertMessageHits(messages []traq.Message) (model.MessageList, error) {
	messageList := model.MessageList{}
	for _, message := range messages {
		messageList = append(messageList, model.MessageItem{
			Id:      message.Id,
			UserId:  message.UserId,
			Content: message.Content,
		})
	}
	return messageList, nil
}
