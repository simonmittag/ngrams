package ngrams

import "fmt"

// Ngram represents an n-gram with its position in the original text
type Ngram struct {
	Text     string
	Position int
}

func ExtractUniqueTrigrams(input string) map[string][]int {
	tg := ExtractTrigrams(input)
	m := make(map[string][]int)
	for _, t := range tg {
		m[t.Text] = append(m[t.Text], t.Position)
	}
	fmt.Printf("number of unique trigrams: %d\n", len(m))
	return m
}

func ExtractTrigrams(input string) []Ngram {
	fmt.Printf("number of runes: %d\n", len([]rune(input)))
	tg := ExtractNgrams(input, 3)
	fmt.Printf("number of trigrams: %d\n", len(tg))
	return tg
}

func ExtractNgrams(input string, n int) []Ngram {
	runes := []rune(input)

	// this is the edge cases when it doesn't work
	if n < 1 || n > len(runes) {
		return []Ngram{}
	}

	result := make([]Ngram, 0, len(runes)-n+1)

	// Extract each n-gram by sliding a window of size n over the runes
	for i := 0; i <= len(runes)-n; i++ {
		ngram := string(runes[i : i+n])
		result = append(result, Ngram{Text: ngram, Position: i})
	}

	return result
}
