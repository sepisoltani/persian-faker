package phonenumber

import (
	"testing"
)

// TestPhoneNumber checks if the generated mobile phone number is correctly formatted.
func TestPhoneNumber(t *testing.T) {
	phoneNumber := &PhoneNumber{}
	generatedNumber := phoneNumber.PhoneNumber()

	if len(generatedNumber) != 11 {
		t.Errorf("PhoneNumber length of 11, but got %d", len(generatedNumber))
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
		t.Errorf("PhoneNumber has an invalid prefix: %s", prefix)
	}

	for _, c := range generatedNumber[4:] {
		if c < '0' || c > '9' {
			t.Errorf("PhoneNumber contains non-digit characters: %s", generatedNumber)
			break
		}
	}
}
