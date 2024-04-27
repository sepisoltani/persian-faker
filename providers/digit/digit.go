package digit

import (
	"math/rand"
	"time"
)

// faNums contains Persian digits.
var faNums = []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"}

// PersianDigit returns a random Persian digit.
func PersianDigit() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	return faNums[rng.Intn(len(faNums))]
}
