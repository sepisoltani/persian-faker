package location

import (
	"testing"
)

// TestRandomProvince checks if the randomly selected province is from the list.
func TestRandomProvince(t *testing.T) {
	randomProvince := RandomProvince()
	found := false
	for _, province := range provinces {
		if province == randomProvince {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomProvince returned a value not in the provinces list: %s", randomProvince)
	}
}

// TestRandomCity checks if the randomly selected city is from the list.
func TestRandomCity(t *testing.T) {
	randomCity := RandomCity()
	found := false
	for _, city := range cities {
		if city == randomCity {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomCity returned a value not in the cities list: %s", randomCity)
	}
}
