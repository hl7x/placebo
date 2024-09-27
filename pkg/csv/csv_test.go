package csv

import (
	"testing"
	"errors"
	"strings"

	"placebo/pkg/random"
)

var csvPatients *random.Collection

func TestMain(m *testing.M) {
	
	testPatients := []*random.Patient{
		{	
			FirstName: "John",
			LastName: "Test",
			MRN: "000001",
			EncounterId: 123456,
			Phone: "123456789",
			DOB: "08/08/1955",
			Hl7DOB: "08081955",
			PatientAddress: &random.Address{RegionInfo: &random.Region{}},
			ArrivalDate: "08/08/2024",
			DischargeDate: "08/09/2024",
			Appointment: "08/08/2024",
			Hl7Info: &random.Hl7Dates{},
		},
	}

	csvPatients = &random.Collection{Patients: testPatients}
	m.Run()

}

func TestdelimiterValidation(t *testing.T) {

	var tests = []struct{
		description	string
		input		string
		expected	error
	}{
		{"No Error When Proper Delimiter is Provided", ",", nil},
		{"Other Proper Delimiter Provided And No Error", "|", nil},
		{"Error When Improper Delimiter Provided", "#", errors.New("")},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := delimiterValidation(tc.input)

			if got != tc.expected {
				t.Fatalf("Got %v, expected %v", got, tc.expected)
			}
		})
	}
}

func TestCsvFormatter(t *testing.T) {

	testSlice := []string{"This", "is", "a", "slice"}

	var tests = []struct{
		description	string
		input		[]string
		delimiter	string
		expected 	string
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

	var tests =  []struct{
		description	string
		input		*random.Collection
		expected	string
	}{
		{"Built out struct to csv", testCollection, "John,Test,000001"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, _ := Builder(tc.input, ",")

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Got %v, expected to contain %v", got, tc.expected)
			}
		})
	}
}



