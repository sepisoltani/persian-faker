package location

import (
	"strconv"
	"strings"
	"testing"
)

// TestGeneratePersianAddress verifies that the format of the generated address is correct.
func TestGeneratePersianAddress(t *testing.T) {
	location := &Location{}
	address := location.GeneratePersianAddress()
	// Split the address into components based on spaces
	parts := strings.Fields(address)

	// Check if there are exactly 7 parts as expected
	if len(parts) != 9 {
		t.Fatalf("Expected 9 parts in the address, got %d: %s", len(parts), address)
	}

	// Check the static parts based on the known prefixes
	if parts[0] != cityPrefix {
		t.Errorf("Expected city prefix '%s', got '%s'", cityPrefix, parts[0])
	}

	// Check if the postcode is numerical
	if _, err := strconv.Atoi(parts[6]); err != nil {
		t.Errorf("Expected a numerical postcode, got '%s'", parts[6])
	}

	// Validate the building number
	if _, err := strconv.Atoi(parts[8]); err != nil {
		t.Errorf("Expected a numerical building number, got '%s'", parts[8])
	}
}
