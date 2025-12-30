package main

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Counts
	}{
		{
			name:  "empty input",
			input: "",
			want:  Counts{Lines: 0, Words: 0, Bytes: 0},
		},
		{
			name:  "single word",
			input: "hello",
			want:  Counts{Lines: 0, Words: 1, Bytes: 5},
		},
		{
			name:  "single word with newline",
			input: "hello\n",
			want:  Counts{Lines: 1, Words: 1, Bytes: 6},
		},
		{
			name:  "two words",
			input: "hello world",
			want:  Counts{Lines: 0, Words: 2, Bytes: 11},
		},
		{
			name:  "multiple lines",
			input: "hello\nworld\n",
			want:  Counts{Lines: 2, Words: 2, Bytes: 12},
		},
		{
			name:  "multiple spaces between words",
			input: "hello   world",
			want:  Counts{Lines: 0, Words: 2, Bytes: 13},
		},
		{
			name:  "tabs and spaces",
			input: "hello\t\tworld\n",
			want:  Counts{Lines: 1, Words: 2, Bytes: 13},
		},
		{
			name:  "only whitespace",
			input: "   \n\t\n",
			want:  Counts{Lines: 2, Words: 0, Bytes: 6},
		},
		{
			name:  "real text",
			input: "The quick brown fox\njumps over the lazy dog\n",
			want:  Counts{Lines: 2, Words: 9, Bytes: 44},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(strings.NewReader(tt.input))

			if got.Lines != tt.want.Lines {
				t.Errorf("Lines = %d, want %d", got.Lines, tt.want.Lines)
			}
			if got.Words != tt.want.Words {
				t.Errorf("Words = %d, want %d", got.Words, tt.want.Words)
			}
			if got.Bytes != tt.want.Bytes {
				t.Errorf("Bytes = %d, want %d", got.Bytes, tt.want.Bytes)
			}
		})
	}
}
