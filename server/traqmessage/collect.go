package traqmessage

import (
	"context"
	"fmt"
	"h23s_15/model"
	"strings"
	"sync"
	"time"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

type MessagePoller struct {
	processor *messageProcessor
}

func NewMessagePoller() *MessagePoller {
	return &MessagePoller{
		processor: &messageProcessor{
			queue:            make(chan *[]traq.Message),
			lastCheckMessage: "",
		},
	}
}

// go routineの中で呼ぶこと
func (m *MessagePoller) Run() {
	go m.processor.run()

	pollingInterval := time.Minute * 3

	lastCheckpoint := time.Now()
	var checkpointMutex sync.Mutex

	ticker := time.Tick(pollingInterval)
	for range ticker {
		slog.Info("Start polling")
		checkpointMutex.Lock()

		now := time.Now()
		collectedMessageCount := 0
		for {
			messages, tmplastmessage, err := collectMessages(lastCheckpoint, now, collectedMessageCount)
			if err != nil {
				slog.Error(fmt.Sprintf("Failled to polling messages: %v", err))
				break
			}

			// オフセット0の時なら検索対象最新メッセージが真に最新メッセージ
			if collectedMessageCount == 0 {
				m.processor.lastCheckMessage = tmplastmessage
			}

			tmpMessageCount := len(messages.Hits)

			slog.Info(fmt.Sprintf("Collect %d messages", tmpMessageCount))
			collectedMessageCount += tmpMessageCount

			// 取得したメッセージを使っての処理の呼び出し
			m.processor.enqueue(&messages.Hits)

			if tmpMessageCount < 100 {
				break
			}
		}

		slog.Info(fmt.Sprintf("%d messages collected totally", collectedMessageCount))

		lastCheckpoint = now
		checkpointMutex.Unlock()
	}
}

// 通知メッセージの検索と通知処理のjobを処理する
type messageProcessor struct {
	queue            chan *[]traq.Message
	lastCheckMessage string //前回ポーリング時の最新メッセージUUID
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
	messageList, err := ConvertMessageHits(messages, m.lastCheckMessage)
	if err != nil {
		slog.Error(fmt.Sprintf("Failled to convert messages: %v", err))
		return
	}
	notifyInfoList, err := model.FindMatchingWords(messageList)
	if err != nil {
		slog.Error(fmt.Sprintf("Failled to process messages: %v", err))
		return
	}

	slog.Info(fmt.Sprintf("Sending %d DMs...", len(notifyInfoList)))

	for _, notifyInfo := range notifyInfoList {
		err := sendMessage(notifyInfo.NotifyTargetTraqUuid, genNotifyMessageContent(notifyInfo.MessageId, notifyInfo.Words...))
		if err != nil {
			slog.Error(fmt.Sprintf("Failled to send message: %v", err))
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

func collectMessages(from time.Time, to time.Time, offset int) (*traq.MessageSearchResult, string, error) {
	if model.ACCESS_TOKEN == "" {
		slog.Info("Skip collectMessage")
		return &traq.MessageSearchResult{}, "", nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, model.ACCESS_TOKEN)

	// 1度での取得上限は100まで　それ以上はoffsetを使うこと
	// ポーリング漏れ防止のために1分余分にメッセージを取得
	// https://github.com/traPtitech/traQ/blob/47ed2cf94b2209c8444533326dee2a588936d5e0/service/search/engine.go#L51
	result, _, err := client.MessageApi.SearchMessages(auth).After(from.Add(-time.Minute)).Before(to).Limit(100).Offset(int32(offset)).Sort(`createdAt`).Execute()
	if err != nil {
		return nil, "", err
	}

	lastCheckMessage := ""
	if offset == 0 {
		lastCheckMessage = result.Hits[0].Id
	}

	return result, lastCheckMessage, nil
}

func ConvertMessageHits(messages []traq.Message, lastcheckmessage string) (model.MessageList, error) {
	messageList := model.MessageList{}
	for _, message := range messages {
		if message.Id == lastcheckmessage {
			break
		}
		messageList = append(messageList, model.MessageItem{
			Id:       message.Id,
			TraqUuid: message.UserId,
			Content:  message.Content,
		})
	}
	return messageList, nil
}
