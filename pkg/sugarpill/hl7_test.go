package sugarpill

import (
	"strings"
	"testing"
	"time"

	"github.com/hl7x/placebo/pkg/random"
)

var sugarpillPatients random.Collection

// Return a collection of random generated patients for testing.
func SugarpillTestPatients() random.Collection {

	patients := random.NewPatients(12)

	return patients

}

func TestMain(m *testing.M) {
	sugarpillPatients = SugarpillTestPatients()
	m.Run()

}

func TestNewPV1Segment(t *testing.T) {

	sugarpillPatients.Patients[0].VisitId = 123456789
	sugarpillPatients.Patients[1].VisitId = 0
	sugarpillPatients.Patients[2].VisitId = -2

	patient1 := sugarpillPatients.Patients[0]
	patient2 := sugarpillPatients.Patients[1]
	patient3 := sugarpillPatients.Patients[2]

	var tests = []struct {
		description string
		input       *random.Patient
		expected    interface{}
	}{
		{"PV1 Visit Number Should Be the Same From Patient", patient1, 123456789},
		{"PV1 Vist Number 0 Should Reflect Visit ID 0", patient2, 0},
		{"PV1 Visit Number Unconventional should match Visit Id", patient3, -2},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewPV1Segment(tc.input)

			if got.VisitNumber != tc.expected {
				t.Fatalf("got %v, expected %v", got.VisitNumber, tc.expected)
			}
		})
	}
}

func TestNewPIDSegment(t *testing.T) {

	patient1 := sugarpillPatients.Patients[0]
	patient2 := sugarpillPatients.Patients[1]
	patient3 := sugarpillPatients.Patients[2]

	patient1.MRN = "MRN000001"
	patient2.MRN = "EEEEEEE"
	patient3.MRN = ""

	var tests = []struct {
		description string
		input       *random.Patient
		expected    interface{}
	}{
		{"PID MRN String Should Reflect the Same From the Patient", patient1, "MRN000001"},
		{"PID MRN STring Should Reflect the Only Letter String From Patient", patient2, "EEEEEEE"},
		{"PID MRN String Should Reflect empty if Patient MRN is empty", patient3, ""},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewPIDSegment(tc.input)

			if got.PatientIdentifierList != tc.expected {
				t.Fatalf("got %v, expected %v", got.PatientIdentifierList, tc.expected)
			}
		})
	}
}

func TestNewEVNSegment(t *testing.T) {

	patient1 := sugarpillPatients.Patients[0]
	patient2 := sugarpillPatients.Patients[1]

	patient1.EventDate = random.PatientDate(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
	patient2.EventDate = random.PatientDate(time.Date(1776, 7, 4, 0, 0, 0, 0, time.UTC))

	var tests = []struct {
		description string
		input       *random.Patient
		expected    string
	}{
		{"EVN Timestamp Should Reflect Patient Event Time in HL7 format", patient1, "20010101"},
		{"EVN Alternate Date Should Be Formatted as HL7", patient2, "17760704"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewEVNSegment(tc.input)

			if got.RecordedDateTime != tc.expected {
				t.Fatalf("got %v, expected %v", got.RecordedDateTime, tc.expected)
			}
		})
	}
}

func TestNewMSHSegment(t *testing.T) {

	patient1 := sugarpillPatients.Patients[0]
	//patient2 := sugarpillPatients.Patients[1]
	//patient3 := sugarpillPatients.Patients[2]

	patient1.EventDate = random.PatientDate(time.Date(1774, 7, 4, 0, 0, 0, 0, time.UTC))

	var tests = []struct {
		description string
		input       *random.Patient
		expected    any
	}{
		{"MSH Should Reflect The Patient Timestamp of the Message Event", patient1, "17740704"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewMSHSegment(tc.input)

			if got.DateTimeOfMessage != tc.expected {
				t.Fatalf("got %v, expected %v", got.DateTimeOfMessage, tc.expected)
			}
			if got.MessageControlID == "123456" {
				t.Fatalf("got %v, expected %v", got.MessageControlID, "random message ID")
			}
		})
	}
}

func TestNewHL7EventMessage(t *testing.T) {

	patient := sugarpillPatients.Patients[0]

	eventType := "ADT"

	var tests = []struct {
		description string
		patient     *random.Patient
		command     string
		expected    string
	}{
		{"Patient Admit Command", patient, "A01", "A01"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewHL7EventMessage(patient, eventType, tc.command)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("got %v, expected to contain %v", got, tc.expected)
			}
		})
	}
}

/*

	other segment functions

*/

/*
func TestNewHL7Message(t *testing.T) {

	//patient := sugarpillPatient

	var tests = []struct{
		description	string
		input		*random.Patient
		expected	interface{}
	}{
		{"HL7 Message Reflect Patient Data", nil, nil },
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := NewHL7Message(tc.input)

			if got != tc.expected {
				t.Fatalf("got %v, expected %v", got, tc.expected)
			}
		})
	}

}
*/
