package wordpattern

import "testing"

func TestIsRegexWord(t *testing.T) {
	t.Parallel()

	t.Run("slash-delimited words", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
		}{
			"wrapped by slashes is treated as regex": {
				registeredWord: "/hello/",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if !IsRegexWord(tt.registeredWord) {
					t.Fatalf("IsRegexWord(%q) = false, want true", tt.registeredWord)
				}
			})
		}
	})

	t.Run("plain words", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
		}{
			"leading slash only stays plain word": {
				registeredWord: "/hello",
			},
			"double leading slash without trailing delimiter stays plain word": {
				registeredWord: "//path",
			},
			"url-like word with internal slashes stays plain word": {
				registeredWord: "https://sample.invalid",
			},
			"trailing slash only stays plain word": {
				registeredWord: "hello/",
			},
			"single slash stays plain word": {
				registeredWord: "/",
			},
			"ordinary word without slash delimiters stays plain word": {
				registeredWord: "hello",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if IsRegexWord(tt.registeredWord) {
					t.Fatalf("IsRegexWord(%q) = true, want false", tt.registeredWord)
				}
			})
		}
	})
}

func TestRegexPattern(t *testing.T) {
	t.Parallel()

	got := RegexPattern("//hello//")
	want := "/hello/"
	if got != want {
		t.Fatalf("RegexPattern returned %q, want %q", got, want)
	}
}

func TestValidateRegisteredWord(t *testing.T) {
	t.Parallel()

	t.Run("plain words", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
		}{
			"plain word with regex metacharacter is accepted without compilation": {
				registeredWord: "[",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if err := ValidateRegisteredWord(tt.registeredWord); err != nil {
					t.Fatalf("ValidateRegisteredWord returned error: %v", err)
				}
			})
		}
	})

	t.Run("accepted regex words", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
		}{
			"valid PCRE regexp word is accepted": {
				registeredWord: "/traQ.+gazer/",
			},
			"fixed-length negative lookbehind is accepted": {
				registeredWord: "/(?<!blocked)token/",
			},
			"fixed-length negative lookbehind with alternatives is accepted": {
				registeredWord: "/(?<!(aa|bb)|prefix|(?<!:)@)token/",
			},
			"lookbehind with backreference-based follow-up is accepted": {
				registeredWord: "/(?<=^|(.))(?!\\1)(.)\\2{3}(?!\\2)/",
			},
			"inline case-insensitive option is accepted": {
				registeredWord: "/(?i)alpha([^s]|s[^c]|sc[^r]|$)/",
			},
			"negative lookahead with escaped dot is accepted": {
				registeredWord: "/[Pp]ackage(?!\\.invalid)/",
			},
			"unicode alternatives and emoji are accepted": {
				registeredWord: "/(alpha|ベータ|😀)/",
			},
			"word boundary assertion is accepted": {
				registeredWord: "/\\b(term|TERM)/",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if err := ValidateRegisteredWord(tt.registeredWord); err != nil {
					t.Fatalf("ValidateRegisteredWord returned error: %v", err)
				}
			})
		}
	})

	t.Run("rejected regex words", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
		}{
			"invalid PCRE regexp word is rejected": {
				registeredWord: "/[/",
			},
			"variable-length lookbehind is rejected": {
				registeredWord: "/(?<=(a|bc))a/",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				if err := ValidateRegisteredWord(tt.registeredWord); err == nil {
					t.Fatal("ValidateRegisteredWord error = nil, want error")
				}
			})
		}
	})
}

