package digit

import (
	"testing"
)

// TestDigit checks if the digit returned by Digit is one of the valid Persian digits.
func TestDigit(t *testing.T) {
	d := &Digit{}
	digit := d.Digit()
	found := false
	for _, d := range faNums {
		if d == digit {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Digit returned an invalid digit: %s", digit)
	}
}
