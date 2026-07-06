package message

import (
	"errors"
	"reflect"
	"testing"

	"traQ-gazer/model"
	"traQ-gazer/wordpattern"
)

func TestFindMatchingWords(t *testing.T) {
	t.Parallel()

	t.Run("empty message list", func(t *testing.T) {
		t.Parallel()

		got, err := findMatchingWords(model.MessageList{}, func() (messageWordMatcher, error) {
			t.Fatal("word matcher loader should not be called for empty messages")
			return nil, nil
		})
		if err != nil {
			t.Fatalf("findMatchingWords returned error: %v", err)
		}
		if got != nil {
			t.Fatalf("notify info list = %#v, want nil", got)
		}
	})

	t.Run("matcher loading", func(t *testing.T) {
		t.Parallel()

		matcher := &fakeMessageWordMatcher{
			matchedWordsList: []model.MatchedWords{
				{
					ContactedWords: "hello\nworld",
					TrapID:         "target",
					TraqUUID:       "target-uuid",
				},
			},
		}

		got, err := findMatchingWords(
			model.MessageList{
				{
					Id:       "message-id",
					TraqUuid: "sender-uuid",
					Content:  "message content",
				},
			},
			func() (messageWordMatcher, error) {
				return matcher, nil
			},
		)
		if err != nil {
			t.Fatalf("findMatchingWords returned error: %v", err)
		}

		want := []*model.NotifyInfo{
			{
				Words:                []string{"hello", "world"},
				NotifyTargetTrapId:   "target",
				NotifyTargetTraqUuid: "target-uuid",
				MessageId:            "message-id",
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("notify info list = %#v, want %#v", got, want)
		}
		if matcher.callCount != 1 {
			t.Fatalf("matcher call count = %d, want 1", matcher.callCount)
		}
	})

	t.Run("matcher loading error", func(t *testing.T) {
		t.Parallel()

		wantErr := errors.New("load failed")

		got, err := findMatchingWords(
			model.MessageList{
				{
					Id:       "message-id",
					TraqUuid: "sender-uuid",
					Content:  "message content",
				},
			},
			func() (messageWordMatcher, error) {
				return nil, wantErr
			},
		)

		if !errors.Is(err, wantErr) {
			t.Fatalf("error = %v, want %v", err, wantErr)
		}
		if got != nil {
			t.Fatalf("notify info list = %#v, want nil", got)
		}
	})
}

func TestWordMatcher_MatchMessage(t *testing.T) {
	t.Parallel()

	matcher, err := newWordMatcher(
		[]model.WordsItem{
			{Word: "hello", TrapId: "target-a", IncludeBot: false, IncludeMe: true},
			{Word: "self", TrapId: "sender", IncludeBot: true, IncludeMe: false},
			{Word: "missing", TrapId: "target-b", IncludeBot: true, IncludeMe: true},
		},
		[]model.UsersItem{
			{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
			{TrapID: "target-a", TraqUUID: "target-a-uuid", IsBot: false},
			{TrapID: "target-b", TraqUUID: "target-b-uuid", IsBot: false},
		},
	)
	if err != nil {
		t.Fatalf("newWordMatcher returned error: %v", err)
	}
	t.Cleanup(func() {
		if err := matcher.close(); err != nil {
			t.Errorf("wordMatcher.close returned error: %v", err)
		}
	})

	got := matcher.matchMessage(model.MessageItem{
		TraqUuid: "sender-uuid",
		Content:  "hello, self",
	})
	want := []model.MatchedWords{
		{
			ContactedWords: "hello",
			TrapID:         "target-a",
			TraqUUID:       "target-a-uuid",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("matched words = %#v, want %#v", got, want)
	}
}

func TestTargetsMatchingContent(t *testing.T) {
	t.Parallel()

	t.Run("message content filtering", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWords  []string
			messageContent   string
			wantMatchedWords []string
		}{
			"only registered words found in the message remain": {
				registeredWords:  []string{"hello", "missing", "world"},
				messageContent:   "hello world",
				wantMatchedWords: []string{"hello", "world"},
			},
			"registered words absent from the message are dropped": {
				registeredWords:  []string{"first", "second"},
				messageContent:   "content",
				wantMatchedWords: []string{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				registeredTargets := make([]wordMatchTarget, 0, len(tt.registeredWords))
				for _, word := range tt.registeredWords {
					registeredTargets = append(registeredTargets, wordMatchTarget{
						word: plainRegisteredWord{original: word, normalized: word},
					})
				}

				gotTargets := targetsMatchingContent(registeredTargets, newMessageContent(tt.messageContent))
				got := make([]string, 0, len(gotTargets))
				for _, target := range gotTargets {
					got = append(got, target.word.text())
				}

				if !reflect.DeepEqual(got, tt.wantMatchedWords) {
					t.Fatalf("matched target words = %#v, want %#v", got, tt.wantMatchedWords)
				}
			})
		}
	})

	t.Run("registration order", func(t *testing.T) {
		t.Parallel()

		gotTargets := targetsMatchingContent(
			[]wordMatchTarget{
				{word: plainRegisteredWord{original: "first", normalized: "first"}},
				{word: plainRegisteredWord{original: "second", normalized: "second"}},
				{word: plainRegisteredWord{original: "third", normalized: "third"}},
			},
			newMessageContent("third first second"),
		)
		got := make([]string, 0, len(gotTargets))
		for _, target := range gotTargets {
			got = append(got, target.word.text())
		}
		want := []string{"first", "second", "third"}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("matched target words = %#v, want %#v", got, want)
		}
	})
}

