package name

import (
	"strings"
	"testing"
)

func TestGenerateFirstName(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		n := &Name{}
		name := n.FirstName()
		if seen[name] {
			continue
		}
		seen[name] = true
		if !contains(firstNames, name) {
			t.Errorf("FirstName returned an unexpected name: %s", name)
		}
	}
	if len(seen) < 10 {
		t.Errorf("FirstName returned too few unique names, got: %d", len(seen))
	}
}

// contains checks if a string is in a slice of strings
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func TestLastName(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		n := &Name{}
		lastName := n.LastName()
		if seen[lastName] {
			continue
		}
		seen[lastName] = true
		if !contains(lastNames, lastName) {
			t.Errorf("LastName returned an unexpected name: %s", lastName)
		}
	}
	if len(seen) < 10 {
		t.Errorf("LastName returned too few unique names, got: %d", len(seen))
	}
}

func TestFullName(t *testing.T) {
	n := &Name{}
	fullName := n.FullName()
	parts := strings.Split(fullName, " ")
	if len(parts) != 2 {
		t.Errorf("FullName did not return a proper full name, got: %s", fullName)
	}
	if !contains(firstNames, parts[0]) || !contains(lastNames, parts[1]) {
		t.Errorf("FullName returned an unexpected full name: %s", fullName)
	}
}
