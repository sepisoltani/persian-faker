package bill

import (
	"testing"
)

// TestRandomBillType checks if the returned bill type is in the billTypes map.
func TestRandomBillType(t *testing.T) {
	result := RandomBillType()
	found := false

	for _, v := range billTypes {
		if v == result {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomBillType returned an invalid bill type: %s", result)
	}
}