func TestTargetsAllowedForSender(t *testing.T) {
	t.Parallel()

	t.Run("self notification setting", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			includeMe      bool
			senderTraqUUID string
			wantAllowed    bool
		}{
			"own message is allowed when includeMe is true": {
				includeMe:      true,
				senderTraqUUID: "target-uuid",
				wantAllowed:    true,
			},
			"own message is dropped when includeMe is false": {
				includeMe:      false,
				senderTraqUUID: "target-uuid",
				wantAllowed:    false,
			},
			"another user's message is allowed even when includeMe is false": {
				includeMe:      false,
				senderTraqUUID: "sender-uuid",
				wantAllowed:    true,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				notifyTarget := wordMatchTarget{
					includeMe:  tt.includeMe,
					includeBot: true,
					trapID:     "target",
					traqUUID:   "target-uuid",
					word:       plainRegisteredWord{original: "keyword", normalized: "keyword"},
				}

				gotTargets := targetsAllowedForSender(
					[]wordMatchTarget{notifyTarget},
					messageSender{traqUUID: tt.senderTraqUUID, isKnown: true},
				)
				gotAllowed := len(gotTargets) == 1
				if gotAllowed != tt.wantAllowed {
					t.Fatalf("target allowed = %v, want %v", gotAllowed, tt.wantAllowed)
				}
			})
		}
	})

	t.Run("bot notification setting", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			includeBot  bool
			sender      messageSender
			wantAllowed bool
		}{
			"unknown sender is allowed when includeBot is true": {
				includeBot:  true,
				sender:      messageSender{traqUUID: "unknown-uuid"},
				wantAllowed: true,
			},
			"known human sender is allowed when includeBot is false": {
				includeBot:  false,
				sender:      messageSender{traqUUID: "human-uuid", isKnown: true, isBot: false},
				wantAllowed: true,
			},
			"known bot sender is dropped when includeBot is false": {
				includeBot:  false,
				sender:      messageSender{traqUUID: "bot-uuid", isKnown: true, isBot: true},
				wantAllowed: false,
			},
			"unknown sender is dropped when includeBot is false": {
				includeBot:  false,
				sender:      messageSender{traqUUID: "unknown-uuid"},
				wantAllowed: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				notifyTarget := wordMatchTarget{
					includeMe:  true,
					includeBot: tt.includeBot,
					trapID:     "target",
					traqUUID:   "target-uuid",
					word:       plainRegisteredWord{original: "keyword", normalized: "keyword"},
				}

				gotTargets := targetsAllowedForSender([]wordMatchTarget{notifyTarget}, tt.sender)
				gotAllowed := len(gotTargets) == 1
				if gotAllowed != tt.wantAllowed {
					t.Fatalf("target allowed = %v, want %v", gotAllowed, tt.wantAllowed)
				}
			})
		}
	})
}

