package location

import (
	"testing"
)

// TestGenerateProvince checks if the randomly selected province is from the list.
func TestGenerateProvince(t *testing.T) {
	location := &Location{}
	randomProvince := location.GenerateProvince()
	found := false
	for _, province := range provinces {
		if province == randomProvince {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GenerateProvince returned a value not in the provinces list: %s", randomProvince)
	}
}

// TestGenerateCity checks if the randomly selected city is from the list.
func TestGenerateCity(t *testing.T) {
	location := &Location{}
	randomCity := location.GenerateCity()
	found := false
	for _, city := range cities {
		if city == randomCity {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GenerateCity returned a value not in the cities list: %s", randomCity)
	}
}

// TestGenerateCountry checks if the randomly selected city is from the list.
func TestGenerateCountry(t *testing.T) {
	location := &Location{}
	randomCountry := location.GenerateCountry()
	found := false
	for _, country := range countries {
		if country == randomCountry {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GenerateCountry returned a value not in the cities list: %s", randomCountry)
	}
}
