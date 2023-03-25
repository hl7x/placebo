package random

import (
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
