package sugarpill

import "placebo/pkg/random"

type PV2 struct {
	PriorPendingLocation           *PatientLocation `json:"PriorPendingLocation"`           // PV2-1
	AccommodationCode              *ServiceCode     `json:"AccommodationCode"`              // PV2-2
	AdmitReason                    *ServiceCode     `json:"AdmitReason"`                    // PV2-3
	TransferReason                 *ServiceCode     `json:"TransferReason"`                 // PV2-4
	PatientValuables               string           `json:"PatientValuables"`               // PV2-5
	PatientValuablesLocation       string           `json:"PatientValuablesLocation"`       // PV2-6
	VisitUserCode                  string           `json:"VisitUserCode"`                  // PV2-7
	ExpectedAdmitDateTime          string           `json:"ExpectedAdmitDateTime"`          // PV2-8
	ExpectedDischargeDateTime      string           `json:"ExpectedDischargeDateTime"`      // PV2-9
	EstimatedLengthOfInpatientStay string           `json:"EstimatedLengthOfInpatientStay"` // PV2-10
	ActualLengthOfInpatientStay    string           `json:"ActualLengthOfInpatientStay"`    // PV2-11
	VisitDescription               string           `json:"VisitDescription"`               // PV2-12
	ReferralSourceCode             *XCN             `json:"ReferralSourceCode"`             // PV2-13
	PreviousServiceDate            string           `json:"PreviousServiceDate"`            // PV2-14
	EmploymentIllnessRelated       string           `json:"EmploymentIllnessRelated"`       // PV2-15
	PurgeStatusCode                string           `json:"PurgeStatusCode"`                // PV2-16
	PurgeStatusDate                string           `json:"PurgeStatusDate"`                // PV2-17
	SpecialProgramCode             string           `json:"SpecialProgramCode"`             // PV2-18
	RetentionIndicator             string           `json:"RetentionIndicator"`             // PV2-19
	ExpectedNumberOfInsurancePlans string           `json:"ExpectedNumberOfInsurancePlans"` // PV2-20
	VisitPublicityCode             string           `json:"VisitPublicityCode"`             // PV2-21
	VisitProtectionIndicator       string           `json:"VisitProtectionIndicator"`       // PV2-22
	ClinicOrganizationName         *XON             `json:"ClinicOrganizationName"`         // PV2-23
	PatientStatusCode              string           `json:"PatientStatusCode"`              // PV2-24
	VisitPriorityCode              string           `json:"VisitPriorityCode"`              // PV2-25
	PreviousTreatmentDate          string           `json:"PreviousTreatmentDate"`          // PV2-26
	ExpectedDischargeDisposition   string           `json:"ExpectedDischargeDisposition"`   // PV2-27
	SignatureOnFileDate            string           `json:"SignatureOnFileDate"`            // PV2-28
	FirstSimilarIllnessDate        string           `json:"FirstSimilarIllnessDate"`        // PV2-29
	PatientChargeAdjustmentCode    *ServiceCode     `json:"PatientChargeAdjustmentCode"`    // PV2-30
	RecurringServiceCode           string           `json:"RecurringServiceCode"`           // PV2-31
	BillingMediaCode               string           `json:"BillingMediaCode"`               // PV2-32
	ExpectationOfBillTypeCode      string           `json:"ExpectationOfBillTypeCode"`      // PV2-33
	MilitaryPartnershipCode        string           `json:"VisitProtectionIndicator"`       // PV2-34
}

func NewPV2Segment(p *random.Patient) *PV2 {

	pv2 := &PV2{
		PriorPendingLocation:        &PatientLocation{},
		AccommodationCode:           &ServiceCode{},
		AdmitReason:                 &ServiceCode{},
		TransferReason:              &ServiceCode{},
		ReferralSourceCode:          &XCN{},
		ClinicOrganizationName:      &XON{},
		PatientChargeAdjustmentCode: &ServiceCode{},
	}

	return pv2
}
