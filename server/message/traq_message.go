package message

import (
	"fmt"
	"strings"
	"traQ-gazer/model"
	"traQ-gazer/repo"

	"golang.org/x/exp/slog"
)

type messageWordMatcher interface {
	matchMessage(model.MessageItem) []model.MatchedWords
}

type wordMatcherLoaderFunc func() (messageWordMatcher, error)

func findMatchingWords(messageList model.MessageList, loadMatcher wordMatcherLoaderFunc) ([]*model.NotifyInfo, error) {
	if len(messageList) == 0 {
		return nil, nil
	}

	notifyInfoList := make([]*model.NotifyInfo, 0)
	matcher, err := loadMatcher()
	if err != nil {
		return nil, err
	}

	// メッセージごとに通知対象を検索する
	for _, messageItem := range messageList {
		// メッセージに含まれている登録単語で、通知条件が合致するものを登録者別にまとめる
		matchedWordsList := matcher.matchMessage(messageItem)

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

func loadWordMatcher() (messageWordMatcher, error) {
	words, err := repo.GetWordsWithoutTime()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to fetch words for matching: %v", err))
		return nil, err
	}

	users, err := repo.GetUserList()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to fetch users for matching: %v", err))
		return nil, err
	}

	matcher, err := newWordMatcher(words, users)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to build word matcher: %v", err))
		return nil, err
	}
	return matcher, nil
}
