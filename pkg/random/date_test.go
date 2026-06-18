package random

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	t.Run("Should produce a random date in the DOB year range", func(t *testing.T) {
		got := Date()
		year := time.Time(got).Year()
		if year < 1970 || year > 2020 {
			t.Fatalf("Date() year %v out of expected range 1970–2020", year)
		}
	})
}

func TestPatientDate_HL7(t *testing.T) {
	var tests = []struct {
		description string
		input       time.Time
		expected    string
	}{
		{"Fourth of July 1776", time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC), "17760704"},
		{"Christmas 2000", time.Date(2000, 12, 25, 0, 0, 0, 0, time.UTC), "20001225"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := PatientDate(tc.input).HL7()
			if got != tc.expected {
				t.Fatalf("HL7()=%v expected %v", got, tc.expected)
			}
		})
	}
}

func TestPatientDate_CSV(t *testing.T) {
	var tests = []struct {
		description string
		input       time.Time
		expected    string
	}{
		{"Fourth of July 1776", time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC), "7/4/1776"},
		{"Christmas 2000", time.Date(2000, 12, 25, 0, 0, 0, 0, time.UTC), "12/25/2000"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := PatientDate(tc.input).CSV()
			if got != tc.expected {
				t.Fatalf("CSV()=%v expected %v", got, tc.expected)
			}
		})
	}
}

func TestMonth(t *testing.T) {

	var tests = []struct {
		description string
		expected    int
	}{
		{"Number 12 or Less", 12},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Month()
			if (got > tc.expected) || (got < 0) {
				t.Fatalf("Month()=%v expected %v", got, tc.expected)
			}
		})
	}

}

func TestDay(t *testing.T) {

	var tests = []struct {
		description string
		expected    int
	}{
		{"Should Produce a Number Less Than 28", 28},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Day()
			if got > tc.expected {
				t.Fatalf("Day()=%v expected %v", got, tc.expected)
			}
		})
	}

}

func TestYear(t *testing.T) {

	var tests = []struct {
		description string
		expected    int
	}{
		{"Should Produce a Number Greater Than 1920", 1920},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Year()
			if got < tc.expected {
				t.Fatalf("Year()=%v expected %v", got, tc.expected)
			}
		})
	}

}
