package persian_faker

import (
	"github.com/sepisoltani/persian-faker/providers/bank"
	"github.com/sepisoltani/persian-faker/providers/bill"
)

// DataGenerator is a facade for accessing all fake data generation functionalities.
type DataGenerator struct {
	Bank *bank.Bank
	Bill *bill.Bill
}

func NewDataGenerator() *DataGenerator {
	return &DataGenerator{
		Bank: &bank.Bank{},
		Bill: &bill.Bill{},
	}
}
