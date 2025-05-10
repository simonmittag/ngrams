package ngrams

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractUniqueTrigramsFromFellowship(t *testing.T) {
	filePath := filepath.Join("texts", "english", "fellowshipofthering.txt")
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read fellowship text: %v", err)
	}

	uniqueTrigrams := ExtractUniqueTrigrams(string(content))

	outputFile := filepath.Join("output", "fellowship_trigrams.json")
	jsonData, err := json.MarshalIndent(uniqueTrigrams, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal trigrams to JSON: %v", err)
	}

	err = os.WriteFile(outputFile, jsonData, 0755)
	if err != nil {
		t.Fatalf("Failed to write JSON file: %v", err)
	}

	t.Logf("Successfully wrote %d unique trigrams to %s", len(uniqueTrigrams), outputFile)
}
