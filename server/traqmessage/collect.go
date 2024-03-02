package traqmessage

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
	"traQ-gazer/model"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

type MessagePoller struct {
	processor *messageProcessor
}

func NewMessagePoller() *MessagePoller {
	return &MessagePoller{
		processor: &messageProcessor{
			queue: make(chan *[]traq.Message),
		},
	}
}

// go routineの中で呼ぶこと
func (m *MessagePoller) Run() {
	go m.processor.run()

	const pollingInterval = time.Minute * 3

	lastCheckpoint, err := model.GetPollingFrom()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get pollinginfo: %v", err))
		lastCheckpoint = time.Now()
	}

	var checkpointMutex sync.Mutex

	ticker := time.Tick(pollingInterval)
	for range ticker {
		slog.Info("Start polling")
		checkpointMutex.Lock()

		onPollingTime := time.Now()
		var collectedMessageCount int

		for page := 0; ; page++ {
			messages, more, err := collectMessages(lastCheckpoint, onPollingTime, page)

			if err != nil {
				slog.Error(fmt.Sprintf("Failed to polling messages: %v", err))
				lastCheckpoint = onPollingTime
				break
			}

			tmpMessageCount := len(*messages)

			// ページ0の初めのメッセージが最新のもの
			if page == 0 {
				lastCheckpoint = (*messages)[0].CreatedAt
			}

			slog.Info(fmt.Sprintf("Collected %d messages", tmpMessageCount))

			collectedMessageCount += tmpMessageCount

			// 取得したメッセージを使っての処理の呼び出し
			m.processor.enqueue(messages)
			if !more {
				if tmpMessageCount <= 0 {
					slog.Info("Message count is 0. Skip logging created at information")
				} else {
					slog.Info(fmt.Sprintf("The first one is created at %v.", (*messages)[tmpMessageCount-1].CreatedAt))
				}
				break
			}
		}

		slog.Info(fmt.Sprintf("%d messages collected totally", collectedMessageCount))

		err := model.RecordPollingTime(lastCheckpoint)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to recording lastCheckpoint: %v", err))
		}
		slog.Info(fmt.Sprintf("Now, lastCheckpoint = %v", lastCheckpoint))
		checkpointMutex.Unlock()
	}
}

// 通知メッセージの検索と通知処理のjobを処理する
type messageProcessor struct {
	queue chan *[]traq.Message
}

// go routineの中で呼ぶ
func (m *messageProcessor) run() {
	for {
		select {
		case messages := <-m.queue:
			m.process(*messages)
		}
	}
}

func (m *messageProcessor) enqueue(messages *[]traq.Message) {
	m.queue <- messages
}

func (m *messageProcessor) process(messages []traq.Message) {
	messageList, err := ConvertMessageHits(messages)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to convert messages: %v", err))
		return
	}
	notifyInfoList, err := model.FindMatchingWords(messageList)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to process messages: %v", err))
		return
	}

	slog.Info(fmt.Sprintf("Sending %d DMs...", len(notifyInfoList)))

	for _, notifyInfo := range notifyInfoList {
		err := sendMessage(notifyInfo.NotifyTargetTraqUuid, genNotifyMessageContent(notifyInfo.MessageId, notifyInfo.Words...))
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to send message: %v", err))
			continue
		}
	}

	slog.Info("End of send DMs")
}

func genNotifyMessageContent(citeMessageId string, words ...string) string {
	list := make([]string, 0)
	for _, word := range words {
		item := fmt.Sprintf("「%s」", word)
		list = append(list, item)
	}

	return fmt.Sprintf("%s\n https://q.trap.jp/messages/%s", strings.Join(list, ""), citeMessageId)
}

func sendMessage(notifyTargetTraqUUID string, messageContent string) error {
	if model.ACCESS_TOKEN == "" {
		slog.Info("Skip sendMessage")
		return nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, model.ACCESS_TOKEN)
	_, _, err := client.UserApi.PostDirectMessage(auth, notifyTargetTraqUUID).PostMessageRequest(traq.PostMessageRequest{
		Content: messageContent,
	}).Execute()
	if err != nil {
		slog.Info("Error sending message: %v", err)
		return err
	}
	return nil
}

func collectMessages(from time.Time, to time.Time, page int) (*[]traq.Message, bool, error) {
	if model.ACCESS_TOKEN == "" {
		slog.Info("Skip collectMessage")
		return &[]traq.Message{}, false, nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, model.ACCESS_TOKEN)

	// 1度での取得上限は100まで　それ以上はoffsetを使うこと
	// https://github.com/traPtitech/traQ/blob/47ed2cf94b2209c8444533326dee2a588936d5e0/service/search/engine.go#L51
	const limit = 100
	result, _, err := client.MessageApi.SearchMessages(auth).After(from).Before(to).Limit(limit).Offset(int32(limit * page)).Sort(`createdAt`).Execute()

	if err != nil {
		return nil, false, err
	}

	messages := result.Hits

	more := limit*(page+1) < int(result.TotalHits)
	return &messages, more, nil
}

func ConvertMessageHits(messages []traq.Message) (model.MessageList, error) {
	messageList := model.MessageList{}
	for _, message := range messages {
		messageList = append(messageList, model.MessageItem{
			Id:       message.Id,
			TraqUuid: message.UserId,
			Content:  message.Content,
		})
	}
	return messageList, nil
}
