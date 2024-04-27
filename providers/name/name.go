package names

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
)

type NameData struct {
	FirstNames []string `json:"firstNames"`
	LastNames  []string `json:"lastNames"`
}

var (
	nameData  NameData
	localRand *rand.Rand
)

func init() {
	localRand = rand.New(rand.NewSource(int64(os.Getpid())))
	loadData()
}

func loadData() {
	dataPath := filepath.Join("data", "names.json")
	file, err := os.ReadFile(dataPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &nameData)
	if err != nil {
		panic(err)
	}
}

func RandomFirstName() string {
	return nameData.FirstNames[localRand.Intn(len(nameData.FirstNames))]
}

func RandomLastName() string {
	return nameData.LastNames[localRand.Intn(len(nameData.LastNames))]
}

func RandomFullName() string {
	return RandomFirstName() + " " + RandomLastName()
}
