package cmd

import (
	"testing"
	"strings"

	"placebo/pkg/random"
)

func TestSendHl7Message(t *testing.T) {

	var tests = []struct{
		description	string
		input		string
		expected	error
	}{
		{"Should Return Nil When No Command Is Given", "", nil},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := SendHl7Message(tc.input)

			if got != tc.expected {
				t.Fatalf("SendHl7Message(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestHl7Format(t *testing.T) {
	
	patientTest := &random.Patient{FirstName: "Joe", Hl7Info: &random.Hl7Dates{}}

	var tests = []struct{
		description	string
		input		*random.Patient
		expected	string
	}{
		{"Should Return Templated Data", patientTest, "MSH"},
		{"Should Return Patient Data", patientTest, "Joe"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Hl7Format(tc.input)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("Hl7Format(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}
}
