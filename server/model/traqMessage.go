package model

import "strings"

func TraqMessageProcessor(messageList MessageList) (SendList, error) {
	words := []Words{}
	err := db.Select(&words, "SELECT word, bot_notification, me_notification, trap_id FROM words")
	if err != nil {
		return nil, err
	}
	var sendList SendList
	for _, message := range messageList {
		for _, word := range words {
			if strings.Contains(message.Content, word.Word) {
				if message.UserId == word.UserId {
					if !word.IncludeMe {
						continue
					}
					if !word.IncludeBot {
						if IsBot(message.UserId) {
							continue
						}
					}
					// 通知する内容を追加
					sendList = append(sendList, Send{
						word:      word.Word,
						userId:    word.UserId,
						messageId: message.Id,
					})
				}
			}
		}
	}
	return sendList, nil
}

func IsBot(userId string) bool {
	// TODO: botかどうか判定する
	return false
}

type MessageItem struct {
	// メッセージUUID
	Id string `json:"id"`
	// 投稿者UUID
	UserId string `json:"userId"`
	// メッセージ本文
	Content string `json:"content"`
}

type MessageList []MessageItem

type Words struct {
	IncludeBot bool   `db:"bot_notification"`
	IncludeMe  bool   `db:"me_notification"`
	UserId     string `db:"trap_id"`
	Word       string `db:"word"`
}

type Send struct {
	// 含んでいた単語
	word string
	// 送信先
	userId string
	// 送信するメッセージのID
	messageId string
}

type SendList []Send
