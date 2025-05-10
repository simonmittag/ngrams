package ngrams

import (
	"reflect"
	"testing"
)

func TestExtractUniqueTrigrams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "string shorter than 3",
			input:    "ab",
			expected: map[string]int{},
		},
		{
			name:  "string exactly 3",
			input: "abc",
			expected: map[string]int{
				"abc": 1,
			},
		},
		{
			name:  "string with repeated trigrams",
			input: "abcabc",
			expected: map[string]int{
				"abc": 2,
				"bca": 1,
				"cab": 1,
			},
		},
		{
			name:  "string with unicode characters",
			input: "hello世界",
			expected: map[string]int{
				"hel": 1,
				"ell": 1,
				"llo": 1,
				"lo世": 1,
				"o世界": 1,
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
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			n:        1,
			expected: []string{},
		},
		{
			name:     "n=0 should return empty slice",
			input:    "hello",
			n:        0,
			expected: []string{},
		},
		{
			name:     "n > len(input) should return empty slice",
			input:    "hello",
			n:        10,
			expected: []string{},
		},
		{
			name:     "unigrams (n=1)",
			input:    "hello",
			n:        1,
			expected: []string{"h", "e", "l", "l", "o"},
		},
		{
			name:     "bigrams (n=2)",
			input:    "hello",
			n:        2,
			expected: []string{"he", "el", "ll", "lo"},
		},
		{
			name:     "trigrams (n=3)",
			input:    "hello",
			n:        3,
			expected: []string{"hel", "ell", "llo"},
		},
		{
			name:     "n=len(input)",
			input:    "hello",
			n:        5,
			expected: []string{"hello"},
		},
		{
			name:     "longer text with spaces",
			input:    "the quick brown fox",
			n:        4,
			expected: []string{"the ", "he q", "e qu", " qui", "quic", "uick", "ick ", "ck b", "k br", " bro", "brow", "rown", "own ", "wn f", "n fo", " fox"},
		},
		{
			name:     "unicode characters",
			input:    "こんにちは",
			n:        2,
			expected: []string{"こん", "んに", "にち", "ちは"},
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
