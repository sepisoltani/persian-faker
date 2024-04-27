package location

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Sample data for the different parts of the address
var (
	buildingNumbers = func() []string {
		numbers := make([]string, 3)
		for i := range numbers {
			numbers[i] = strconv.Itoa(rand.Intn(99) + 1)
		}
		return numbers
	}()
	streetNames    = []string{"خیابان بهار", "خیابان انقلاب", "بلوار کشاورز", "خیابان پاسداران", "میدان آزادی"}
	cityPrefix     = "شهر"
	postcodePrefix = "کد پستی"
	unitPrefix     = "واحد"
	postcodes      = func() []string {
		codes := make([]string, 100)
		for i := range codes {
			codes[i] = strconv.Itoa(rand.Intn(899999) + 100000) // Six digit codes
		}
		return codes
	}()
)

// randomElement selects a random element from a slice of strings
func randomElement(items []string) string {
	return items[rand.Intn(len(items))]
}

// RandomPersianAddress generates a random Persian address
func RandomPersianAddress() string {
	// Construct the address from the parts
	return fmt.Sprintf("%s %s %s %s %s %s %s",
		cityPrefix,
		randomElement(cities),
		randomElement(streetNames),
		postcodePrefix,
		randomElement(postcodes),
		unitPrefix,
		randomElement(buildingNumbers))
}
