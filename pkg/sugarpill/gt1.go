package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type GT1 struct {
	SetID                          string       `json:"SetID"`                          // GT1-1
	GuarantorNumber                *CX          `json:"GuarantorNumber"`                // GT1-2
	GuarantorName                  *XPN         `json:"GuarantorName"`                  // GT1-3
	GuarantorSpouseName            *XPN         `json:"GuarantorSpouseName"`            // GT1-4
	GuarantorAddress               *XAD         `json:"GuarantorAddress"`               // GT1-5
	GuarantorPhoneNumber           *XTN         `json:"GuarantorPhoneNumber"`           // GT1-6
	GuarantorBusinessPhoneNumber   *XTN         `json:"GuarantorBusinessPhoneNumber"`   // GT1-7
	GuarantorDateOfBirth           string       `json:"GuarantorDateOfBirth"`           // GT1-8
	GuarantorAdministrativeSex     string       `json:"GuarantorAdministrativeSex"`     // GT1-9
	GuarantorType                  string       `json:"GuarantorType"`                  // GT1-10
	GuarantorRelationship          *ServiceCode `json:"GuarantorRelationship"`          // GT1-11
	GuarantorSSN                   string       `json:"GuarantorSSN"`                   // GT1-12
	GuarantorDateBegin             string       `json:"GuarantorDateBegin"`             // GT1-13
	GuarantorDateEnd               string       `json:"GuarantorDateEnd"`               // GT1-14
	GuarantorPriority              string       `json:"GuarantorPriority"`              // GT1-15
	GuarantorEmployerName          *XPN         `json:"GuarantorEmployerName"`          // GT1-16
	GuarantorEmployerAddress       *XAD         `json:"GuarantorEmployerAddress"`       // GT1-17
	GuarantorEmployerPhoneNumber   *XTN         `json:"GuarantorEmployerPhoneNumber"`   // GT1-18
	GuarantorEmployeeIDNumber      *CX          `json:"GuarantorEmployeeIDNumber"`      // GT1-19
	GuarantorEmploymentStatus      string       `json:"GuarantorEmploymentStatus"`      // GT1-20
	GuarantorOrganizationName      *XON         `json:"GuarantorOrganizationName"`      // GT1-21
	GuarantorCreditRatingCode      *ServiceCode `json:"GuarantorCreditRatingCode"`      // GT1-23
	GuarantorDeathDateAndTime      string       `json:"GuarantorDeathDateAndTime"`      // GT1-24
	GuarantorDeathFlag             string       `json:"GuarantorDeathFlag"`             // GT1-25
	GuarantorChargeAdjustmentCode  *ServiceCode `json:"GuarantorChargeAdjustmentCode"`  // GT1-26
	GuarantorHouseholdAnnualIncome string       `json:"GuarantorHouseholdAnnualIncome"` // GT1-27
	GuarantorHouseholdSize         string       `json:"GuarantorHouseholdSize"`         // GT1-28
	GuarantorEmployerIDNumber      *CX          `json:"GuarantorEmployerIDNumber"`      // GT1-29
	GuarantorMaritalStatusCode     *ServiceCode `json:"GuarantorMaritalStatusCode"`     // GT1-30
	ContactPersonsName             *XPN         `json:"ContactPersonsName"`             // GT1-45
	ContactPersonsTelephoneNumber  *XTN         `json:"ContactPersonsTelephoneNumber"`  // GT1-46
	ContactReason                  *ServiceCode `json:"ContactReason"`                  // GT1-47
	ContactRelationship            *ServiceCode `json:"ContactRelationship"`            // GT1-48
	JobTitle                       string       `json:"JobTitle"`                       // GT1-49
	JobCodeClass                   *JCC         `json:"JobCodeClass"`                   // GT1-50
	GuarantorRace                  *ServiceCode `json:"GuarantorRace"`                  // GT1-55
	GuarantorBirthPlace            string       `json:"GuarantorBirthPlace"`            // GT1-56
	VIPIndicator                   string       `json:"VIPIndicator"`                   // GT1-57
}

func NewGT1Segment(p *random.Patient) *GT1 {

	gt1 := &GT1{
		GuarantorNumber:               &CX{},
		GuarantorName:                 &XPN{},
		GuarantorSpouseName:           &XPN{},
		GuarantorAddress:              &XAD{},
		GuarantorPhoneNumber:          &XTN{},
		GuarantorBusinessPhoneNumber:  &XTN{},
		GuarantorRelationship:         &ServiceCode{},
		GuarantorEmployerName:         &XPN{},
		GuarantorEmployerAddress:      &XAD{},
		GuarantorEmployerPhoneNumber:  &XTN{},
		GuarantorEmployeeIDNumber:     &CX{},
		GuarantorOrganizationName:     &XON{},
		GuarantorCreditRatingCode:     &ServiceCode{},
		GuarantorChargeAdjustmentCode: &ServiceCode{},
		GuarantorEmployerIDNumber:     &CX{},
		GuarantorMaritalStatusCode:    &ServiceCode{},
		ContactPersonsName:            &XPN{},
		ContactPersonsTelephoneNumber: &XTN{},
		ContactReason:                 &ServiceCode{},
		ContactRelationship:           &ServiceCode{},
		JobCodeClass:                  &JCC{},
		GuarantorRace:                 &ServiceCode{},
	}

	return gt1
}
