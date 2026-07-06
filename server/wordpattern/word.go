package wordpattern

import (
	"fmt"
	"strings"

	"go.arsenm.dev/pcre"
)

const regexCompileOptions = pcre.UTF

type RegexWord struct {
	regex *pcre.Regexp
}

func IsRegexWord(word string) bool {
	return len(word) >= 2 && strings.HasPrefix(word, "/") && strings.HasSuffix(word, "/")
}

func RegexPattern(word string) string {
	if !IsRegexWord(word) {
		return word
	}
	return word[1 : len(word)-1]
}

// CompileRegexWord defines the supported regex dialect for registered words.
// Matching and registration validation both use PCRE, matching MariaDB REGEXP more closely.
func CompileRegexWord(word string) (*RegexWord, error) {
	regex, err := pcre.CompileOpts(RegexPattern(word), regexCompileOptions)
	if err != nil {
		return nil, err
	}
	return &RegexWord{regex: regex}, nil
}

func (r *RegexWord) MatchString(s string) (matched bool, err error) {
	if r == nil || r.regex == nil {
		return false, nil
	}
	defer func() {
		if recovered := recover(); recovered != nil {
			matched = false
			err = fmt.Errorf("regex match failed: %v", recovered)
		}
	}()
	return r.regex.MatchString(s), nil
}

func (r *RegexWord) Close() error {
	if r == nil || r.regex == nil {
		return nil
	}
	err := r.regex.Close()
	r.regex = nil
	return err
}

func ValidateRegisteredWord(word string) error {
	if !IsRegexWord(word) {
		return nil
	}
	regex, err := CompileRegexWord(word)
	if err != nil {
		return fmt.Errorf("invalid regex word: %w", err)
	}
	return regex.Close()
}

// NormalizePlainWord defines app-side plain-word matching semantics.
// It folds case and hiragana/katakana only; width, accents, marks, and emoji stay distinct.
func NormalizePlainWord(word string) string {
	return strings.Map(foldHiraganaToKatakana, strings.ToLower(word))
}

func foldHiraganaToKatakana(r rune) rune {
	if r >= 'ぁ' && r <= 'ゖ' {
		return r + ('ァ' - 'ぁ')
	}
	switch r {
	case 'ゝ':
		return 'ヽ'
	case 'ゞ':
		return 'ヾ'
	default:
		return r
	}
}
