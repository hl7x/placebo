package random

import (
	"fmt"
	"strings"
	"testing"
	"time"
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
		} else if got.MRN == "" {
			t.Fatalf("NewPatient()=%v should not be ''", got.MRN)
		} else if got.VisitId == 0 {
			t.Fatalf("NewPatient()=%v should not be 0", got.VisitId)
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

/*
func TestPatient_Mrn(t *testing.T) {

	testPatient := &Patient{MRN: "test"}

	var tests = []struct {
		description string
		expected string
	}{
		{"Default Case for MRN", "test"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatient.Mrn()
			if got.MRN != tc.expected {
				t.Fatalf("MrnAndencounterID()=%v expected  %v and got %v", got.MRN, tc.expected, got.MRN)
			}
		})
	}
}
*/

func TestPatient_VisitID(t *testing.T) {

	testPatient := &Patient{}

	var tests = []struct {
		description string
		maxExpected int
	}{
		{"Default Case for Visit ID", 1000000000},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatient.VisitID()

			if (got.VisitId < 0) || (got.VisitId > tc.maxExpected) {
				t.Fatalf("MrnAndencounterID()=%v expected range max to %v and got %v, out of range", got.VisitId, tc.maxExpected, got.VisitId)
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

func TestPatient_Hl7DateOfBirthFmt(t *testing.T) {

	firstTestPatient := &Patient{DOB: "9-18-1988"}
	secondTestPatient := &Patient{DOB: "12-2-2001"}

	var tests = []struct {
		description string
		input       *Patient
		expected    string
	}{
		{"Should Return Reformatted Date For HL7 Message Format", firstTestPatient, "19880918"},
		{"Should Ensure That Reformatting Handles Single Digits Well", secondTestPatient, "20011202"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.input.Hl7DateOfBirthFmt()

			if got.Hl7DOB != tc.expected {
				t.Fatalf("Hl7DateOfBirth(%v)=%v expected %v", tc.input, got, tc.expected)
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

func TestEventDate(t *testing.T) {

	time := time.Now()

	timeFormated := fmt.Sprintf("%v/%v/%v", int(time.Month()), time.Day(), time.Year())

	var tests = []struct {
		description string
		expected    string
	}{
		{"Should Return Matching Current Date", timeFormated},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := EventDate()

			if got != tc.expected {
				t.Fatalf("EventDate()=%v expected %v", got, tc.expected)
			}
		})
	}

}

func TestPatient_HL7Info(t *testing.T) {

	testPatient := &Patient{ArrivalDate: "12/10/1998", DischargeDate: "12/11/1998", DOB: "3/10/1978"}

	hl7Info := &Hl7Dates{HL7Arrival: "19981210", HL7Discharge: "19981211", HL7DOB: "19780310"}

	var tests = []struct {
		description string
		input       *Patient
		expected    *Hl7Dates
	}{
		{"Should Return HL7 Info Based On Earlier Constructed Data", testPatient, hl7Info},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.input.HL7Info()

			if got.Hl7Info.HL7Arrival != tc.expected.HL7Arrival {
				t.Fatalf("got %v, expected %v", got.Hl7Info.HL7Arrival, tc.expected.HL7Arrival)
			}

			if got.Hl7Info.HL7Discharge != tc.expected.HL7Discharge {
				t.Fatalf("got %v, expected %v", got.Hl7Info.HL7Discharge, tc.expected.HL7Discharge)
			}

			if got.Hl7Info.HL7DOB != tc.expected.HL7DOB {
				t.Fatalf("got %v, expected %v", got.Hl7Info.HL7DOB, tc.expected.HL7DOB)
			}
		})
	}
}

func TestHL7DateConstructor(t *testing.T) {

	var tests = []struct {
		description string
		input       string
		expected    string
	}{
		{"Should Return a Date Formatted for HL7", "7/4/1776", "17760704"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := HL7DateConstructor(tc.input, "", "", "")

			if got.HL7Arrival != tc.expected {
				t.Fatalf("HL7DateConstructor(%v)=%v expected %v", tc.input, got.HL7Arrival, tc.expected)
			}
		})
	}
}
