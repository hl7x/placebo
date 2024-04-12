package random

import (
	"strings"
	"testing"
)

func TestPatient_PhoneNumber(t *testing.T) {

	testPatient := &Patient{}

	testPatientAddressAndPhone := testPatient.NewAddress()

	var tests = []struct {
		description string
		expected    string
	}{
		{"Default Case", "5551212"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := testPatientAddressAndPhone.PhoneNumber()
			if !strings.Contains(got.Phone, tc.expected) {
				t.Fatalf("Phone()=%v should containt %v", got.Phone, tc.expected)
			}
		})
	}
}

func TestAreaCode(t *testing.T) {
	var tests = []struct {
		description string
		input       string
		expected    string
	}{
		{"Default Case of City in the List", "DOVER", "302"},
		{"Case Where City isn't in the List", "ST. PAUL", ""},
		{"Case Where City is Bad Format", "Dover", ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := AreaCode(tc.input)
			if got != tc.expected {
				t.Fatalf("AreaCode(%v)=%v, expected %v and got %v", tc.input, got, tc.expected, got)
			}
		})
	}
}
