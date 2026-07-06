package message

import (
	"errors"
	"reflect"
	"testing"

	"traQ-gazer/model"
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

	t.Run("word matching", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"plain and regex words for same user are grouped together": {
				words: []model.WordsItem{
					{Word: "hello", TrapId: "target-a", IncludeBot: false, IncludeMe: true},
					{Word: "/traQ.+gazer/", TrapId: "target-a", IncludeBot: false, IncludeMe: true},
					{Word: "missing", TrapId: "target-b", IncludeBot: false, IncludeMe: true},
				},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target-a", TraqUUID: "target-a-uuid", IsBot: false},
					{TrapID: "target-b", TraqUUID: "target-b-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "hello, traQ gazer!"},
				want: []model.MatchedWords{
					{
						ContactedWords: "hello\n/traQ.+gazer/",
						TrapID:         "target-a",
						TraqUUID:       "target-a-uuid",
					},
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("matched word grouping", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"matched words are grouped by trapID in registration order": {
				words: []model.WordsItem{
					{Word: "beta", TrapId: "target-b", IncludeBot: true, IncludeMe: true},
					{Word: "alpha", TrapId: "target-a", IncludeBot: true, IncludeMe: true},
					{Word: "beta-2", TrapId: "target-b", IncludeBot: true, IncludeMe: true},
				},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target-a", TraqUUID: "target-a-uuid", IsBot: false},
					{TrapID: "target-b", TraqUUID: "target-b-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "alpha beta beta-2"},
				want: []model.MatchedWords{
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
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("self notification setting", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"includeMe true allows own message": {
				words:   []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users:   []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "target-uuid", Content: "keyword"},
				want: []model.MatchedWords{
					{ContactedWords: "keyword", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"includeMe false rejects own message": {
				words:   []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: false}},
				users:   []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "target-uuid", Content: "keyword"},
				want:    []model.MatchedWords{},
			},
			"includeMe false allows another user's message": {
				words: []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: false}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "keyword"},
				want: []model.MatchedWords{
					{ContactedWords: "keyword", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("bot notification setting", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"includeBot true allows unknown sender": {
				words:   []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users:   []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "unknown-uuid", Content: "keyword"},
				want: []model.MatchedWords{
					{ContactedWords: "keyword", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"includeBot false allows known human sender": {
				words: []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: false, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "human", TraqUUID: "human-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "human-uuid", Content: "keyword"},
				want: []model.MatchedWords{
					{ContactedWords: "keyword", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"includeBot false rejects known bot sender": {
				words: []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: false, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "bot", TraqUUID: "bot-uuid", IsBot: true},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "bot-uuid", Content: "keyword"},
				want:    []model.MatchedWords{},
			},
			"includeBot false rejects unknown sender": {
				words:   []model.WordsItem{{Word: "keyword", TrapId: "target", IncludeBot: false, IncludeMe: true}},
				users:   []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "unknown-uuid", Content: "keyword"},
				want:    []model.MatchedWords{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("plain word literal matching", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"percent is treated as a literal plain word": {
				words: []model.WordsItem{{Word: "%", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "hello"},
				want:    []model.MatchedWords{},
			},
			"underscore is treated as a literal plain word": {
				words: []model.WordsItem{{Word: "_", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "hello"},
				want:    []model.MatchedWords{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("plain word normalization", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"identical plain words match": {
				words: []model.WordsItem{{Word: "hello", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "hello"},
				want: []model.MatchedWords{
					{ContactedWords: "hello", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"latin case differences are ignored": {
				words: []model.WordsItem{{Word: "hello", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "HELLO"},
				want: []model.MatchedWords{
					{ContactedWords: "hello", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"case variants registered by same user are reported together": {
				words: []model.WordsItem{
					{Word: "hello", TrapId: "target", IncludeBot: true, IncludeMe: true},
					{Word: "HELLO", TrapId: "target", IncludeBot: true, IncludeMe: true},
				},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "hello"},
				want: []model.MatchedWords{
					{ContactedWords: "hello\nHELLO", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"hiragana and katakana variants are folded": {
				words: []model.WordsItem{{Word: "あ", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ア"},
				want: []model.MatchedWords{
					{ContactedWords: "あ", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"voiced hiragana and katakana variants are folded": {
				words: []model.WordsItem{{Word: "が", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ガ"},
				want: []model.MatchedWords{
					{ContactedWords: "が", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"different dakuten state stays distinct": {
				words: []model.WordsItem{{Word: "は", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ば"},
				want:    []model.MatchedWords{},
			},
			"small hiragana and small katakana variants are folded": {
				words: []model.WordsItem{{Word: "ぁ", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ァ"},
				want: []model.MatchedWords{
					{ContactedWords: "ぁ", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"large and small kana stay distinct": {
				words: []model.WordsItem{{Word: "あ", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ぁ"},
				want:    []model.MatchedWords{},
			},
			"fullwidth latin stays distinct": {
				words: []model.WordsItem{{Word: "a", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "ａ"},
				want:    []model.MatchedWords{},
			},
			"accented latin stays distinct": {
				words: []model.WordsItem{{Word: "e", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "é"},
				want:    []model.MatchedWords{},
			},
			"different emoji code points stay distinct": {
				words: []model.WordsItem{{Word: "😀", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "😃"},
				want:    []model.MatchedWords{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("emoji substring matching", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"base emoji matches variation-selector sequence": {
				words: []model.WordsItem{{Word: "☹", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "☹️"},
				want: []model.MatchedWords{
					{ContactedWords: "☹", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
			"base emoji matches skin-tone modifier sequence": {
				words: []model.WordsItem{{Word: "👍", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target", TraqUUID: "target-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "👍🏽"},
				want: []model.MatchedWords{
					{ContactedWords: "👍", TrapID: "target", TraqUUID: "target-uuid"},
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})

	t.Run("regex matching", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"regex words use raw content without plain-word normalization": {
				words: []model.WordsItem{
					{Word: "/hello/", TrapId: "target-a", IncludeBot: true, IncludeMe: true},
					{Word: "/あ/", TrapId: "target-b", IncludeBot: true, IncludeMe: true},
				},
				users: []model.UsersItem{
					{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false},
					{TrapID: "target-a", TraqUUID: "target-a-uuid", IsBot: false},
					{TrapID: "target-b", TraqUUID: "target-b-uuid", IsBot: false},
				},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "HELLO ア"},
				want:    []model.MatchedWords{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})
}

func TestNewWordMatcher(t *testing.T) {
	t.Parallel()

	t.Run("target construction", func(t *testing.T) {
		tests := map[string]struct {
			words   []model.WordsItem
			users   []model.UsersItem
			message model.MessageItem
			want    []model.MatchedWords
		}{
			"invalid regex word is skipped": {
				words:   []model.WordsItem{{Word: "/[/", TrapId: "target", IncludeBot: true, IncludeMe: true}},
				users:   []model.UsersItem{{TrapID: "target", TraqUUID: "target-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "["},
				want:    []model.MatchedWords{},
			},
			"word without target user is skipped": {
				words:   []model.WordsItem{{Word: "keyword", TrapId: "missing-user", IncludeBot: true, IncludeMe: true}},
				users:   []model.UsersItem{{TrapID: "sender", TraqUUID: "sender-uuid", IsBot: false}},
				message: model.MessageItem{TraqUuid: "sender-uuid", Content: "keyword"},
				want:    []model.MatchedWords{},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				matcher := mustNewWordMatcher(t, tt.words, tt.users)
				got := matcher.matchMessage(tt.message)
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("matched words = %#v, want %#v", got, tt.want)
				}
			})
		}
	})
}

func mustNewWordMatcher(t *testing.T, words []model.WordsItem, users []model.UsersItem) *wordMatcher {
	t.Helper()

	matcher, err := newWordMatcher(words, users)
	if err != nil {
		t.Fatalf("newWordMatcher returned error: %v", err)
	}
	return matcher
}

type fakeMessageWordMatcher struct {
	callCount        int
	matchedWordsList []model.MatchedWords
}

func (m *fakeMessageWordMatcher) matchMessage(model.MessageItem) []model.MatchedWords {
	m.callCount++
	return m.matchedWordsList
}