func TestRegexWord_MatchString(t *testing.T) {
	t.Parallel()

	t.Run("PCRE matching", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			registeredWord string
			messageContent string
			wantMatches    bool
		}{
			"fixed-length negative lookbehind matches when excluded prefix is absent": {
				registeredWord: "/(?<!blocked)token/",
				messageContent: "allowed token",
				wantMatches:    true,
			},
			"fixed-length negative lookbehind rejects when excluded prefix is present": {
				registeredWord: "/(?<!blocked)token/",
				messageContent: "blockedtoken",
				wantMatches:    false,
			},
			"fixed-length negative lookbehind alternatives reject a blocked prefix": {
				registeredWord: "/(?<!(aa|bb)|prefix|(?<!:)@)token/",
				messageContent: "prefixtoken",
				wantMatches:    false,
			},
			"fixed-length negative lookbehind alternatives match an allowed prefix": {
				registeredWord: "/(?<!(aa|bb)|prefix|(?<!:)@)token/",
				messageContent: "allowed token",
				wantMatches:    true,
			},
			"inline case-insensitive option matches case variants": {
				registeredWord: "/(?i)alpha([^s]|s[^c]|sc[^r]|$)/",
				messageContent: "ALPHA!",
				wantMatches:    true,
			},
			"negative lookahead rejects an excluded suffix": {
				registeredWord: "/[Pp]ackage(?!\\.invalid)/",
				messageContent: "Package.invalid",
				wantMatches:    false,
			},
			"negative lookahead matches a non-excluded suffix": {
				registeredWord: "/[Pp]ackage(?!\\.invalid)/",
				messageContent: "Package.localhost",
				wantMatches:    true,
			},
			"unicode alternatives match emoji": {
				registeredWord: "/(alpha|ベータ|😀)/",
				messageContent: "😀",
				wantMatches:    true,
			},
			"word boundary assertion matches at a boundary": {
				registeredWord: "/\\b(term|TERM)/",
				messageContent: "a term",
				wantMatches:    true,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				regex, err := CompileRegexWord(tt.registeredWord)
				if err != nil {
					t.Fatalf("CompileRegexWord returned error: %v", err)
				}
				t.Cleanup(func() {
					if err := regex.Close(); err != nil {
						t.Errorf("RegexWord.Close returned error: %v", err)
					}
				})

				got, err := regex.MatchString(tt.messageContent)
				if err != nil {
					t.Fatalf("MatchString returned error: %v", err)
				}
				if got != tt.wantMatches {
					t.Fatalf("MatchString(%q) = %v, want %v", tt.messageContent, got, tt.wantMatches)
				}
			})
		}
	})

	t.Run("runtime match errors", func(t *testing.T) {
		t.Parallel()

		regex, err := CompileRegexWord("/.+/")
		if err != nil {
			t.Fatalf("CompileRegexWord returned error: %v", err)
		}
		t.Cleanup(func() {
			if err := regex.Close(); err != nil {
				t.Errorf("RegexWord.Close returned error: %v", err)
			}
		})

		got, err := regex.MatchString(string([]byte{0xff}))
		if err == nil {
			t.Fatal("MatchString error = nil, want error")
		}
		if got {
			t.Fatal("MatchString returned true on runtime match error")
		}
	})
}

func TestNormalizePlainWord(t *testing.T) {
	t.Parallel()

	t.Run("case and kana folding", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			plainWord      string
			wantNormalized string
		}{
			"latin case is folded": {
				plainWord:      "HELLO",
				wantNormalized: "hello",
			},
			"hiragana and iteration marks are folded to katakana": {
				plainWord:      "がゞ",
				wantNormalized: "ガヾ",
			},
			"small hiragana is folded to small katakana": {
				plainWord:      "ぁ",
				wantNormalized: "ァ",
			},
			"fullwidth latin width stays distinct while case is folded": {
				plainWord:      "Ａ",
				wantNormalized: "ａ",
			},
			"accent stays distinct while case is folded": {
				plainWord:      "É",
				wantNormalized: "é",
			},
			"emoji code point stays unchanged": {
				plainWord:      "😀",
				wantNormalized: "😀",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got := NormalizePlainWord(tt.plainWord)
				if got != tt.wantNormalized {
					t.Fatalf("NormalizePlainWord(%q) = %q, want %q", tt.plainWord, got, tt.wantNormalized)
				}
			})
		}
	})

	t.Run("distinct forms", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			leftPlainWord  string
			rightPlainWord string
		}{
			"different dakuten state stays distinct": {
				leftPlainWord:  "は",
				rightPlainWord: "ば",
			},
			"large and small kana stay distinct": {
				leftPlainWord:  "あ",
				rightPlainWord: "ぁ",
			},
			"ascii and fullwidth latin stay distinct": {
				leftPlainWord:  "a",
				rightPlainWord: "ａ",
			},
			"ascii and accented latin stay distinct": {
				leftPlainWord:  "e",
				rightPlainWord: "é",
			},
			"different emoji code points stay distinct": {
				leftPlainWord:  "😀",
				rightPlainWord: "😃",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				left := NormalizePlainWord(tt.leftPlainWord)
				right := NormalizePlainWord(tt.rightPlainWord)
				if left == right {
					t.Fatalf("NormalizePlainWord(%q) and NormalizePlainWord(%q) both returned %q", tt.leftPlainWord, tt.rightPlainWord, left)
				}
			})
		}
	})
}
