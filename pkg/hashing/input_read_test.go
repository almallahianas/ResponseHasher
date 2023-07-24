package hashing

import (
	"flag"
	"log"
	"testing"
)

type InputReaderMock struct {
	MaxConcurrent int
	Urls          []string
}

func (m *InputReaderMock) Read() {
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	// set test input args
	flag.IntVar(&m.MaxConcurrent, "parallel", 10, "Your number of parallel processes")
	err := flag.CommandLine.Parse([]string{"-parallel=5", "example1", "example2"})
	if err != nil {
		log.Printf("test error parsing the input flags, with error:%v", err)
	}
	m.Urls = flag.Args()
}

func TestReadCliInput(t *testing.T) {
	mockInput := &InputReaderMock{}
	mockInput.Read()

	// Test the values set by Read
	if mockInput.MaxConcurrent != 5 {
		t.Errorf("MaxConcurrent got: %d, expected: %d", mockInput.MaxConcurrent, 5)
	}

	expectedUrls := []string{"example1", "example2"}
	if !compareStringSlices(mockInput.Urls, expectedUrls) {
		t.Errorf("urls got: %v, expected: %v", mockInput.Urls, expectedUrls)
	}
}

func compareStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
