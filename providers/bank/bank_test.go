package bank

import (
	"regexp"
	"testing"
)

// TestIBAN ensures that the IBAN generated is in the correct format.
func TestIBAN(t *testing.T) {
	bank := &Bank{}
	iban := bank.IBAN()
	expectedPattern := `^IR[0-9]{24}$`

	re := regexp.MustCompile(expectedPattern)
	if !re.MatchString(iban) {
		t.Errorf("Generated IBAN is invalid. Got: %s, Expected pattern: %s", iban, expectedPattern)
	}

	if len(iban) != 26 {
		t.Errorf("Generated IBAN length is incorrect. Expected 26, got: %d", len(iban))
	}
}

// TestCardNumber ensures that the bank card number generated is in the correct format.
func TestCardNumber(t *testing.T) {
	bank := &Bank{}
	cardNumber := bank.CardNumber()
	expectedPattern := `^6037[0-9]{12}$`

	re := regexp.MustCompile(expectedPattern)
	if !re.MatchString(cardNumber) {
		t.Errorf("Generated bank card number is invalid. Got: %s, Expected pattern: %s", cardNumber, expectedPattern)
	}

	if len(cardNumber) != 16 {
		t.Errorf("Generated bank card number length is incorrect. Expected 16, got: %d", len(cardNumber))
	}
}

// TestBankName checks if the returned bank name is from the list of defined bank names.
func TestBankName(t *testing.T) {
	bank := &Bank{}
	bankName := bank.BankName() // Use the GenerateBankName method
	validBankNames := map[string]bool{
		"بانک ملی ایران":            true,
		"بانک سپه":                  true,
		"بانک تجارت":                true,
		"بانک ملت":                  true,
		"بانک صادرات ایران":         true,
		"بانک کشاورزی":              true,
		"بانک مسکن":                 true,
		"بانک توسعه صادرات":         true,
		"بانک توسعه تعاون":          true,
		"پست بانک ایران":            true,
		"بانک اقتصاد نوین":          true,
		"بانک پارسیان":              true,
		"بانک پاسارگاد":             true,
		"بانک کارآفرین":             true,
		"بانک سامان":                true,
		"بانک سینا":                 true,
		"بانک خاورمیانه":            true,
		"بانک شهر":                  true,
		"بانک دی":                   true,
		"بانک صنعت و معدن":          true,
		"بانک آینده":                true,
		"بانک گردشگری":              true,
		"بانک ایران زمین":           true,
		"بانک قرض الحسنه مهر ایران": true,
	}

	if !validBankNames[bankName] {
		t.Errorf("Generated bank name '%s' is not in the list of valid bank names.", bankName)
	}
}
