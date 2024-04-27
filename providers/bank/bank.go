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

// RandomPersianBankName returns a random Persian bank name from a predefined list.
func RandomPersianBankName() string {
	bankNames := []string{
		"بانک ملی ایران",
		"بانک سپه",
		"بانک تجارت",
		"بانک ملت",
		"بانک صادرات ایران",
		"بانک کشاورزی",
		"بانک مسکن",
		"بانک صنعت و معدن",
		"بانک توسعه صادرات",
		"بانک توسعه تعاون",
		"پست بانک ایران",
		"بانک اقتصاد نوین",
		"بانک پارسیان",
		"بانک پاسارگاد",
		"بانک کارآفرین",
		"بانک سامان",
		"بانک سینا",
		"بانک خاورمیانه",
		"بانک شهر",
		"بانک دی",
		"بانک صنعت و معدن",
		"بانک آینده",
		"بانک گردشگری",
		"بانک ایران زمین",
		"بانک قرض الحسنه مهر ایران",
	}

	// Create a new random source and a random generator for better concurrency support
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Select a random bank name
	return bankNames[rng.Intn(len(bankNames))]
}
