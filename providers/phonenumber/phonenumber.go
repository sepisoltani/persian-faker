package phonenumber

import (
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type phoneNumber struct {
}

// GeneratePersianMobileNumber generates a Persian mobile phone number.
func (phoneNumber) GeneratePersianMobileNumber() string {
	prefixes := []string{"0912", "0913", "0914", "0915", "0916"}

	prefix := prefixes[rng.Intn(len(prefixes))]

	randomNumber := rng.Intn(10000000)

	formattedNumber := fmt.Sprintf("%07d", randomNumber)

	return prefix + formattedNumber
}
