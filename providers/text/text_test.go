package text

import (
	"testing"
)

func TestSentence(t *testing.T) {
	text := &Text{}
	sentence := text.Sentence()

	// Check if the returned sentence is non-empty
	if sentence == "" {
		t.Error("Expected a non-empty string, got an empty string")
	}

	// Check if the returned sentence is in the predefined list
	found := false
	for _, s := range exampleSentences {
		if s == sentence {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("The sentence '%s' is not in the predefined list of sentences", sentence)
	}
}
