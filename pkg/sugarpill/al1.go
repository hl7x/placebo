package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type AL1 struct {
	SetID                   string       `json:"SetID"`                   // AL1-1
	AllergenTypeCode        *ServiceCode `json:"AllergenTypeCode"`        // AL1-2
	AllergenCodeDescription *ServiceCode `json:"AllergenCodeDescription"` // AL1-3
	AllergySeverityCode     *ServiceCode `json:"AllergySeverityCode"`     // AL1-4
	AllergyReactionCode     string       `json:"AllergyReactionCode"`     // AL1-5
	IdentificationDate      string       `json:"IdentificationDate"`      // AL1-6
}

func NewAL1Segment(p *random.Patient) *AL1 {

	al1 := &AL1{
		AllergenTypeCode:        &ServiceCode{},
		AllergenCodeDescription: &ServiceCode{},
		AllergySeverityCode:     &ServiceCode{},
	}

	return al1
}