func TestMatchedWordsFromTargets(t *testing.T) {
	t.Parallel()

	got := matchedWordsFromTargets([]wordMatchTarget{
		{
			trapID:   "target-b",
			traqUUID: "target-b-uuid",
			word:     plainRegisteredWord{original: "beta", normalized: "beta"},
		},
		{
			trapID:   "target-a",
			traqUUID: "target-a-uuid",
			word:     plainRegisteredWord{original: "alpha", normalized: "alpha"},
		},
		{
			trapID:   "target-b",
			traqUUID: "target-b-uuid",
			word:     plainRegisteredWord{original: "beta-2", normalized: "beta-2"},
		},
	})
	want := []model.MatchedWords{
		{
			ContactedWords: "beta\nbeta-2",
			TrapID:         "target-b",
			TraqUUID:       "target-b-uuid",
		},
		{
			ContactedWords: "alpha",
			TrapID:         "target-a",
			TraqUUID:       "target-a-uuid",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("matched words = %#v, want %#v", got, want)
	}
}

func TestNewWordMatcher(t *testing.T) {
	t.Parallel()

	t.Run("notification target user lookup", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWords []model.WordsItem
			knownUsers      []model.UsersItem
			wantTargetWords []string
		}{
			"word is kept when its notification target user exists": {
				registeredWords: []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				knownUsers:      []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				wantTargetWords: []string{"keyword"},
			},
			"word is skipped when its notification target user is unknown": {
				registeredWords: []model.WordsItem{
					{Word: "missing", TrapId: "missing-user", IncludeBot: true, IncludeMe: true},
					{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: true},
				},
				knownUsers:      []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				wantTargetWords: []string{"keyword"},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher, err := newWordMatcher(tt.registeredWords, tt.knownUsers)
				if err != nil {
					t.Fatalf("newWordMatcher returned error: %v", err)
				}
				t.Cleanup(func() {
					if err := matcher.close(); err != nil {
						t.Errorf("wordMatcher.close returned error: %v", err)
					}
				})

				gotWords := make([]string, 0, len(matcher.targets))
				for _, target := range matcher.targets {
					gotWords = append(gotWords, target.word.text())
				}
				if !reflect.DeepEqual(gotWords, tt.wantTargetWords) {
					t.Fatalf("target words = %#v, want %#v", gotWords, tt.wantTargetWords)
				}
			})
		}
	})

	t.Run("invalid registered word handling", func(t *testing.T) {
		t.Parallel()

		matcher, err := newWordMatcher(
			[]model.WordsItem{
				{Word: "/[/", TrapId: "target", IncludeBot: true, IncludeMe: true},
				{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: true},
			},
			[]model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
		)
		if err != nil {
			t.Fatalf("newWordMatcher returned error: %v", err)
		}
		t.Cleanup(func() {
			if err := matcher.close(); err != nil {
				t.Errorf("wordMatcher.close returned error: %v", err)
			}
		})

		gotWords := make([]string, 0, len(matcher.targets))
		for _, target := range matcher.targets {
			gotWords = append(gotWords, target.word.text())
		}
		wantWords := []string{"keyword"}
		if !reflect.DeepEqual(gotWords, wantWords) {
			t.Fatalf("target words = %#v, want %#v", gotWords, wantWords)
		}
	})
}

func TestWordMatcher_MessageSender(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		senderTraqUUID string
		wantSender     messageSender
	}{
		"known human sender": {
			senderTraqUUID: "human-uuid",
			wantSender:     messageSender{traqUUID: "human-uuid", isKnown: true, isBot: false},
		},
		"known bot sender": {
			senderTraqUUID: "bot-uuid",
			wantSender:     messageSender{traqUUID: "bot-uuid", isKnown: true, isBot: true},
		},
		"unknown sender": {
			senderTraqUUID: "unknown-uuid",
			wantSender:     messageSender{traqUUID: "unknown-uuid"},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			matcher, err := newWordMatcher(
				nil,
				[]model.UsersItem{
					{TrapID: "human", TraqUUID: "human-uuid", IsBot: false},
					{TrapID: "bot", TraqUUID: "bot-uuid", IsBot: true},
				},
			)
			if err != nil {
				t.Fatalf("newWordMatcher returned error: %v", err)
			}
			t.Cleanup(func() {
				if err := matcher.close(); err != nil {
					t.Errorf("wordMatcher.close returned error: %v", err)
				}
			})

			got := matcher.messageSender(tt.senderTraqUUID)
			if got != tt.wantSender {
				t.Fatalf("message sender = %#v, want %#v", got, tt.wantSender)
			}
		})
	}
}

