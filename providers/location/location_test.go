package location

import (
	"testing"
)

// TestProvince checks if the randomly selected province is from the list.
func TestProvince(t *testing.T) {
	location := &Location{}
	randomProvince := location.Province()
	found := false
	for _, province := range provinces {
		if province == randomProvince {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Province returned a value not in the provinces list: %s", randomProvince)
	}
}

// TestCity checks if the randomly selected city is from the list.
func TestCity(t *testing.T) {
	location := &Location{}
	randomCity := location.City()
	found := false
	for _, city := range cities {
		if city == randomCity {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("City returned a value not in the cities list: %s", randomCity)
	}
}

// TestCountry checks if the randomly selected country is from the list.
func TestCountry(t *testing.T) {
	location := &Location{}
	randomCountry := location.Country()
	found := false
	for _, country := range countries {
		if country == randomCountry {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Country returned a value not in the countires list: %s", randomCountry)
	}
}
