package persianfaker

import (
	"reflect"
	"testing"

	"github.com/sepisoltani/persian-faker/providers/bank"
	"github.com/sepisoltani/persian-faker/providers/bill"
	"github.com/sepisoltani/persian-faker/providers/digit"
	"github.com/sepisoltani/persian-faker/providers/location"
	"github.com/sepisoltani/persian-faker/providers/name"
	"github.com/sepisoltani/persian-faker/providers/phonenumber"
)

// TestNewDataGenerator tests the New function for proper initialization of DataGenerator.
func TestNewDataGenerator(t *testing.T) {
	dg := New()

	// Test if all the struct fields are initialized and not nil
	if dg.Bank == nil {
		t.Error("Expected Bank to be initialized, got nil")
	}
	if dg.Bill == nil {
		t.Error("Expected Bill to be initialized, got nil")
	}
	if dg.Digit == nil {
		t.Error("Expected Digit to be initialized, got nil")
	}
	if dg.PhoneNumber == nil {
		t.Error("Expected PhoneNumber to be initialized, got nil")
	}
	if dg.Location == nil {
		t.Error("Expected Location to be initialized, got nil")
	}
	if dg.Name == nil {
		t.Error("Expected Name to be initialized, got nil")
	}

	// Test if the types of all fields are correctly set
	if reflect.TypeOf(dg.Bank).Elem() != reflect.TypeOf(bank.Bank{}) {
		t.Errorf("Expected Bank type to be *bank.Bank, got %T", dg.Bank)
	}
	if reflect.TypeOf(dg.Bill).Elem() != reflect.TypeOf(bill.Bill{}) {
		t.Errorf("Expected Bill type to be *bill.Bill, got %T", dg.Bill)
	}
	if reflect.TypeOf(dg.Digit).Elem() != reflect.TypeOf(digit.Digit{}) {
		t.Errorf("Expected Digit type to be *digit.Digit, got %T", dg.Digit)
	}
	if reflect.TypeOf(dg.PhoneNumber).Elem() != reflect.TypeOf(phonenumber.PhoneNumber{}) {
		t.Errorf("Expected PhoneNumber type to be *phonenumber.PhoneNumber, got %T", dg.PhoneNumber)
	}
	if reflect.TypeOf(dg.Location).Elem() != reflect.TypeOf(location.Location{}) {
		t.Errorf("Expected Location type to be *location.Location, got %T", dg.Location)
	}
	if reflect.TypeOf(dg.Name).Elem() != reflect.TypeOf(name.Name{}) {
		t.Errorf("Expected Name type to be *name.Name, got %T", dg.Name)
	}
}
