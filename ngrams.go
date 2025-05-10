package ngrams

func ExtractUniqueTrigrams(input string) map[string]int {
	tg := ExtractTrigrams(input)
	m := make(map[string]int)
	for _, t := range tg {
		m[t]++
	}
	return m
}

func ExtractTrigrams(input string) []string {
	return ExtractNgrams(input, 3)
}

func ExtractNgrams(input string, n int) []string {
	runes := []rune(input)

	// this is the edge cases when it doesn't work
	if n < 1 || n > len(runes) {
		return []string{}
	}

	result := make([]string, 0, len(runes)-n+1)

	// Extract each n-gram by sliding a window of size n over the runes
	for i := 0; i <= len(runes)-n; i++ {
		ngram := string(runes[i : i+n])
		result = append(result, ngram)
	}

	return result
}
