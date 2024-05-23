package bot

import (
	"fmt"
	"strings"
	"traQ-gazer/db"
	"traQ-gazer/model"

	"golang.org/x/exp/slog"
)

func TraqMessageProcessor(messageList model.MessageList) (model.SendList, error) {
	wordsList, err := db.GetWordsWithoutTime()
	if err != nil {
		slog.Info("Error selecting words: %v", err)
		return nil, err
	}

	usersItem, err := db.GetUserList()
	if err != nil {
		slog.Info("Error selecting users: %v", err)
		return nil, err
	}

	traqUuidToTrapId := make(map[string]model.UsersItem)
	trapIdToTraqUuid := make(map[string]model.UsersItem)

	for _, item := range usersItem {
		trapIdToTraqUuid[item.TrapID] = item
		traqUuidToTrapId[item.TraqUUID] = item
	}

	var sendList model.SendList
	// TODO: Sotatsu リファクタリングと確認頼んだ！
	for _, message := range messageList {
		var messageOwnerTrapId string
		messageOwner, ok := traqUuidToTrapId[message.TraqUuid]
		if ok {
			messageOwnerTrapId = messageOwner.TrapID
		}

		for _, wordsItem := range wordsList {
			notifyTarget, ok := trapIdToTraqUuid[wordsItem.TrapId]
			if !ok {
				continue
			}
			if strings.Contains(message.Content, wordsItem.Word) {
				if !wordsItem.IncludeMe {
					if messageOwnerTrapId == notifyTarget.TrapID {
						continue
					}
				}

				if !wordsItem.IncludeBot {
					if messageOwner.IsBot {
						continue
					}
				}
				// 通知する内容を追加
				sendList = append(sendList, &model.Send{
					// wordがワードを登録しているUserの情報
					// messageが投稿されたワードの情報
					Word:                 wordsItem.Word,
					NotifyTargetTrapId:   notifyTarget.TrapID,
					NotifyTargetTraqUuid: notifyTarget.TraqUUID,
					MessageId:            message.Id,
					IsBot:                messageOwner.IsBot,
				})
			}
		}
	}
	// slog.Info("sendList: %v", sendList)
	return sendList, nil
}

func FindMatchingWords(messageList model.MessageList) ([]*model.NotifyInfo, error) {
	notifyInfoList := make([]*model.NotifyInfo, 0)

	// メッセージごとに通知対象を検索する
	for _, messageItem := range messageList {
		// メッセージに含まれている登録単語で、通知条件が合致するものを登録者別にまとめる
		matchedWordsList, err := db.GetMatchedWordList(messageItem)
		// err := db.Select(&matchedWordsList, `
		// 	SELECT
		// 		group_concat(words.word SEPARATOR '\n') AS contacted_words,
		// 		words.trap_id AS trap_id,
		// 		users.traq_uuid AS traq_uuid
		// 	FROM words
		// 	JOIN users ON words.trap_id = users.trap_id
		// 		WHERE (
		// 		    	((word NOT LIKE '/%/')
		// 		    		AND (? LIKE concat('%', word, '%')))
		//     			OR ((word LIKE '/%/')
		//     				AND (BINARY ? REGEXP trim(BOTH '/' FROM word)))
		// 		    )
		// 		AND (me_notification OR
		// 			 users.traq_uuid != ?)
		// 		AND (bot_notification OR
		// 			 (SELECT is_bot FROM users WHERE traq_uuid = ? LIMIT 1) = FALSE)
		// 	GROUP BY words.trap_id`,
		// 	messageItem.Content, messageItem.Content, messageItem.TraqUuid, messageItem.TraqUuid)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to search words with message: `%s`", messageItem.Id))
			return nil, err
		}

		for _, matchedWords := range matchedWordsList {
			notifyInfo := &model.NotifyInfo{
				Words:                strings.Split(matchedWords.ContactedWords, "\n"),
				NotifyTargetTrapId:   matchedWords.TrapID,
				NotifyTargetTraqUuid: matchedWords.TraqUUID,
				MessageId:            messageItem.Id,
			}

			notifyInfoList = append(notifyInfoList, notifyInfo)
		}
	}

	return notifyInfoList, nil
}
