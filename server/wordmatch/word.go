package wordmatch

import (
	"fmt"
	"regexp"
	"strings"
)

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
// Matching and registration validation both use Go's regexp package.
func CompileRegexWord(word string) (*regexp.Regexp, error) {
	return regexp.Compile(RegexPattern(word))
}

func ValidateRegisteredWord(word string) error {
	if !IsRegexWord(word) {
		return nil
	}
	if _, err := CompileRegexWord(word); err != nil {
		return fmt.Errorf("invalid regex word: %w", err)
	}
	return nil
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
