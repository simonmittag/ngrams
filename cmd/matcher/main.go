package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Trigram represents a trigram with its frequency
type Trigram struct {
	Text      string `json:"trigram"`
	Frequency int    `json:"frequency"`
}

func main() {
	// Read and parse fellowship.txt_trigrams.json
	plainTrigrams, err := readTrigramsFromFile("./output/fellowship.txt_trigrams.json")
	if err != nil {
		fmt.Printf("Error reading plain trigrams: %v\n", err)
		os.Exit(1)
	}

	// Find the most frequent trigram without spaces (should be "the")
	var mostFrequentNonSpaceTrigram string
	var maxFrequency int
	for _, t := range plainTrigrams {
		if !strings.Contains(t.Text, " ") && t.Frequency > maxFrequency {
			mostFrequentNonSpaceTrigram = t.Text
			maxFrequency = t.Frequency
		}
	}

	fmt.Printf("Most frequent trigram without spaces: %s (frequency: %d)\n",
		mostFrequentNonSpaceTrigram, maxFrequency)

	// Read and parse fellowship_enc.txt_trigrams.json
	encryptedTrigrams, err := readTrigramsFromFile("./output/fellowship_enc.txt_trigrams.json")
	if err != nil {
		fmt.Printf("Error reading encrypted trigrams: %v\n", err)
		os.Exit(1)
	}

	// Filter out trigrams with spaces
	var filteredEncTrigrams []Trigram
	for _, t := range encryptedTrigrams {
		if !strings.Contains(t.Text, " ") {
			filteredEncTrigrams = append(filteredEncTrigrams, t)
		}
	}

	fmt.Printf("Number of encrypted trigrams without spaces: %d\n", len(filteredEncTrigrams))

	// Sort filtered trigrams by frequency (descending)
	sort.Slice(filteredEncTrigrams, func(i, j int) bool {
		return filteredEncTrigrams[i].Frequency > filteredEncTrigrams[j].Frequency
	})

	// Calculate Vigenere shifts and print keywords for top trigrams
	fmt.Println("\nVigenere keywords for encrypted trigrams (sorted by frequency):")

	// Limit the number of trigrams to display
	maxTrigramsToDisplay := 50
	if len(filteredEncTrigrams) < maxTrigramsToDisplay {
		maxTrigramsToDisplay = len(filteredEncTrigrams)
	}

	for i := 0; i < maxTrigramsToDisplay; i++ {
		encTrigram := filteredEncTrigrams[i]
		keyword := calculateVigenereKeyword(mostFrequentNonSpaceTrigram, encTrigram.Text)
		fmt.Printf("Plain: %s, Encrypted: %s, Keyword: %s, Frequency: %d\n",
			mostFrequentNonSpaceTrigram, encTrigram.Text, keyword, encTrigram.Frequency)
	}
}

// readTrigramsFromFile reads and parses trigrams from a JSON file
func readTrigramsFromFile(filename string) ([]Trigram, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var trigrams []Trigram
	err = json.Unmarshal(data, &trigrams)
	if err != nil {
		return nil, err
	}

	return trigrams, nil
}

// calculateVigenereKeyword calculates the Vigenere keyword (shifts) needed to transform
// plaintext to ciphertext (a=0, b=1, etc.)
func calculateVigenereKeyword(plaintext, ciphertext string) string {
	if len(plaintext) != len(ciphertext) {
		return "Error: lengths don't match"
	}

	var keyword strings.Builder
	for i := 0; i < len(plaintext); i++ {
		// Convert characters to 0-25 range (a=0, b=1, ..., z=25)
		var plainChar, cipherChar int

		// Handle lowercase letters
		if plaintext[i] >= 'a' && plaintext[i] <= 'z' {
			plainChar = int(plaintext[i] - 'a')
		} else if plaintext[i] >= 'A' && plaintext[i] <= 'Z' {
			// Handle uppercase letters
			plainChar = int(plaintext[i] - 'A')
		} else {
			// For non-alphabetic characters, use 0 as the shift
			plainChar = 0
		}

		// Same for cipher character
		if ciphertext[i] >= 'a' && ciphertext[i] <= 'z' {
			cipherChar = int(ciphertext[i] - 'a')
		} else if ciphertext[i] >= 'A' && ciphertext[i] <= 'Z' {
			cipherChar = int(ciphertext[i] - 'A')
		} else {
			cipherChar = 0
		}

		// Calculate the shift needed (may need to add 26 if result is negative)
		shift := (cipherChar - plainChar) % 26
		if shift < 0 {
			shift += 26
		}

		// Convert shift to letter (a=0, b=1, etc.)
		keyChar := 'a' + rune(shift)
		keyword.WriteRune(keyChar)
	}

	return keyword.String()
}
