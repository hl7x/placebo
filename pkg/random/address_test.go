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

	var tests = []struct {
		description string
		expected    []string
	}{
		{"Function Should Return Street From List", streetAddress},
	}
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Street()
			for _, street := range tc.expected {
				if strings.Contains(got.Street, street) {
					return
				}
			}
			t.Fatalf("Street()=%v, expected %v", got, tc.expected)

		})
	}
}

func TestAddress_Number(t *testing.T) {
	testAddress := Address{}
	var tests = []struct {
		description string
		maxExpected int
	}{
		{"Default should return a number between 0 and 100000", 100000},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testAddress.Number()
			if got.StructureNumber > tc.maxExpected {
				t.Fatalf("Number()=%v, expected number less than %v", got.StructureNumber, tc.maxExpected)
			}
		})
	}
}

func TestCity(t *testing.T) {
	Cities := []string{"MOBILE", "MONTGOMERY", "BIRMINGHAM"}

	var tests = []struct {
		description string
		input       string
		expected    []string
	}{
		{"State Abbreviation should return city", "AL", Cities},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := City(tc.input)
			for _, city := range Cities {
				if got == city {
					return
				}
			}
			t.Fatalf("City(%v)=%v, wanted one of %v, got %v", tc.input, got, tc.expected, got)
		})
	}
}

func TestState(t *testing.T) {
	testState := []string{
		"AL",
		"AK",
		"AZ",
		"AR",
		"CA",
		"CO",
		"CT",
		"DE",
		"DC",
		"FL",
		"GA",
		"HI",
		"ID",
		"IL",
		"IN",
		"IA",
		"KS",
		"KY",
		"LA",
		"ME",
		"MD",
		"MA",
		"MI",
		"MN",
		"MS",
		"MO",
		"MT",
		"NE",
		"NV",
		"NH",
		"NJ",
		"NM",
		"NY",
		"NC",
		"ND",
		"OH",
		"OK",
		"OR",
		"PA",
		"RI",
		"SC",
		"SD",
		"TN",
		"TX",
		"UT",
		"VT",
		"VA",
		"WA",
		"WV",
		"WI",
		"WY"}

	var tests = []struct {
		description string
		expected    []string
	}{
		{"Should Return Random State Abbreviation", testState},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := State()
			for _, state := range tc.expected {
				if got == state {
					return
				}
			}
			t.Fatalf("State()=%v, got %v expected %v", got, got, tc.expected)
		})
	}
}
