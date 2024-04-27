package bank

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomBankShebaNumber generates a random Persian Bank Sheba number.
func RandomBankShebaNumber() string {

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	sheba := "IR"

	for i := 0; i < 24; i++ {
		digit := rng.Intn(10) // Generate a single digit (0-9)
		sheba += fmt.Sprintf("%d", digit)
	}

	return sheba
}

// RandomBankCardNumber generates a bank card number starting with "6037" followed by 12 random digits.
func RandomBankCardNumber() string {
	// Create a new random source and a random generator for better concurrency support
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Start with the prefix "6037"
	cardNumber := "6037"

	// Append 12 random digits to the card number
	for i := 0; i < 12; i++ {
		digit := rng.Intn(10) // Generate a single digit (0-9)
		cardNumber += fmt.Sprintf("%d", digit)
	}

	return cardNumber
}
