package message

import (
	"errors"
	"log/slog"
	"strings"

	"traQ-gazer/model"
	"traQ-gazer/wordpattern"
)

// notificationWordMatcher applies registered words to messages and notification settings.
type notificationWordMatcher struct {
	senderIsBotByTraqUUID map[string]bool
	targets               []notificationWordTarget
}

type registeredWord interface {
	matches(messageContent) bool
	text() string
}

type closeableRegisteredWord interface {
	close() error
}

type messageContent struct {
	raw        string
	normalized string
}

type messageSender struct {
	traqUUID string
	isKnown  bool
	isBot    bool
}

type notificationWordTarget struct {
	includeBot bool
	includeMe  bool
	trapID     string
	traqUUID   string
	word       registeredWord
}

type plainRegisteredWord struct {
	original   string
	normalized string
}

type regexRegisteredWord struct {
	original string
	regex    *wordpattern.RegexWord
}

func newNotificationWordMatcher(words []model.WordsItem, users []model.UsersItem) (*notificationWordMatcher, error) {
	usersByTrapID := make(map[string]model.UsersItem, len(users))
	senderIsBotByTraqUUID := make(map[string]bool, len(users))
	for _, user := range users {
		usersByTrapID[user.TrapID] = user
		senderIsBotByTraqUUID[user.TraqUUID] = user.IsBot
	}

	targets := make([]notificationWordTarget, 0, len(words))
	for _, word := range words {
		user, exists := usersByTrapID[word.TrapId]
		if !exists {
			continue
		}

		target, err := newNotificationWordTarget(word, user)
		if err != nil {
			slog.Warn("skip invalid registered regex word")
			continue
		}
		targets = append(targets, target)
	}

	return &notificationWordMatcher{
		senderIsBotByTraqUUID: senderIsBotByTraqUUID,
		targets:               targets,
	}, nil
}

func newNotificationWordTarget(word model.WordsItem, user model.UsersItem) (notificationWordTarget, error) {
	matchedWord, err := newRegisteredWord(word.Word)
	if err != nil {
		return notificationWordTarget{}, err
	}

	target := notificationWordTarget{
		includeBot: word.IncludeBot,
		includeMe:  word.IncludeMe,
		trapID:     word.TrapId,
		traqUUID:   user.TraqUUID,
		word:       matchedWord,
	}
	return target, nil
}

func newRegisteredWord(wordText string) (registeredWord, error) {
	if wordpattern.IsRegexWord(wordText) {
		regex, err := wordpattern.CompileRegexWord(wordText)
		if err != nil {
			return nil, err
		}
		return regexRegisteredWord{original: wordText, regex: regex}, nil
	}

	return plainRegisteredWord{
		original:   wordText,
		normalized: wordpattern.NormalizePlainWord(wordText),
	}, nil
}

func (m *notificationWordMatcher) close() error {
	var errs []error
	for _, target := range m.targets {
		if err := target.close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (m *notificationWordMatcher) matchMessage(messageItem model.MessageItem) []model.MatchedWords {
	content := newMessageContent(messageItem.Content)
	sender := m.messageSender(messageItem.TraqUuid)
	matchedTargets := targetsMatchingContent(m.targets, content)
	allowedTargets := targetsAllowedForSender(matchedTargets, sender)
	return matchedWordsFromTargets(allowedTargets)
}

func newMessageContent(raw string) messageContent {
	return messageContent{
		raw:        raw,
		normalized: wordpattern.NormalizePlainWord(raw),
	}
}

func (m *notificationWordMatcher) messageSender(traqUUID string) messageSender {
	isBot, exists := m.senderIsBotByTraqUUID[traqUUID]
	return messageSender{
		traqUUID: traqUUID,
		isKnown:  exists,
		isBot:    isBot,
	}
}

func targetsMatchingContent(targets []notificationWordTarget, content messageContent) []notificationWordTarget {
	matchedTargets := make([]notificationWordTarget, 0, len(targets))
	for _, target := range targets {
		if target.matchesContent(content) {
			matchedTargets = append(matchedTargets, target)
		}
	}
	return matchedTargets
}

func targetsAllowedForSender(targets []notificationWordTarget, sender messageSender) []notificationWordTarget {
	allowedTargets := make([]notificationWordTarget, 0, len(targets))
	for _, target := range targets {
		if target.allowsSender(sender) {
			allowedTargets = append(allowedTargets, target)
		}
	}
	return allowedTargets
}

func matchedWordsFromTargets(targets []notificationWordTarget) []model.MatchedWords {
	wordsByTrapID := map[string][]string{}
	targetsByTrapID := map[string]model.MatchedWords{}
	trapIDOrder := []string{}

	for _, target := range targets {
		if _, exists := targetsByTrapID[target.trapID]; !exists {
			trapIDOrder = append(trapIDOrder, target.trapID)
			targetsByTrapID[target.trapID] = model.MatchedWords{
				TrapID:   target.trapID,
				TraqUUID: target.traqUUID,
			}
		}
		wordsByTrapID[target.trapID] = append(wordsByTrapID[target.trapID], target.word.text())
	}

	matchedWordsList := make([]model.MatchedWords, 0, len(trapIDOrder))
	for _, trapID := range trapIDOrder {
		matchedWords := targetsByTrapID[trapID]
		matchedWords.ContactedWords = strings.Join(wordsByTrapID[trapID], "\n")
		matchedWordsList = append(matchedWordsList, matchedWords)
	}
	return matchedWordsList
}

func (t notificationWordTarget) close() error {
	if closeableWord, ok := t.word.(closeableRegisteredWord); ok {
		return closeableWord.close()
	}
	return nil
}

func (t notificationWordTarget) matchesContent(content messageContent) bool {
	if t.word == nil {
		return false
	}
	return t.word.matches(content)
}

func (t notificationWordTarget) allowsSender(sender messageSender) bool {
	return t.allowsSelfNotification(sender) && t.allowsBotNotification(sender)
}

func (t notificationWordTarget) allowsSelfNotification(sender messageSender) bool {
	return t.includeMe || t.traqUUID != sender.traqUUID
}

func (t notificationWordTarget) allowsBotNotification(sender messageSender) bool {
	if t.includeBot {
		return true
	}
	return sender.isKnown && !sender.isBot
}

func (w plainRegisteredWord) matches(content messageContent) bool {
	return strings.Contains(content.normalized, w.normalized)
}

func (w plainRegisteredWord) text() string {
	return w.original
}

func (w regexRegisteredWord) close() error {
	if w.regex == nil {
		return nil
	}
	return w.regex.Close()
}

func (w regexRegisteredWord) matches(content messageContent) bool {
	matched, err := w.regex.MatchString(content.raw)
	if err != nil {
		slog.Warn("skip regex word after match error")
		return false
	}
	return matched
}

func (w regexRegisteredWord) text() string {
	return w.original
}
