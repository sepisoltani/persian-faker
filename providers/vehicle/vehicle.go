package vehicle

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// CarPlateNumber generates a random Persian car plate number.
func CarPlateNumber() string {
	letterRunes := []rune("الفبپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی")
	numbers := make([]rune, 0, 10)
	for i := '۰'; i <= '۹'; i++ {
		numbers = append(numbers, i)
	}

	plate := make([]rune, 0, 8)
	plate = append(plate, numbers[rng.Intn(len(numbers))], numbers[rng.Intn(len(numbers))])
	plate = append(plate, ' ')
	plate = append(plate, letterRunes[rng.Intn(len(letterRunes))])
	plate = append(plate, ' ')
	plate = append(plate, numbers[rng.Intn(len(numbers))], numbers[rng.Intn(len(numbers))], numbers[rng.Intn(len(numbers))])
	plate = append(plate, ' ')
	plate = append(plate, numbers[rng.Intn(len(numbers))], numbers[rng.Intn(len(numbers))])

	return string(plate)
}