func TestNewWordMatchTarget(t *testing.T) {
	t.Parallel()

	user := model.UsersItem{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}

	t.Run("target metadata", func(t *testing.T) {
		t.Parallel()

		target, err := newWordMatchTarget(
			model.WordsItem{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: false},
			user,
		)
		if err != nil {
			t.Fatalf("newWordMatchTarget returned error: %v", err)
		}
		if target.includeBot != true || target.includeMe != false ||
			target.trapID != "target" || target.traqUUID != "target-uuid" ||
			target.word.text() != "keyword" {
			t.Fatalf("target = %#v, want metadata copied from word and user", target)
		}
	})

	t.Run("invalid regex returns error", func(t *testing.T) {
		t.Parallel()

		target, err := newWordMatchTarget(model.WordsItem{Word: "/[/", TrapId: "target"}, user)
		if err == nil {
			t.Fatal("newWordMatchTarget error = nil, want error")
		}
		if target.word != nil {
			t.Fatalf("target word = %#v, want nil", target.word)
		}
	})
}

func TestNewRegisteredWord(t *testing.T) {
	t.Parallel()

	t.Run("plain word", func(t *testing.T) {
		t.Parallel()

		word, err := newRegisteredWord("HELLO")
		if err != nil {
			t.Fatalf("newRegisteredWord returned error: %v", err)
		}

		plainWord, ok := word.(plainRegisteredWord)
		if !ok {
			t.Fatalf("registered word type = %T, want plainRegisteredWord", word)
		}
		if plainWord.original != "HELLO" || plainWord.normalized != "hello" {
			t.Fatalf("plain word = %#v, want original and normalized text", plainWord)
		}
	})

	t.Run("regex word", func(t *testing.T) {
		t.Parallel()

		word, err := newRegisteredWord("/hello/")
		if err != nil {
			t.Fatalf("newRegisteredWord returned error: %v", err)
		}

		regexWord, ok := word.(regexRegisteredWord)
		if !ok {
			t.Fatalf("registered word type = %T, want regexRegisteredWord", word)
		}
		t.Cleanup(func() {
			if err := regexWord.close(); err != nil {
				t.Errorf("regexRegisteredWord.close returned error: %v", err)
			}
		})
		if regexWord.original != "/hello/" {
			t.Fatalf("regex word = %#v, want original text", regexWord)
		}
	})

	t.Run("invalid regex", func(t *testing.T) {
		t.Parallel()

		word, err := newRegisteredWord("/[/")
		if err == nil {
			t.Fatal("newRegisteredWord error = nil, want error")
		}
		if word != nil {
			t.Fatalf("registered word = %#v, want nil", word)
		}
	})
}

func TestWordMatchTarget_MatchesContent(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		notifyTarget   wordMatchTarget
		messageContent messageContent
		wantMatches    bool
	}{
		"target without a registered word never matches": {
			notifyTarget:   wordMatchTarget{},
			messageContent: messageContent{raw: "keyword", normalized: "keyword"},
			wantMatches:    false,
		},
		"target matches when its registered word matches normalized content": {
			notifyTarget:   wordMatchTarget{word: plainRegisteredWord{normalized: "keyword"}},
			messageContent: messageContent{raw: "missing", normalized: "keyword"},
			wantMatches:    true,
		},
		"target does not match when its registered word misses normalized content": {
			notifyTarget:   wordMatchTarget{word: plainRegisteredWord{normalized: "keyword"}},
			messageContent: messageContent{raw: "keyword", normalized: "missing"},
			wantMatches:    false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.notifyTarget.matchesContent(tt.messageContent)
			if got != tt.wantMatches {
				t.Fatalf("matchesContent(%#v) = %v, want %v", tt.messageContent, got, tt.wantMatches)
			}
		})
	}
}

