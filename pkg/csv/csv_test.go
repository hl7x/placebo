package csv

import (
	"errors"
	"strings"
	"testing"

	"placebo/pkg/random"
)

var csvPatients *random.Collection

func TestMain(m *testing.M) {

	testPatients := []*random.Patient{
		{
			FirstName:      "John",
			MiddleName:     "William",
			LastName:       "Test",
			MiddleInitial:  "W",
			MRN:            "000001",
			VisitId:        123456,
			Phone:          "123456789",
			DOB:            "08/08/1955",
			PatientAddress: &random.Address{RegionInfo: &random.Region{}},
			ArrivalDate:    "08/08/2024",
			DischargeDate:  "08/09/2024",
			Appointment:    "08/08/2024",
			Hl7Info:        &random.Hl7Dates{},
		},
	}

	csvPatients = &random.Collection{Patients: testPatients}
	m.Run()

}

func TestDelimiterValidation(t *testing.T) {

	var tests = []struct {
		description string
		input       string
		expected    error
	}{
		{"No Error When Proper Delimiter is Provided", ",", nil},
		{"Other Proper Delimiter Provided And No Error", "|", nil},
		{"Error When Improper Delimiter Provided", "#", errors.New("Delimiter Format Not Supported.")},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := delimiterValidation(tc.input)
			if tc.expected == nil {
				if got != nil {
					t.Fatalf("Expected no error, but got %v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Expected error %v, but got nil", tc.expected)
				}
				if got.Error() != "Delimiter Format Not Supported." {
					t.Fatalf("Got error message %q, expected %q", got.Error(), "Delimiter Format Not Supported.")
				}
			}
		})
	}
}

func TestCsvFormatter(t *testing.T) {

	testSlice := []string{"This", "is", "a", "slice"}

	var tests = []struct {
		description string
		input       []string
		delimiter   string
		expected    string
	}{
		{"Join slice with provided delimiter", testSlice, ",", "This,is,a,slice"},
		{"Join slice with another provided delimiter", testSlice, "|", "This|is|a|slice"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := CsvFormatter(tc.input, tc.delimiter)

			if got != tc.expected {
				t.Fatalf("Got %v, expected %v", got, tc.expected)
			}
		})
	}
}

func TestBuilder(t *testing.T) {

	testCollection := csvPatients

	var tests = []struct {
		description string
		input       *random.Collection
		delimiter   string
		expected    string
	}{
		{"Built out struct to csv", testCollection, ",", "John,William,Test,W,000001"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, _ := Builder(tc.input, tc.delimiter)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Got %v, expected to contain %v", got, tc.expected)
			}
		})
	}
}

func TestDataProcess(t *testing.T) {

	testCollection := csvPatients

	var tests = []struct {
		description string
		input       *random.Collection
		delimiter   string
		expected    string
	}{
		{"Process out header and body from struct", testCollection, ",", "John,William,Test,W,000001"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {

			got, _ := DataProcess(tc.input, tc.delimiter)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Got %v, expected to contain %v", got, tc.expected)
			}
		})
	}
}

func TestValueExtration(t *testing.T) {

	testPatient := csvPatients.Patients[0]

	slice := []string{"John", "William", "Test", "W", "000001"}

	var tests = []struct {
		description string
		input       *random.Patient
		expected    []string
	}{
		{"Struct input returns string of the associated values", testPatient, slice},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := ValueExtraction(tc.input)

			// iterate over the returned slice to match the expected
			for i, s := range tc.expected {
				if !strings.Contains(s, got[i]) {
					t.Fatalf("got %v, expected to match %v", got[i], s)
				}
			}
		})
	}
}

func TestFieldExtraction(t *testing.T) {

	testPatient := csvPatients.Patients[0]

	slice := []string{"FirstName", "MiddleName", "LastName", "MiddleInitial", "MRN"}

	var tests = []struct {
		description string
		input       *random.Patient
		expected    []string
	}{
		{"Struct input returns slice of the associated fields", testPatient, slice},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := FieldExtraction(tc.input)

			for i, s := range tc.expected {
				if !strings.Contains(s, got[i]) {
					t.Fatalf("got %v, expected to match %v", got[i], s)
				}
			}
		})
	}
}
