package csv

import (
	"testing"
	"errors"
)

func TestdelimiterValidation(t *testing.T) {

	var tests = []struct{
		description	string
		input		string
		expected	error
	}{
		{"No Error When Proper Delimiter is Provided", ",", nil},
		{"Other Proper Delimiter Provided And No Error", "|", nil},
		{"Error When Improper Delimiter Provided", "#", errors.New("")},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := delimiterValidation(tc.input)

			if got != tc.expected {
				t.Fatalf("Got %v, expected %v", got, tc.expected)
			}
		})
	}
}
