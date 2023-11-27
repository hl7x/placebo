package random

import (
	"testing"
	"strings"
)

func TestDate(t *testing.T) {

	var tests = []struct{
		description	string
		expected	string
	}{
		{"Should produce a random date.", "-"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Date()
			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Date=%v, expected a date with %v", got, tc.expected)
			}
		})
	}

}

func TestMonth(t *testing.T) {
	
	var tests = []struct {
		description	string
		expected 	int
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
		description	string
		expected	int
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
		description	string
		expected	int
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

func TestHl7DateFormatter(t *testing.T) {

	var tests = []struct {
		description	string
		input		string
		expected	string
	}{
		{"Date With '-' Converted to HL7 Format", "12-25-2000","20001225"},
		{"Date with '/' Converted to HL7 Format", "7/4/1988", "19880704"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Hl7DateFormatter(tc.input)

			if got != tc.expected {
				t.Fatalf("Hl7DateFormatter(%v)=%v expected %v", tc.input,  got, tc.expected)
			}
		})
	}

}
