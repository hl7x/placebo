package event

import (
	"testing"
	"strings"

	"placebo/pkg/random"
	"placebo/pkg/templates"

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

func TestTemplateFinder(t *testing.T) {

	admitEventType := "ADT^A01"
	dischargeEventType := "ADT^A03"

	var tests = []struct {
		description	string
		input		string
		expected	string
	}{
		{"Should Return Admit Event Based Template", "admit", admitEventType},
		{"Should Return Discharge Event Based Template", "discharge", dischargeEventType},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := TemplateFinder(tc.input)

			gotString := string(got)

			if !strings.Contains(gotString, tc.expected) {
				t.Fatalf("TemplateFinder(%v)=%v expected %v", tc.input, gotString, tc.expected)
			}
		})
	}
}

func TestTemplateMapper(t *testing.T) {

	testPatient := random.NewPatient()

	testPatient.FirstName = "John"

	testTemplate := templates.SimpleHl7Info()

	var tests = []struct {
		description	string
		patient		*random.Patient
		template	[]byte
		expected	string
	}{
		{"Should Return Patient Info Mapped to HL7 template", testPatient, testTemplate, "John"},
		{"Shoule Return the Right Template", testPatient, testTemplate, "ADT^A01"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := TemplateMapper(tc.patient, tc.template)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("TemplateMapper(%v, %v)=%v expected to include: %v", tc.patient, tc.template, got, tc.expected)
			}
		})
	}
}
