package number

import (
	"testing"
)

// TestRandomPersianMobileNumber checks if the generated mobile number is correctly formatted.
func TestRandomPersianMobileNumber(t *testing.T) {
	generatedNumber := RandomPersianMobileNumber()

	if len(generatedNumber) != 11 {
		t.Errorf("Expected mobile number length of 11, but got %d", len(generatedNumber))
	}

	validPrefixes := []string{"0912", "0913", "0914", "0915", "0916"}
	prefix := generatedNumber[:4]

	found := false
	for _, validPrefix := range validPrefixes {
		if prefix == validPrefix {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Generated mobile number has an invalid prefix: %s", prefix)
	}

	for _, c := range generatedNumber[4:] {
		if c < '0' || c > '9' {
			t.Errorf("Generated mobile number contains non-digit characters: %s", generatedNumber)
			break
		}
	}
}
