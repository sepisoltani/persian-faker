package bank

import (
	"regexp"
	"testing"
)

// TestRandomBankShebaNumber ensures that the Sheba number generated is in the correct format.
func TestRandomBankShebaNumber(t *testing.T) {
	sheba := RandomBankShebaNumber()
	expectedPattern := `^IR[0-9]{24}$`

	re := regexp.MustCompile(expectedPattern)
	if !re.MatchString(sheba) {
		t.Errorf("Generated Sheba number is invalid. Got: %s, Expected pattern: %s", sheba, expectedPattern)
	}

	if len(sheba) != 26 {
		t.Errorf("Generated Sheba number length is incorrect. Expected 26, got: %d", len(sheba))
	}
}

// TestRandomBankShebaNumber ensures that the bank card number generated is in the correct format.
func TestRandomBankCardNumber(t *testing.T) {
	cardNumber := RandomBankCardNumber()
	expectedPattern := `^6037[0-9]{12}$`

	re := regexp.MustCompile(expectedPattern)
	if !re.MatchString(cardNumber) {
		t.Errorf("Generated bank card number is invalid. Got: %s, Expected pattern: %s", cardNumber, expectedPattern)
	}

	if len(cardNumber) != 16 {
		t.Errorf("Generated bank card number length is incorrect. Expected 16, got: %d", len(cardNumber))
	}
}
