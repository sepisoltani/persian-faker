package names

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
)

// NameData represents the structure for names data.
type NameData struct {
	FirstNames []string `json:"firstNames"`
	LastNames  []string `json:"lastNames"`
}

// DataLoader is the interface for loading name data.
type DataLoader interface {
	LoadData(path string) (NameData, error)
}

// FileDataLoader loads name data from a file.
type FileDataLoader struct{}

func (fdl FileDataLoader) LoadData(path string) (NameData, error) {
	var data NameData
	file, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(file, &data)
	return data, err
}

// NameGenerator holds the loaded name data and the random generator.
type NameGenerator struct {
	Data NameData
	Rand *rand.Rand
}

// NewNameGenerator creates a new NameGenerator with data loaded using the given DataLoader and random source.
func NewNameGenerator(loader DataLoader, seed int64) (*NameGenerator, error) {
	dataPath := filepath.Join("data", "names.json")
	data, err := loader.LoadData(dataPath)
	if err != nil {
		return nil, err
	}
	src := rand.NewSource(seed)
	rng := rand.New(src)
	return &NameGenerator{Data: data, Rand: rng}, nil
}

// RandomFirstName returns a random Persian first name.
func (ng *NameGenerator) RandomFirstName() string {
	return ng.Data.FirstNames[ng.Rand.Intn(len(ng.Data.FirstNames))]
}

// RandomLastName returns a random Persian last name.
func (ng *NameGenerator) RandomLastName() string {
	return ng.Data.LastNames[ng.Rand.Intn(len(ng.Data.LastNames))]
}

// RandomFullName returns a random Persian full name.
func (ng *NameGenerator) RandomFullName() string {
	return ng.RandomFirstName() + " " + ng.RandomLastName()
}
