package vehicle

import (
	"regexp"
	"testing"
)

func TestCarPlateNumber(t *testing.T) {
	plate := CarPlateNumber()
	t.Logf("Generated Car Plate: %s", plate)

	// Persian numbers and letters
	numbers := "۰۱۲۳۴۵۶۷۸۹"
	letters := "الفبپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی"

	// Regex to match the format: "DD L DDD DD"
	// D: Persian digit, L: Persian letter
	plateRegex := `^[` + numbers + `]{2} [` + letters + `] [` + numbers + `]{3} [` + numbers + `]{2}$`

	// Compile the regex
	re := regexp.MustCompile(plateRegex)

	if !re.MatchString(plate) {
		t.Errorf("Car plate does not match expected format: %s", plate)
	}
}
