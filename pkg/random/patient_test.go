package random

import (
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
		} else if (time.Time(got.DOB)).IsZero() {
			t.Fatalf("NewPatient() DOB should not be zero")
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
		got := PatientName()
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

	t.Run("Default Case", func(t *testing.T) {
		got := testPatient.DateOfBirth()
		year := time.Time(got.DOB).Year()
		if year < 1970 || year > 2020 {
			t.Fatalf("DateOfBirth() year %v out of expected range 1970–2020", year)
		}
	})
}

func TestPatient_Arrival(t *testing.T) {

	testPatient := &Patient{}

	expected := time.Now().AddDate(0, 0, -5)

	t.Run("Default Case", func(t *testing.T) {
		got := testPatient.Arrival()
		gotTime := time.Time(got.ArrivalDate)
		if gotTime.Year() != expected.Year() || gotTime.Month() != expected.Month() || gotTime.Day() != expected.Day() {
			t.Fatalf("ArrivalDate() got %v expected %v", gotTime.Format("2006-01-02"), expected.Format("2006-01-02"))
		}
	})
}

func TestPatient_Discharge(t *testing.T) {

	testPatient := &Patient{}

	expected := time.Now().AddDate(0, 0, -3)

	t.Run("Default Test Case", func(t *testing.T) {
		got := testPatient.Discharge()
		gotTime := time.Time(got.DischargeDate)
		if gotTime.Year() != expected.Year() || gotTime.Month() != expected.Month() || gotTime.Day() != expected.Day() {
			t.Fatalf("DischargeDate() got %v expected %v", gotTime.Format("2006-01-02"), expected.Format("2006-01-02"))
		}
	})
}

func TestPatient_EventDate(t *testing.T) {

	testPatient := &Patient{}

	t.Run("Should Return Matching Current Date", func(t *testing.T) {
		got := testPatient.SetEventDate()
		gotTime := time.Time(got.EventDate)
		now := time.Now()
		if gotTime.Year() != now.Year() || gotTime.Month() != now.Month() || gotTime.Day() != now.Day() {
			t.Fatalf("EventDate got %v expected today %v", gotTime.Format("2006-01-02"), now.Format("2006-01-02"))
		}
	})
}

func TestPatient_HL7Format(t *testing.T) {

	arrival := time.Date(1998, 12, 10, 0, 0, 0, 0, time.UTC)
	discharge := time.Date(1998, 12, 11, 0, 0, 0, 0, time.UTC)
	dob := time.Date(1978, 3, 10, 0, 0, 0, 0, time.UTC)

	testPatient := &Patient{
		ArrivalDate:   PatientDate(arrival),
		DischargeDate: PatientDate(discharge),
		DOB:           PatientDate(dob),
	}

	var tests = []struct {
		description string
		got         string
		expected    string
	}{
		{"ArrivalDate HL7 format", testPatient.ArrivalDate.HL7(), "19981210"},
		{"DischargeDate HL7 format", testPatient.DischargeDate.HL7(), "19981211"},
		{"DOB HL7 format", testPatient.DOB.HL7(), "19780310"},
		{"ArrivalDate CSV format", testPatient.ArrivalDate.CSV(), "12/10/1998"},
		{"DOB CSV format", testPatient.DOB.CSV(), "3/10/1978"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			if tc.got != tc.expected {
				t.Fatalf("got %v, expected %v", tc.got, tc.expected)
			}
		})
	}
}
