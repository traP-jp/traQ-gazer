package traqmessage

import (
	"context"
	"fmt"
	"h23s_15/model"
	"time"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

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
		messageProcessor(messages.Hits)
	}
}

func messageProcessor(messages []traq.Message) {
	check := true
	messageList, err := ConvertMessageHits(messages)
	if err != nil {
		slog.Error(fmt.Sprintf("Failled to convert messages: %v", err))
		check = false
	}
	sendList, err := model.TraqMessageProcessor(messageList)
	if err != nil {
		slog.Error(fmt.Sprintf("Failled to process messages: %v", err))
		check = false
	}
	if check {
		for _, message := range sendList {
			err := sendMessage(*message)
			if err != nil {
				slog.Error(fmt.Sprintf("Failled to send message: %v", err))
				continue
			}
		}
	}
}

func sendMessage(message model.Send) error {
	// TODO: 送信処理
	// 送信先User: message.userId 送信先: message.dmId
	// 送信内容: "ワード:"+message.word+"\n https://q.trap.jp/messages/"+message.messageId
	// client := traq.NewAPIClient(traq.NewConfiguration())
	// auth := context.WithValue(context.Background(), traq.ContextAccessToken, model.ACCESS_TOKEN)
	// _, _, err := client.UserApi.PostDirectMessage(auth, message.userUUID).PostMessageRequest(&Post
	// )
	// if err != nil {
	// 	return err
	// }
	return nil
}

func collectMessages(from time.Time, to time.Time) (*traq.MessageSearchResult, error) {
	if model.ACCESS_TOKEN == "" {
		slog.Info("Skip collectMessage")
		return &traq.MessageSearchResult{}, nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, model.ACCESS_TOKEN)

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
