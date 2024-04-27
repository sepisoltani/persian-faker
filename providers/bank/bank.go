package main

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

func main() {
	randomShebaNumber := RandomBankShebaNumber()
	fmt.Println("Random Persian Sheba number:", randomShebaNumber)
}
