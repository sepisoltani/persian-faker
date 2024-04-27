package digit

import (
	"testing"
)

// TestPersianDigit checks if the digit returned by PersianDigit is one of the valid Persian digits.
func TestPersianDigit(t *testing.T) {
	digit := PersianDigit()
	found := false
	for _, d := range faNums {
		if d == digit {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("PersianDigit returned an invalid digit: %s", digit)
	}
}
