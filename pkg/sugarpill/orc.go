package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type ORC struct {
	OrderControl                  string            `json:"OrderControl"`                  // ORC-1
	PlacerOrderNumber             *EntityIdentifier `json:"PlacerOrderNumber"`             // ORC-2
	FillerOrderNumber             *EntityIdentifier `json:"FillerOrderNumber"`             // ORC-3
	PlacerGroupNumber             *EntityIdentifier `json:"PlacerGroupNumber"`             // ORC-4
	OrderStatus                   string            `json:"OrderStatus"`                   // ORC-5
	ResponseFlag                  string            `json:"ResponseFlag"`                  // ORC-6
	DateTimeOfTransaction         string            `json:"DateTimeOfTransaction"`         // ORC-9
	EnteredBy                     *XCN              `json:"EnteredBy"`                     // ORC-10
	VerifiedBy                    *XCN              `json:"VerifiedBy"`                    // ORC-11
	OrderingProvider              *XCN              `json:"OrderingProvider"`              // ORC-12
	EntererLocation               *PatientLocation  `json:"EntererLocation"`               // ORC-13
	CallbackPhoneNumber           *XTN              `json:"CallbackPhoneNumber"`           // ORC-14
	OrderEffectiveDateTime        string            `json:"OrderEffectiveDateTime"`        // ORC-15
	OrderControlCodeReason        *ServiceCode      `json:"OrderControlCodeReason"`        // ORC-16
	EnteringOrganization          *ServiceCode      `json:"EnteringOrganization"`          // ORC-17
	EnteringDevice                *ServiceCode      `json:"EnteringDevice"`                // ORC-18
	ActionBy                      *XCN              `json:"ActionBy"`                      // ORC-19
	AdvancedBeneficiaryNoticeCode *ServiceCode      `json:"AdvancedBeneficiaryNoticeCode"` // ORC-20
	OrderingFacilityName          *XON              `json:"OrderingFacilityName"`          // ORC-21
	OrderingFacilityAddress       *XAD              `json:"OrderingFacilityAddress"`       // ORC-22
	OrderingFacilityPhoneNumber   *XTN              `json:"OrderingFacilityPhoneNumber"`   // ORC-23
	OrderingProviderAddress       *XAD              `json:"OrderingProviderAddress"`       // ORC-24
	OrderStatusModifier           *ServiceCode      `json:"OrderStatusModifier"`           // ORC-25
}

func NewORCSegment(p *random.Patient) *ORC {

	orc := &ORC{
		PlacerOrderNumber:             &EntityIdentifier{},
		FillerOrderNumber:             &EntityIdentifier{},
		PlacerGroupNumber:             &EntityIdentifier{},
		EnteredBy:                     &XCN{},
		VerifiedBy:                    &XCN{},
		OrderingProvider:              &XCN{},
		EntererLocation:               &PatientLocation{},
		CallbackPhoneNumber:           &XTN{},
		OrderControlCodeReason:        &ServiceCode{},
		EnteringOrganization:          &ServiceCode{},
		EnteringDevice:                &ServiceCode{},
		ActionBy:                      &XCN{},
		AdvancedBeneficiaryNoticeCode: &ServiceCode{},
		OrderingFacilityName:          &XON{},
		OrderingFacilityAddress:       &XAD{},
		OrderingFacilityPhoneNumber:   &XTN{},
		OrderingProviderAddress:       &XAD{},
		OrderStatusModifier:           &ServiceCode{},
	}

	return orc
}
