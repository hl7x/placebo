package event

import (
	"testing"
	"strings"

)

func TestBuilder(t *testing.T) {

	admit := []string{"admit"}
	discharge := []string{"admit", "discharge"}

	var tests = []struct {
		description	string
		input		[]string
		expected	string
	}{
		{"Should Build HL7 Event Message Based on Scenario Input", admit, "MSH"},
		{"Should Build Multipe HL7s Based On Scenario", discharge, "MSH"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Builder(tc.input)

			for _, test := range got {
				if !strings.Contains(test, tc.expected) {
					t.Fatalf("Builder(%v)=%v expected to contain '%v'", tc.input, got, tc.expected)
				}
			}

		})
	}

}
