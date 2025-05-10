package ngrams

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

type TrigramEntry struct {
	Trigram   string `json:"trigram"`
	Frequency int    `json:"frequency"`
}

func TestExtractUniqueTrigramsFromFellowship(t *testing.T) {
	filePath := filepath.Join("texts", "english", "fellowship_enc.txt")
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read fellowship text: %v", err)
	}

	uniqueTrigrams := ExtractUniqueTrigrams(string(content))

	trigramEntries := make([]TrigramEntry, 0, len(uniqueTrigrams))
	for trigram, frequency := range uniqueTrigrams {
		trigramEntries = append(trigramEntries, TrigramEntry{
			Trigram:   trigram,
			Frequency: frequency,
		})
	}

	sort.SliceStable(trigramEntries, func(i, j int) bool {
		return trigramEntries[i].Frequency > trigramEntries[j].Frequency
	})

	var buf bytes.Buffer
	buf.WriteString("[\n")
	for i, entry := range trigramEntries {
		if i > 0 {
			buf.WriteString(",\n")
		}

		buf.WriteString("  ")
		entryJSON, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("Failed to marshal trigram entry: %v", err)
		}
		buf.Write(entryJSON)
	}
	buf.WriteString("\n]")

	outputFile := filepath.Join("output", "fellowship_trigrams.json")

	jsonData := buf.Bytes()

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write JSON file: %v", err)
	}

	t.Logf("Successfully wrote %d unique trigrams to %s", len(uniqueTrigrams), outputFile)
}
