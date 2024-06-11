package sugarpill

import (
	"testing"

	"placebo/pkg/random"
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

	sugarpillPatients.Patients[0].EncounterId = 123456789
	sugarpillPatients.Patients[1].EncounterId = 0
	sugarpillPatients.Patients[2].EncounterId = -2

	patient1 := sugarpillPatients.Patients[0]
	patient2 := sugarpillPatients.Patients[1]
	patient3 := sugarpillPatients.Patients[2]

	var tests = []struct {
		description string
		input       *random.Patient
		expected    interface{}
	}{
		{"PV1 Visit Number Should Be the Same From Patient", patient1, 123456789},
		{"PV1 Vist Number 0 Should Reflect Encounter ID 0", patient2, 0},
		{"PV1 Visit Number Unconventional should match Encounter Id", patient3, -2},
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
		description	string
		input 		*random.Patient
		expected	interface{}
	}{
		{"PID MRN String Should Reflect the Same From the Patient", patient1, "MRN000001"},
		{"PID MRN STring Should Reflect the Only Letter String From Patient", patient2, "EEEEEEE"},
		{"PID MRN String Should Reflect empty if Patient MRN is empty", patient3 , ""},
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
