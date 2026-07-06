package message

import (
	"regexp"
	"strings"

	"traQ-gazer/model"
	"traQ-gazer/wordmatch"

	"golang.org/x/exp/slog"
)

type wordMatcher struct {
	senderIsBotByTraqUUID map[string]bool
	targets               []wordMatchTarget
}

type wordMatchTarget struct {
	includeBot bool
	includeMe  bool
	isRegex    bool
	plainWord  string
	regex      *regexp.Regexp
	trapID     string
	traqUUID   string
	word       string
}

func newWordMatcher(words []model.WordsItem, users []model.UsersItem) (*wordMatcher, error) {
	usersByTrapID := make(map[string]model.UsersItem, len(users))
	senderIsBotByTraqUUID := make(map[string]bool, len(users))
	for _, user := range users {
		usersByTrapID[user.TrapID] = user
		senderIsBotByTraqUUID[user.TraqUUID] = user.IsBot
	}

	targets := make([]wordMatchTarget, 0, len(words))
	for _, word := range words {
		user, exists := usersByTrapID[word.TrapId]
		if !exists {
			continue
		}

		target, err := newWordMatchTarget(word, user)
		if err != nil {
			slog.Error("skip invalid registered regex word")
			continue
		}
		targets = append(targets, target)
	}

	return &wordMatcher{
		senderIsBotByTraqUUID: senderIsBotByTraqUUID,
		targets:               targets,
	}, nil
}

func newWordMatchTarget(word model.WordsItem, user model.UsersItem) (wordMatchTarget, error) {
	target := wordMatchTarget{
		includeBot: word.IncludeBot,
		includeMe:  word.IncludeMe,
		trapID:     word.TrapId,
		traqUUID:   user.TraqUUID,
		word:       word.Word,
	}

	if wordmatch.IsRegexWord(word.Word) {
		regex, err := wordmatch.CompileRegexWord(word.Word)
		if err != nil {
			return wordMatchTarget{}, err
		}
		target.isRegex = true
		target.regex = regex
		return target, nil
	}

	target.plainWord = wordmatch.NormalizePlainWord(word.Word)
	return target, nil
}

func (m *wordMatcher) matchMessage(messageItem model.MessageItem) []model.MatchedWords {
	wordsByTrapID := map[string][]string{}
	targetsByTrapID := map[string]model.MatchedWords{}
	trapIDOrder := []string{}

	content := wordmatch.NormalizePlainWord(messageItem.Content)
	for _, target := range m.targets {
		if !target.matches(messageItem.Content, content) ||
			!target.allowsSelfNotification(messageItem.TraqUuid) ||
			!target.allowsBotNotification(m.senderIsBotByTraqUUID, messageItem.TraqUuid) {
			continue
		}

		if _, exists := targetsByTrapID[target.trapID]; !exists {
			trapIDOrder = append(trapIDOrder, target.trapID)
			targetsByTrapID[target.trapID] = model.MatchedWords{
				TrapID:   target.trapID,
				TraqUUID: target.traqUUID,
			}
		}
		wordsByTrapID[target.trapID] = append(wordsByTrapID[target.trapID], target.word)
	}

	matchedWordsList := make([]model.MatchedWords, 0, len(trapIDOrder))
	for _, trapID := range trapIDOrder {
		matchedWords := targetsByTrapID[trapID]
		matchedWords.ContactedWords = strings.Join(wordsByTrapID[trapID], "\n")
		matchedWordsList = append(matchedWordsList, matchedWords)
	}
	return matchedWordsList
}

func (t wordMatchTarget) matches(rawContent, lowerContent string) bool {
	if t.isRegex {
		return t.regex.MatchString(rawContent)
	}
	return strings.Contains(lowerContent, t.plainWord)
}

func (t wordMatchTarget) allowsSelfNotification(senderTraqUUID string) bool {
	return t.includeMe || t.traqUUID != senderTraqUUID
}

func (t wordMatchTarget) allowsBotNotification(senderIsBotByTraqUUID map[string]bool, senderTraqUUID string) bool {
	if t.includeBot {
		return true
	}
	senderIsBot, exists := senderIsBotByTraqUUID[senderTraqUUID]
	return exists && !senderIsBot
}
