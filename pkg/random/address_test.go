package random

import (
	"strings"
	"testing"
)

func TestPatient_NewAddress(t *testing.T) {

	t.Run("Return Address Properties", func(t *testing.T) {
		testPatient := &Patient{}
		address := testPatient.NewAddress()
		got := address
		if got.PatientAddress.Street == "" {
			t.Fatalf("Got %v, Street should not be empty", got.PatientAddress.Street)
		} else if got.PatientAddress.StructureNumber > 10000 || got.PatientAddress.StructureNumber < 0 {
			t.Fatalf("Got %v, number should be greater than 0 and less than 10000", got.PatientAddress.StructureNumber)
		} else if got.PatientAddress.RegionInfo.City == "" {
			t.Fatalf("Got %v, City Should not be empty", got.PatientAddress.RegionInfo.City)
		}
	})
}

// TODO: Correct the test
func TestAddress_RegionSpecific(t *testing.T) {
	testAddress := Address{}

	var tests = []struct {
		description string
		expected    string
	}{
		{"Region should include State", ""},
		{"Region should include City", ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testAddress.RegionSpecific()
			if !strings.Contains(got.RegionInfo.City, tc.expected) {
				t.Fatalf("Test Failed %v", got.RegionInfo.City)
			} else if !strings.Contains(got.RegionInfo.State, tc.expected) {
				t.Fatalf("Test Failed %v", got.RegionInfo.State)
			}
		})
	}
}

func TestStreet(t *testing.T) {

	streetAddress := []string{
		"MAIN",
		"FIRST",
		"SECOND",
		"THIRD",
		"FOURTH",
		"FIFTH",
		"SIXTH",
		"SEVENTH",
		"EIGTH",
		"NINTH",
		"TENTH",
		"ELEVENTH",
		"TWELFTH",
		"MAPLE",
		"ELM",
		"BIRCH",
		"RED",
		"ORANGE",
		"YELLOW",
		"BLUE",
		"PURPLE"}

	var tests = []struct{
		description	string
		expected	[]string
	}{
		{"Function Should Return Street From List", streetAddress},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T){
		got := Street()
		if got != strings.Contains() {
			t.Fatalf("Street()=%v, incorrect amount, expected %v", got.Street.len(), tc.maxExpected)
		}

		})
	}
}
