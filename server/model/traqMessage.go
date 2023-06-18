package model

import "strings"

func TraqMessageProcessor(messageList MessageList) (SendList, error) {
	words := []Words{}
	err := db.Select(&words, "SELECT words.word, words.bot_notification, words.me_notification, words.trap_id, users.traq_uuid, users.is_bot, users.dm_id FROM words INNER JOIN users ON users.trap_id = words.trap_id")
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
					word:      word.Word,
					userId:    word.UserId,
					userUUID:  word.UserUUID,
					messageId: message.Id,
					isBot:     word.IsBot,
					dmId:      word.DmId,
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
	DmId       string `db:"dm_id"`
}

type Send struct {
	// 含んでいた単語
	word string
	// 送信先のuser
	userId string
	// 送信先のuserUUID
	userUUID string
	// 送信するメッセージのID
	messageId string
	// BOTかどうか
	isBot bool
	// 送信先のDMのID
	dmId string
}

type SendList []*Send

type IsBotStruct struct {
	IsBot bool `db:"is_bot"`
}
