package parser_test

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"nem12/models"
	"nem12/pkg/parser"
	"os"
	"testing"
)

// Compare function to compare two slices of *models.MeterReading
func compare(a, b []*models.MeterReading) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !(a[i].NMI == b[i].NMI &&
			a[i].Consumption == b[i].Consumption &&
			a[i].Timestamp == b[i].Timestamp) {
			return false
		}
	}
	return true
}

// test with valid format
func TestParserValid(t *testing.T) {
	file, err := os.Open("../../data/data.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the expected JSON output from the file
	expectedData, err := os.ReadFile("../../data/dataOutput.json")
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}

	// Unmarshal the expected JSON data into a map
	var expectedOutput []*models.MeterReading
	if err := json.Unmarshal(expectedData, &expectedOutput); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON data: %v", err)
	}

	reader := csv.NewReader(file)
	actualOutput, err := parser.ParseRecord(reader)
	if !compare(expectedOutput, actualOutput) || err != nil {
		t.Fatalf(`TestParserValid = %v, %v, want match for %v, nil`, actualOutput, err, expectedOutput)
	}
}

// test with invalid 200 format
func TestParserInvalid200(t *testing.T) {
	file, err := os.Open("../../data/dataInvalid200.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	actualOutput, err := parser.ParseRecord(reader)
	if actualOutput != nil || err == nil {
		t.Fatalf(`TestParserInvalid200 = %v, %v, want "", error`, actualOutput, err)
	}
}

// test with invalid 300 format
func TestParserNo300(t *testing.T) {
	file, err := os.Open("../../data/dataInvalid300.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	actualOutput, err := parser.ParseRecord(reader)
	if actualOutput != nil || err == nil {
		t.Fatalf(`TestParserInvalid300 = %v, %v, want "", error`, actualOutput, err)
	}
}
