package persian_faker

import (
	"github.com/sepisoltani/persian-faker/providers/bank"
	"github.com/sepisoltani/persian-faker/providers/bill"
	"github.com/sepisoltani/persian-faker/providers/digit"
	"github.com/sepisoltani/persian-faker/providers/location"
	"github.com/sepisoltani/persian-faker/providers/phonenumber"
)

// DataGenerator is a facade for accessing all fake data generation functionalities.
type DataGenerator struct {
	Bank        *bank.Bank
	Bill        *bill.Bill
	Digit       *digit.Digit
	PhoneNumber *phonenumber.PhoneNumber
	Location    *location.Location
}

func NewDataGenerator() *DataGenerator {
	return &DataGenerator{
		Bank:        &bank.Bank{},
		Bill:        &bill.Bill{},
		Digit:       &digit.Digit{},
		PhoneNumber: &phonenumber.PhoneNumber{},
		Location:    &location.Location{},
	}
}
