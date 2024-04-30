package persian_faker

import "github.com/sepisoltani/persian-faker/providers/bank"

// DataGenerator is a facade for accessing all fake data generation functionalities.
type DataGenerator struct {
	Bank *bank.Bank
}

func NewDataGenerator() *DataGenerator {
	return &DataGenerator{
		Bank: &bank.Bank{},
	}
}
