package bill

import (
	"testing"
)

// TestGenerateBillType checks if the returned bill type is in the billTypes map.
func TestGenerateBillType(t *testing.T) {
	bill := &Bill{}
	result := bill.GenerateBillType()
	found := false

	for _, v := range billTypes {
		if v == result {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GenerateBillType returned an invalid bill type: %s", result)
	}
}
