package model

import "strings"

func TraqMessageProcessor(messageList MessageList) (SendList, error) {
	words := []Words{}
	err := db.Select(&words, `
		SELECT 
			words.word, words.bot_notification, words.me_notification, words.trap_id, users.traq_uuid, users.is_bot 
		FROM
			words 
		INNER JOIN 
			users 
		ON 
			users.trap_id = words.trap_id`)
	if err != nil {
		return nil, err
	}
	var sendList SendList
	// TODO: Sotatsu リファクタリングと確認頼んだ！
	for _, message := range messageList {
		for _, word := range words {
			if strings.Contains(message.Content, word.Word) {
				if message.UserId == word.UserId {
					if !word.IncludeMe {
						continue
					}
				}
				if !word.IncludeBot {
					isBot, err := IsBot(message.UserId)
					if err != nil {
						continue
					}
					if isBot {
						continue
					}
				}
				// 通知する内容を追加
				sendList = append(sendList, &Send{
					// wordがワードを登録しているUserの情報
					// messageが投稿されたワードの情報
					Word:      word.Word,
					UserId:    word.UserId,
					UserUUID:  word.UserUUID,
					MessageId: message.Id,
					IsBot:     word.IsBot,
				})
			}
		}
	}
	return sendList, nil
}

func IsBot(userId string) (bool, error) {
	// TODO: botかどうか判定する
	var isBot IsBotStruct
	err := db.Get(&isBot, "SELECT is_bot FROM users WHERE user_id = ?", userId)
	if err != nil {
		return true, err
	}
	return isBot.IsBot, nil
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
	UserUUID   string `db:"traq_uuid"`
	Word       string `db:"word"`
	IsBot      bool   `db:"is_bot"`
}

type Send struct {
	// 含んでいた単語
	Word string
	// 送信先のuser
	UserId string
	// 送信先のuserUUID
	UserUUID string
	// 送信するメッセージのID
	MessageId string
	// BOTかどうか
	IsBot bool
}

type SendList []*Send

type IsBotStruct struct {
	IsBot bool `db:"is_bot"`
}
