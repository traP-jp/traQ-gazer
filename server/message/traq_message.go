package message

import (
	"fmt"
	"strings"
	"traQ-gazer/model"
	"traQ-gazer/repo"

	"golang.org/x/exp/slog"
)

func findMatchingWords(messageList model.MessageList) ([]*model.NotifyInfo, error) {
	notifyInfoList := make([]*model.NotifyInfo, 0)

	// メッセージごとに通知対象を検索する
	for _, messageItem := range messageList {
		// メッセージに含まれている登録単語で、通知条件が合致するものを登録者別にまとめる
		matchedWordsList, err := repo.GetMatchedWordList(messageItem)
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
