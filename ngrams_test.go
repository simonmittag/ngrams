package ngrams

import (
	"reflect"
	"testing"
)

func TestExtractUniqueTrigrams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string][]int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: map[string][]int{},
		},
		{
			name:     "string shorter than 3",
			input:    "ab",
			expected: map[string][]int{},
		},
		{
			name:  "string exactly 3",
			input: "abc",
			expected: map[string][]int{
				"abc": {0},
			},
		},
		{
			name:  "string with repeated trigrams",
			input: "abcabc",
			expected: map[string][]int{
				"abc": {0, 3},
				"bca": {1},
				"cab": {2},
			},
		},
		{
			name:  "string with unicode characters",
			input: "hello世界",
			expected: map[string][]int{
				"hel": {0},
				"ell": {1},
				"llo": {2},
				"lo世": {3},
				"o世界": {4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractUniqueTrigrams(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExtractUniqueTrigrams(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

func TestExtractNgrams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		n        int
		expected []Ngram
	}{
		{
			name:     "empty string",
			input:    "",
			n:        1,
			expected: []Ngram{},
		},
		{
			name:     "n=0 should return empty slice",
			input:    "hello",
			n:        0,
			expected: []Ngram{},
		},
		{
			name:     "n > len(input) should return empty slice",
			input:    "hello",
			n:        10,
			expected: []Ngram{},
		},
		{
			name:  "unigrams (n=1)",
			input: "hello",
			n:     1,
			expected: []Ngram{
				{Text: "h", Position: 0},
				{Text: "e", Position: 1},
				{Text: "l", Position: 2},
				{Text: "l", Position: 3},
				{Text: "o", Position: 4},
			},
		},
		{
			name:  "bigrams (n=2)",
			input: "hello",
			n:     2,
			expected: []Ngram{
				{Text: "he", Position: 0},
				{Text: "el", Position: 1},
				{Text: "ll", Position: 2},
				{Text: "lo", Position: 3},
			},
		},
		{
			name:  "trigrams (n=3)",
			input: "hello",
			n:     3,
			expected: []Ngram{
				{Text: "hel", Position: 0},
				{Text: "ell", Position: 1},
				{Text: "llo", Position: 2},
			},
		},
		{
			name:  "n=len(input)",
			input: "hello",
			n:     5,
			expected: []Ngram{
				{Text: "hello", Position: 0},
			},
		},
		{
			name:  "longer text with spaces",
			input: "the quick brown fox",
			n:     4,
			expected: []Ngram{
				{Text: "the ", Position: 0},
				{Text: "he q", Position: 1},
				{Text: "e qu", Position: 2},
				{Text: " qui", Position: 3},
				{Text: "quic", Position: 4},
				{Text: "uick", Position: 5},
				{Text: "ick ", Position: 6},
				{Text: "ck b", Position: 7},
				{Text: "k br", Position: 8},
				{Text: " bro", Position: 9},
				{Text: "brow", Position: 10},
				{Text: "rown", Position: 11},
				{Text: "own ", Position: 12},
				{Text: "wn f", Position: 13},
				{Text: "n fo", Position: 14},
				{Text: " fox", Position: 15},
			},
		},
		{
			name:  "unicode characters",
			input: "こんにちは",
			n:     2,
			expected: []Ngram{
				{Text: "こん", Position: 0},
				{Text: "んに", Position: 1},
				{Text: "にち", Position: 2},
				{Text: "ちは", Position: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractNgrams(tt.input, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExtractNgrams(%q, %d) = %v, want %v", tt.input, tt.n, result, tt.expected)
			}
		})
	}
}
