package sugarpill

import (
	"testing"

	"placebo/pkg/random"
)

var sugarpillPatient *random.Patient

func SugarpillPatient() *random.Patient {

	patient := random.NewPatient()

	return patient

}

func TestMain(m *testing.M) {
	sugarpillPatient = SugarpillPatient()
	m.Run()

}

func TestNewPV1Segment(t *testing.T) {
	
	var tests = []struct{
		description	string
		input		*random.Patient
		expected	interface{}
	}{
		{"PV1 Visit Number Should Be the Same From Patient", sugarpillPatient, sugarpillPatient.EncounterId},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T ) {
			got := NewPV1Segment(tc.input)

			if got.VisitNumber != tc.expected {
				t.Fatalf("got %v, expected %v", got.VisitNumber, tc.expected)
			}
		})
	}
}