func TestPlainRegisteredWord_Matches(t *testing.T) {
	t.Parallel()

	t.Run("substring matching", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
			messageContent string
			wantMatches    bool
		}{
			"plain word uses substring matching": {
				registeredWord: "hello",
				messageContent: "hello world",
				wantMatches:    true,
			},
			"plain word absent from content does not match": {
				registeredWord: "hello",
				messageContent: "world",
				wantMatches:    false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := plainRegisteredWord{
					original:   tt.registeredWord,
					normalized: wordpattern.NormalizePlainWord(tt.registeredWord),
				}
				got := matcher.matches(newMessageContent(tt.messageContent))
				if got != tt.wantMatches {
					t.Fatalf("matches(%q, %q) = %v, want %v", tt.registeredWord, tt.messageContent, got, tt.wantMatches)
				}
			})
		}
	})

	t.Run("literal matching", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
			messageContent string
			wantMatches    bool
		}{
			"treats metacharacters literally": {
				registeredWord: "%",
				messageContent: "100%",
				wantMatches:    true,
			},
			"does not treat underscore as wildcard": {
				registeredWord: "_",
				messageContent: "hello",
				wantMatches:    false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := plainRegisteredWord{
					original:   tt.registeredWord,
					normalized: wordpattern.NormalizePlainWord(tt.registeredWord),
				}
				got := matcher.matches(newMessageContent(tt.messageContent))
				if got != tt.wantMatches {
					t.Fatalf("matches(%q, %q) = %v, want %v", tt.registeredWord, tt.messageContent, got, tt.wantMatches)
				}
			})
		}
	})

	t.Run("emoji code point matching", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
			messageContent string
			wantMatches    bool
		}{
			"base emoji matches variation-selector sequence": {
				registeredWord: "☹",
				messageContent: "☹️",
				wantMatches:    true,
			},
			"base emoji matches skin-tone modifier sequence": {
				registeredWord: "👍",
				messageContent: "👍🏽",
				wantMatches:    true,
			},
			"different emoji code points stay distinct": {
				registeredWord: "😀",
				messageContent: "😃",
				wantMatches:    false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := plainRegisteredWord{
					original:   tt.registeredWord,
					normalized: wordpattern.NormalizePlainWord(tt.registeredWord),
				}
				got := matcher.matches(newMessageContent(tt.messageContent))
				if got != tt.wantMatches {
					t.Fatalf("matches(%q, %q) = %v, want %v", tt.registeredWord, tt.messageContent, got, tt.wantMatches)
				}
			})
		}
	})
}

func TestRegexRegisteredWord_Matches(t *testing.T) {
	t.Parallel()

	t.Run("raw message content matching", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
			messageContent string
			wantMatches    bool
		}{
			"regex word matches raw message content": {
				registeredWord: "/hello/",
				messageContent: "hello",
				wantMatches:    true,
			},
			"regex word is case-sensitive unless the pattern opts in": {
				registeredWord: "/hello/",
				messageContent: "HELLO",
				wantMatches:    false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				regex, err := wordpattern.CompileRegexWord(tt.registeredWord)
				if err != nil {
					t.Fatalf("CompileRegexWord returned error: %v", err)
				}
				matcher := regexRegisteredWord{original: tt.registeredWord, regex: regex}
				t.Cleanup(func() {
					if err := matcher.close(); err != nil {
						t.Errorf("regexRegisteredWord.close returned error: %v", err)
					}
				})

				got := matcher.matches(messageContent{raw: tt.messageContent, normalized: "missing"})
				if got != tt.wantMatches {
					t.Fatalf("matches(%q, %q) = %v, want %v", tt.registeredWord, tt.messageContent, got, tt.wantMatches)
				}
			})
		}
	})

	t.Run("runtime match errors", func(t *testing.T) {
		t.Parallel()

		regex, err := wordpattern.CompileRegexWord("/.+/")
		if err != nil {
			t.Fatalf("CompileRegexWord returned error: %v", err)
		}
		matcher := regexRegisteredWord{original: "/.+/", regex: regex}
		t.Cleanup(func() {
			if err := matcher.close(); err != nil {
				t.Errorf("regexRegisteredWord.close returned error: %v", err)
			}
		})

		got := matcher.matches(messageContent{raw: string([]byte{0xff}), normalized: "missing"})
		if got {
			t.Fatal("matches returned true on runtime match error")
		}
	})
}

type fakeMessageWordMatcher struct {
	callCount        int
	matchedWordsList []model.MatchedWords
}

func (m *fakeMessageWordMatcher) matchMessage(model.MessageItem) []model.MatchedWords {
	m.callCount++
	return m.matchedWordsList
}
