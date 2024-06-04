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
	
	sugarpillPatient.EncounterId = 123456789

	var tests = []struct{
		description	string
		input		*random.Patient
		expected	int
	}{
		{"PV1 Visit Number Should Be the Same From Patient", sugarpillPatient, 123456789},
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

