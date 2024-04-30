package name

import (
	"strings"
	"testing"
)

func TestGenerateFirstName(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		n := &Name{}
		name := n.GenerateFirstName()
		if seen[name] {
			continue
		}
		seen[name] = true
		if !contains(firstNames, name) {
			t.Errorf("GenerateFirstName returned an unexpected name: %s", name)
		}
	}
	if len(seen) < 10 { // Check for some variability
		t.Errorf("GenerateFirstName returned too few unique names, got: %d", len(seen))
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

func TestGenerateLastName(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		n := &Name{}
		name := n.GenerateLastName()
		if seen[name] {
			continue
		}
		seen[name] = true
		if !contains(lastNames, name) {
			t.Errorf("GenerateLastName returned an unexpected name: %s", name)
		}
	}
	if len(seen) < 10 { // Check for some variability
		t.Errorf("GenerateLastName returned too few unique names, got: %d", len(seen))
	}
}

func TestGenerateFullName(t *testing.T) {
	n := &Name{}
	name := n.GenerateFullName()
	parts := strings.Split(name, " ")
	if len(parts) != 2 {
		t.Errorf("GenerateFullName did not return a proper full name, got: %s", name)
	}
	if !contains(firstNames, parts[0]) || !contains(lastNames, parts[1]) {
		t.Errorf("GenerateFullName returned an unexpected full name: %s", name)
	}
}
