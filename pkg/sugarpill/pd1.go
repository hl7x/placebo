package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type PD1 struct {
	LivingDependency                        string       `json:"LivingDependency"`                        // PD1-1
	LivingArrangement                       string       `json:"LivingArrangement"`                       // PD1-2
	PatientPrimaryFacility                  *XON         `json:"PatientPrimaryFacility"`                  // PD1-3
	StudentIndicator                        string       `json:"StudentIndicator"`                        // PD1-5
	Handicap                                string       `json:"Handicap"`                                // PD1-6
	LivingWillCode                          string       `json:"LivingWillCode"`                          // PD1-7
	OrganDonorCode                          string       `json:"OrganDonorCode"`                          // PD1-8
	SeparateBillFlag                        string       `json:"SeparateBillFlag"`                        // PD1-9
	DuplicatePatient                        *CX          `json:"DuplicatePatient"`                        // PD1-10
	PublicityCode                           *ServiceCode `json:"PublicityCode"`                           // PD1-11
	ProtectionIndicator                     string       `json:"ProtectionIndicator"`                     // PD1-12
	ProtectionIndicatorEffectiveDate        string       `json:"ProtectionIndicatorEffectiveDate"`        // PD1-13
	PlaceOfWorship                          *XON         `json:"PlaceOfWorship"`                          // PD1-14
	AdvanceDirectiveCode                    *ServiceCode `json:"AdvanceDirectiveCode"`                    // PD1-15
	ImmunizationRegistryStatus              string       `json:"ImmunizationRegistryStatus"`              // PD1-16
	ImmunizationRegistryStatusEffectiveDate string       `json:"ImmunizationRegistryStatusEffectiveDate"` // PD1-17
	PublicityCodeEffectiveDate              string       `json:"PublicityCodeEffectiveDate"`              // PD1-18
	MilitaryBranch                          string       `json:"MilitaryBranch"`                          // PD1-19
	MilitaryRankGrade                       string       `json:"MilitaryRankGrade"`                       // PD1-20
	MilitaryStatus                          string       `json:"MilitaryStatus"`                          // PD1-21
}

func NewPD1Segment(p *random.Patient) *PD1 {

	pd1 := &PD1{
		PatientPrimaryFacility: &XON{},
		DuplicatePatient:       &CX{},
		PublicityCode:          &ServiceCode{},
		PlaceOfWorship:         &XON{},
		AdvanceDirectiveCode:   &ServiceCode{},
	}

	return pd1
}
