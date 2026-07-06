package message

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"sync"
	"time"
	"traQ-gazer/model"
	"traQ-gazer/repo"

	"github.com/traPtitech/go-traq"
)

type MessagePoller struct {
	processor *messageProcessor
}

func NewMessagePoller() *MessagePoller {
	return &MessagePoller{
		processor: &messageProcessor{
			queue:                       make(chan *[]traq.Message),
			loadNotificationWordMatcher: loadNotificationWordMatcher,
		},
	}
}

// go routineの中で呼ぶこと
func (m *MessagePoller) Run() {
	go m.processor.run()

	const pollingInterval = time.Minute * 3

	lastCheckpoint, err := repo.GetPollingFrom()
	if err != nil {
		slog.Error("failed to get polling info", "err", err)
		lastCheckpoint = time.Now()
	}

	var checkpointMutex sync.Mutex

	ticker := time.Tick(pollingInterval)
	for range ticker {
		slog.Info("start polling", "from", lastCheckpoint)
		checkpointMutex.Lock()

		onPollingTime := time.Now()
		var collectedMessageCount int

		var allMessages []traq.Message

		for page := 0; ; page++ {
			messages, more, err := collectMessages(lastCheckpoint, onPollingTime, page)

			if err != nil {
				slog.Error("failed to poll messages", "from", lastCheckpoint, "to", onPollingTime, "page", page, "err", err)
				break
			}

			tmpMessageCount := len(*messages)

			// ページ0の初めのメッセージが最新のもの
			if page == 0 && tmpMessageCount != 0 {
				lastCheckpoint = (*messages)[0].CreatedAt
			}

			slog.Info("collected messages", "count", tmpMessageCount, "page", page)

			collectedMessageCount += tmpMessageCount

			// 取得したメッセージをallMessagesに追加
			allMessages = append(allMessages, *messages...)
			if !more {
				if tmpMessageCount <= 0 {
					slog.Info("no messages collected")
				} else {
					slog.Info("oldest collected message", "created_at", (*messages)[tmpMessageCount-1].CreatedAt)
				}
				break
			}
		}

		// 通知処理にメッセージを渡す
		m.processor.enqueue(&allMessages)

		slog.Info("messages collected", "count", collectedMessageCount)

		err := repo.RecordPollingTime(lastCheckpoint)
		if err != nil {
			slog.Error("failed to record last checkpoint", "last_checkpoint", lastCheckpoint, "err", err)
		}
		slog.Info("updated last checkpoint", "last_checkpoint", lastCheckpoint)
		checkpointMutex.Unlock()
	}
}

// 通知メッセージの検索と通知処理のjobを処理する
type messageProcessor struct {
	queue                       chan *[]traq.Message
	loadNotificationWordMatcher notificationWordMatcherLoaderFunc
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
	messageList, err := convertMessageHits(messages)
	if err != nil {
		slog.Error("failed to convert messages", "err", err)
		return
	}
	notifyInfoList, err := findMatchingWords(messageList, m.notificationWordMatcherLoader())
	if err != nil {
		slog.Error("failed to process messages", "err", err)
		return
	}

	slog.Info("sending direct messages", "count", len(notifyInfoList))

	// 元投稿の時系列に沿ってDMを送るためにnotifyInfoListをIdでソート
	sort.Slice(notifyInfoList, func(i, j int) bool {
		return notifyInfoList[i].MessageId < notifyInfoList[j].MessageId
	})

	for _, notifyInfo := range notifyInfoList {
		err := sendMessage(notifyInfo.NotifyTargetTraqUuid, genNotifyMessageContent(notifyInfo.MessageId, notifyInfo.Words...))
		if err != nil {
			slog.Error("failed to send direct message", "err_type", fmt.Sprintf("%T", err))
			continue
		}
	}

	slog.Info("finished sending direct messages")
}

func (m *messageProcessor) notificationWordMatcherLoader() notificationWordMatcherLoaderFunc {
	if m.loadNotificationWordMatcher != nil {
		return m.loadNotificationWordMatcher
	}
	return loadNotificationWordMatcher
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
	if repo.AccessToken == "" {
		slog.Info("skip send message")
		return nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, repo.AccessToken)
	_, _, err := client.UserApi.PostDirectMessage(auth, notifyTargetTraqUUID).PostMessageRequest(traq.PostMessageRequest{
		Content: messageContent,
	}).Execute()
	if err != nil {
		return err
	}
	return nil
}

func collectMessages(from time.Time, to time.Time, page int) (*[]traq.Message, bool, error) {
	if repo.AccessToken == "" {
		slog.Info("skip collect messages")
		return &[]traq.Message{}, false, nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, repo.AccessToken)

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

func convertMessageHits(messages []traq.Message) (model.MessageList, error) {
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
