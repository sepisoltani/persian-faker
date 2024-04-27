package phonenumber

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomPersianMobileNumber generates a Persian mobile phonenumber.
func RandomPersianMobileNumber() string {
	prefixes := []string{"0912", "0913", "0914", "0915", "0916"}

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	prefix := prefixes[rng.Intn(len(prefixes))]

	randomNumber := rng.Intn(10000000)

	formattedNumber := fmt.Sprintf("%07d", randomNumber)

	return prefix + formattedNumber
}
