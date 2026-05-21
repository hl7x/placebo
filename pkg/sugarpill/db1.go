package sugarpill

import "placebo/pkg/random"

type DB1 struct {
	SetID            string `json:"SetID"`            // DB1-1
	PersonCode       string `json:"PersonCode"`       // DB1-2
	PersonIdentifier *CX    `json:"PersonIdentifier"` // DB1-3
	Indicator        string `json:"Indicator"`        // DB1-4
	StartDate        string `json:"StartDate"`        // DB1-5
	EndDate          string `json:"EndDate"`          // DB1-6
	ReturnToWork     string `json:"ReturnToWork"`     // DB1-7
	UnableToWork     string `json:"UnableToWork"`     // DB1-8
}

func NewDB1Segment(p *random.Patient) *DB1 {

	db1 := &DB1{
		PersonIdentifier: &CX{},
	}

	return db1
}
