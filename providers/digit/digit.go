package digit

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// faNums contains Persian digits.
var faNums = []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"}

type Digit struct {
}

// Digit returns a random Persian digit.
func (Digit) Digit() string {
	return faNums[rng.Intn(len(faNums))]
}
