package bank

import (
	"math/rand"
	"strconv"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Bank struct{}

// GenerateBankIBAN generates a random Persian bank IBAN number.
func (Bank) GenerateBankIBAN() string {
	sheba := "IR"
	for i := 0; i < 24; i++ {
		digit := rng.Intn(10)
		sheba += strconv.Itoa(digit)
	}
	return sheba
}

// GenerateBankCardNumber generates a bank card number.
func (Bank) GenerateBankCardNumber() string {
	cardNumber := "6037"
	for i := 0; i < 12; i++ {
		digit := rng.Intn(10)
		cardNumber += strconv.Itoa(digit)
	}
	return cardNumber
}

// GenerateBankName returns a random Persian bank name from a predefined list.
func (Bank) GenerateBankName() string {
	bankNames := []string{
		"بانک ملی ایران", "بانک سپه", "بانک تجارت", "بانک ملت",
		"بانک صادرات ایران", "بانک کشاورزی", "بانک مسکن", "بانک صنعت و معدن",
		"بانک توسعه صادرات", "بانک توسعه تعاون", "پست بانک ایران", "بانک اقتصاد نوین",
		"بانک پارسیان", "بانک پاسارگاد", "بانک کارآفرین", "بانک سامان",
		"بانک سینا", "بانک خاورمیانه", "بانک شهر", "بانک دی", "بانک صنعت و معدن",
		"بانک آینده", "بانک گردشگری", "بانک ایران زمین", "بانک قرض الحسنه مهر ایران",
	}
	return bankNames[rng.Intn(len(bankNames))]
}
