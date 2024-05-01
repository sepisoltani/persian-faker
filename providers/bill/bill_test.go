package bill

import (
	"testing"
)

// TestBillType checks if the returned bill type is in the billTypes map.
func TestBillType(t *testing.T) {
	bill := &Bill{}
	result := bill.BillType()
	found := false

	for _, v := range billTypes {
		if v == result {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("BillType returned an invalid bill type: %s", result)
	}
}
