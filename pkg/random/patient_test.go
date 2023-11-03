package random

import (
	"fmt"
	"time"
	"strings"
	"testing"
)

func TestNewPatients(t *testing.T) {
	var tests = []struct {
		description string
		expectedMax int
	}{
		{"Single Default Case", 1},
		{"Multiple Case", 5},
		{"Empty Case", 0},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewPatients(tc.expectedMax)
			if len(got.Patients) != tc.expectedMax {
				t.Fatalf("NewPatients(%v)=%v, Expected %v", tc.expectedMax, got, tc.expectedMax)
			}
		})
	}
}

func TestNewPatient(t *testing.T) {
	t.Run("Return Fake Patient", func(t *testing.T) {
		got := NewPatient()
		if got.FirstName == "" {
			t.Fatalf("NewPatient()=%v should not be empty", got.FirstName)
		} else if got.LastName == "" {
			t.Fatalf("NewPatient()=%v should not be empty", got.LastName)
		} else if got.MRN == 0 {
			t.Fatalf("NewPatient()=%v should not be 0", got.MRN)
		} else if got.EncounterId == 0 {
			t.Fatalf("NewPatient()=%v should not be 0", got.EncounterId)
		} else if got.Phone == "" {
			t.Fatalf("NewPatient()=%v should not be empty", got.Phone)
		} else if got.DOB == "" {
			t.Fatalf("NewPatient()=%v should not be empty", got.DOB)
		}
	})
}

// Placeholder
func TestName(t *testing.T) {

	names := append(FIRSTNAME, LASTNAME...)

	var tests = []struct {
		description string
		expected    []string
	}{
		{"First Names Test", names},
		{"Last Names Test", names},
	}

	for _, tc := range tests {
		got := Name()
		t.Run(tc.description, func(t *testing.T) {
			for _, name := range tc.expected {
				if strings.Contains(got.FirstName, name) {
					return
				}
			}
			t.Fatalf("Name()=%v expected one of %v got %v", got.FirstName, tc.expected, got.FirstName)
		})
	}
}

func TestPatient_Mrn(t *testing.T) {

	testPatient := &Patient{}

	var tests = []struct {
		description string
		maxExpected int
	}{
		{"Default Case for MRN", 1000000000},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatient.Mrn()
			if (got.MRN < 0) || (got.MRN > tc.maxExpected) {
				t.Fatalf("MrnAndencounterID()=%v expected range max to %v and got %v, out of range", got.MRN, tc.maxExpected, got.MRN)
			}
		})
	}
}

func TestPatient_EncounterID(t *testing.T) {

	testPatient := &Patient{}

	var tests = []struct {
		description string
		maxExpected int
	}{
		{"Default Case for Encounter ID", 1000000000},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatient.EncounterID()
	
			if (got.EncounterId < 0) || (got.EncounterId > tc.maxExpected) {
				t.Fatalf("MrnAndencounterID()=%v expected range max to %v and got %v, out of range", got.EncounterId, tc.maxExpected, got.EncounterId)
			}
		})
	}
}

func TestPatient_DateOfBirth(t *testing.T) {

	testPatient := &Patient{}

	var tests = []struct {
		description string
		expected    string
	}{
		{"Default Case", "-"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatient.DateOfBirth()
			if !strings.Contains(got.DOB, tc.expected) {
				t.Fatalf("DateOfBirth()=%v got %v expected %v", got.DOB, got.DOB, tc.expected)
			}
		})
	}
}

func TestPatient_Arrival(t *testing.T) {

	testPatient := &Patient{}

	testTime := time.Now().AddDate(0, 0, -5)

	arrival := fmt.Sprintf("%v/%v/%v", int(testTime.Month()), testTime.Day(), testTime.Year())

	var test = struct {
		description string
		expected    string
	}{
		"Default Case", arrival,
	}

	t.Run(test.description, func(t *testing.T) {
		got := testPatient.Arrival()
		if got.ArrivalDate != test.expected {
			t.Fatalf("ArrivalDate()=%v got %v and expected %v", got.ArrivalDate, got.ArrivalDate, test.expected)
		}
	})

}

func TestPatient_Discharge(t *testing.T) {

	testPatient := Patient{}

        testTime := time.Now().AddDate(0, 0, -3)

        discharge := fmt.Sprintf("%v/%v/%v", int(testTime.Month()), testTime.Day(), testTime.Year())

	var test = struct {
		description string
		expected    string
	}{
		"Default Test Case", discharge,
	}

	t.Run(test.description, func(t *testing.T) {
		got := testPatient.Discharge()
		if got.DischargeDate != test.expected {
			t.Fatalf("DischargeDate()=%v got %v and expected %v", got.DischargeDate, got.DischargeDate, test.expected)
		}
	})
}
