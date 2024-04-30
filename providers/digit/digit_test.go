package digit

import (
	"testing"
)

// TestGeneratePersianDigit checks if the digit returned by GeneratePersianDigit is one of the valid Persian digits.
func TestGeneratePersianDigit(t *testing.T) {
	d := &Digit{}
	digit := d.GeneratePersianDigit()
	found := false
	for _, d := range faNums {
		if d == digit {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GeneratePersianDigit returned an invalid digit: %s", digit)
	}
}
