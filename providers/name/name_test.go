package names

import (
	"testing"
)

// MockDataLoader is a mock implementation of DataLoader for testing.
type MockDataLoader struct{}

func (mdl MockDataLoader) LoadData(path string) (NameData, error) {
	return NameData{
		FirstNames: []string{"Ali", "Reza"},
		LastNames:  []string{"Kazemi", "Mousavi"},
	}, nil
}

// TestNameGenerator tests the functionality of NameGenerator.
func TestNameGenerator(t *testing.T) {
	loader := MockDataLoader{}
	seed := int64(1) // Fixed seed for reproducibility
	ng, err := NewNameGenerator(loader, seed)
	if err != nil {
		t.Fatalf("Failed to create NameGenerator: %v", err)
	}

	firstName := ng.RandomFirstName()
	if firstName != "Reza" {
		t.Errorf("Expected 'Reza', got '%s'", firstName)
	}

	lastName := ng.RandomLastName()
	if lastName != "Mousavi" {
		t.Errorf("Expected 'Kazemi', got '%s'", lastName)
	}

	fullName := ng.RandomFullName()
	if fullName != "Reza Mousavi" {
		t.Errorf("Expected 'Reza Kazemi', got '%s'", fullName)
	}
}
