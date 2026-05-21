package sugarpill

import "placebo/pkg/random"

type IN1 struct {
	SetID                       string       `json:"SetID"`                       // IN1-1
	InsurancePlanID             *ServiceCode `json:"InsurancePlanID"`             // IN1-2
	InsuranceCompanyID          *CX          `json:"InsuranceCompanyID"`          // IN1-3
	InsuranceCompanyName        *XON         `json:"InsuranceCompanyName"`        // IN1-4
	InsuranceCompanyAddress     *XAD         `json:"InsuranceCompanyAddress"`     // IN1-5
	InsuranceCompanyContact     *XPN         `json:"InsuranceCompanyContact"`     // IN1-6
	InsuranceCompanyPhone       *XTN         `json:"InsuranceCompanyPhone"`       // IN1-7
	GroupNumber                 string       `json:"GroupNumber"`                 // IN1-8
	GroupName                   *XON         `json:"GroupName"`                   // IN1-9
	InsuredGroupEmpID           *CX          `json:"InsuredGroupEmpID"`           // IN1-10
	InsuredGroupEmpName         *XON         `json:"InsuredGroupEmpName"`         // IN1-11
	PlanEffectiveDate           string       `json:"PlanEffectiveDate"`           // IN1-12
	PlanExpirationDate          string       `json:"PlanExpirationDate"`          // IN1-13
	PlanType                    string       `json:"PlanType"`                    // IN1-15
	NameOfInsured               *XPN         `json:"NameOfInsured"`               // IN1-16
	InsuredsRelationshipToPatient *ServiceCode `json:"InsuredsRelationshipToPatient"` // IN1-17
	InsuredsDateOfBirth         string       `json:"InsuredsDateOfBirth"`         // IN1-18
	InsuredsAddress             *XAD         `json:"InsuredsAddress"`             // IN1-19
	AssignmentOfBenefits        string       `json:"AssignmentOfBenefits"`        // IN1-20
	CoordinationOfBenefits      string       `json:"CoordinationOfBenefits"`      // IN1-21
	PolicyNumber                string       `json:"PolicyNumber"`                // IN1-36
	PolicyDeductible            string       `json:"PolicyDeductible"`            // IN1-37
	InsuredsEmploymentStatus    *ServiceCode `json:"InsuredsEmploymentStatus"`    // IN1-42
	InsuredsAdministrativeSex   string       `json:"InsuredsAdministrativeSex"`   // IN1-43
	InsuredsEmployersAddress    *XAD         `json:"InsuredsEmployersAddress"`    // IN1-44
	CoverageType                string       `json:"CoverageType"`                // IN1-47
	InsuredsIDNumber            *CX          `json:"InsuredsIDNumber"`            // IN1-49
}

func NewIN1Segment(p *random.Patient) *IN1 {

	in1 := &IN1{
		InsurancePlanID:               &ServiceCode{},
		InsuranceCompanyID:            &CX{},
		InsuranceCompanyName:          &XON{},
		InsuranceCompanyAddress:       &XAD{},
		InsuranceCompanyContact:       &XPN{},
		InsuranceCompanyPhone:         &XTN{},
		GroupName:                     &XON{},
		InsuredGroupEmpID:             &CX{},
		InsuredGroupEmpName:           &XON{},
		NameOfInsured:                 &XPN{},
		InsuredsRelationshipToPatient: &ServiceCode{},
		InsuredsAddress:               &XAD{},
		InsuredsEmploymentStatus:      &ServiceCode{},
		InsuredsEmployersAddress:      &XAD{},
		InsuredsIDNumber:              &CX{},
	}

	return in1
}
