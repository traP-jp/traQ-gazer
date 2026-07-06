package wordmatch

import "testing"

func TestIsRegexWord(t *testing.T) {
	t.Parallel()

	t.Run("slash delimiters", func(t *testing.T) {
		tests := map[string]struct {
			word string
			want bool
		}{
			"wrapped by slashes is treated as regex": {
				word: "/hello/",
				want: true,
			},
			"leading slash only stays plain word": {
				word: "/hello",
				want: false,
			},
			"trailing slash only stays plain word": {
				word: "hello/",
				want: false,
			},
			"single slash stays plain word": {
				word: "/",
				want: false,
			},
			"ordinary word without slash delimiters stays plain word": {
				word: "hello",
				want: false,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got := IsRegexWord(tt.word)
				if got != tt.want {
					t.Fatalf("IsRegexWord(%q) = %v, want %v", tt.word, got, tt.want)
				}
			})
		}
	})
}

func TestRegexPattern(t *testing.T) {
	t.Parallel()

	t.Run("slash delimiters", func(t *testing.T) {
		t.Parallel()

		got := RegexPattern("//hello//")
		want := "/hello/"
		if got != want {
			t.Fatalf("RegexPattern returned %q, want %q", got, want)
		}
	})
}

func TestValidateRegisteredWord(t *testing.T) {
	t.Parallel()

	t.Run("regex syntax validation", func(t *testing.T) {
		tests := map[string]struct {
			word    string
			wantErr bool
		}{
			"plain word with regex metacharacter is not compiled": {
				word:    "[",
				wantErr: false,
			},
			"valid Go regexp word is accepted": {
				word:    "/traQ.+gazer/",
				wantErr: false,
			},
			"invalid Go regexp word is rejected": {
				word:    "/[/",
				wantErr: true,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				err := ValidateRegisteredWord(tt.word)
				if (err != nil) != tt.wantErr {
					t.Fatalf("ValidateRegisteredWord error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
}

func TestNormalizePlainWord(t *testing.T) {
	t.Parallel()

	t.Run("case and kana folding", func(t *testing.T) {
		tests := map[string]struct {
			word string
			want string
		}{
			"latin case is folded": {
				word: "HELLO",
				want: "hello",
			},
			"hiragana and iteration marks are folded to katakana": {
				word: "がゞ",
				want: "ガヾ",
			},
			"fullwidth latin width stays distinct while case is folded": {
				word: "Ａ",
				want: "ａ",
			},
			"accent stays distinct while case is folded": {
				word: "É",
				want: "é",
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				got := NormalizePlainWord(tt.word)
				if got != tt.want {
					t.Fatalf("NormalizePlainWord(%q) = %q, want %q", tt.word, got, tt.want)
				}
			})
		}
	})
}
