package model

import "strings"

func TraqMessageProcessor(messageList MessageList) (SendList, error) {
	wordsItem := []WordsItem{}
	err := db.Select(&wordsItem, `
		SELECT 
			word, bot_notification, me_notification, trap_id 
		FROM
			words`)
	if err != nil {
		return nil, err
	}
	wordsItemMap := make(map[string]WordsItem)
	for _, item := range wordsItem {
		wordsItemMap[item.UserId] = item
	}

	usersItem := []UsersItem{}
	err = db.Select(&usersItem, `
		SELECT 
			traq_uuid, traq_id, is_bot 
		FROM
			users`)
	if err != nil {
		return nil, err
	}
	usersItemMap := make(map[string]UsersItem)
	for _, item := range usersItem {
		usersItemMap[item.UserId] = item
	}

	var sendList SendList
	// TODO: Sotatsu リファクタリングと確認頼んだ！
	for _, message := range messageList {
		for _, wordsItem := range wordsItemMap {
			if strings.Contains(message.Content, wordsItem.Word) {
				if message.UserId == wordsItem.UserId {
					if !wordsItem.IncludeMe {
						continue
					}
				}
				if !wordsItem.IncludeBot {
					if usersItemMap[message.UserId].IsBot {
						continue
					}
				}
				// 通知する内容を追加
				sendList = append(sendList, &Send{
					// wordがワードを登録しているUserの情報
					// messageが投稿されたワードの情報
					Word:      wordsItem.Word,
					UserId:    wordsItem.UserId,
					UserUUID:  usersItemMap[wordsItem.UserId].UserUUID,
					MessageId: message.Id,
					IsBot:     usersItemMap[wordsItem.UserId].IsBot,
				})
			}
		}
	}
	return sendList, nil
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

type WordsItem struct {
	IncludeBot bool   `db:"bot_notification"`
	IncludeMe  bool   `db:"me_notification"`
	UserId     string `db:"trap_id"`
	Word       string `db:"word"`
}

type UsersItem struct {
	UserId   string `db:"trap_id"`
	UserUUID string `db:"traq_uuid"`
	IsBot    bool   `db:"is_bot"`
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
