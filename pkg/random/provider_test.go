package random

import (
	"slices"
	"strings"
	"testing"
)

func TestNewProvider(t *testing.T) {

	var tests = []struct {
		description string
		expected    []string
	}{
		{"Creates a randeom provider with a first name", FIRSTNAME},
		{"Creates a randeom provider with a last name", LASTNAME},
		{"Creates a randeom provider with a middle name", MIDDLENAME},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewProvider()
			middleSplit := strings.Split(got.MiddleName, "")

			if slices.Equal(tc.expected, FIRSTNAME) {
				if slices.Contains(FIRSTNAME, got.FirstName) {
					return
				}
				t.Fatalf("NewProvider()=%v, expected in %v", got.FirstName, tc.expected)
			}
			if slices.Equal(tc.expected, LASTNAME) {
				if slices.Contains(LASTNAME, got.LastName) {
					return
				}
				t.Fatalf("NewProvider()=%v, expected in %v", got.LastName, tc.expected)
			}
			if slices.Equal(tc.expected, MIDDLENAME) {
				if slices.Contains(MIDDLENAME, got.MiddleName) {
					return
				}
				t.Fatalf("NewProvider()=%v, expected in %v", got.MiddleName, tc.expected)
			}

			if got.MiddleInitial != middleSplit[0] {
				t.Fatalf("NewProvider()=%v, expected %v", got.MiddleInitial, middleSplit[0])
			}

		})
	}
}
