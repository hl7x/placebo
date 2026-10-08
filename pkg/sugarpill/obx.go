package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type OBX struct {
	SetID                                 string            `json:"SetID"`                                 // OBX-1
	ValueType                             string            `json:"ValueType"`                             // OBX-2
	ObservationIdentifier                 string            `json:"ObservationIdentifier"`                 // OBX-3
	ObservationSubID                      string            `json:"ObservationSubID"`                      // OBX-4
	ObservationValue                      string            `json:"ObservationValue"`                      // OBX-5
	Units                                 *ServiceCode      `json:"Units"`                                 // OBX-6
	ReferencesRange                       string            `json:"ReferencesRange"`                       // OBX-7
	AbnormalFlags                         string            `json:"AbnormalFlags"`                         // OBX-8
	Probability                           string            `json:"Probability"`                           // OBX-9
	NatureOfAbnormalTest                  string            `json:"NatureOfAbnormalTest"`                  // OBX-10
	ObservationResultStatus               string            `json:"ObservationResultStatus"`               // OBX-11
	EffectiveDateOfReferenceRange         string            `json:"EffectiveDateOfReferenceRange"`         // OBX-12
	UserDefinedAccessChecks               string            `json:"UserDefinedAccessChecks"`               // OBX-13
	DateTimeOfTheObservation              string            `json:"DateTimeOfTheObservation"`              // OBX-14
	ProducerID                            *ServiceCode      `json:"ProducerID"`                            // OBX-15
	ResponsibleObserver                   *XCN              `json:"ResponsibleObserver"`                   // OBX-16
	ObservationMethod                     *ServiceCode      `json:"ObservationMethod"`                     // OBX-17
	EquipmentInstanceIdentifier           *EntityIdentifier `json:"EquipmentInstanceIdentifier"`           // OBX-18
	DateTimeOfTheAnalysis                 string            `json:"DateTimeOfTheAnalysis"`                 // OBX-19
	ObservationSite                       *ServiceCode      `json:"ObservationSite"`                       // OBX-20
	ObservationInstanceIdentifier         *EntityIdentifier `json:"ObservationInstanceIdentifier"`         // OBX-21
	MoodCode                              *ServiceCode      `json:"MoodCode"`                              // OBX-22
	PerformingOrganizationName            *XON              `json:"PerformingOrganizationName"`            // OBX-23
	PerformingOrganizationAddress         *XAD              `json:"PerformingOrganizationAddress"`         // OBX-24
	PerformingOrganizationMedicalDirector *XCN              `json:"PerformingOrganizationMedicalDirector"` // OBX-25
}

func NewOBXSegment(p *random.Patient) *OBX {

	obx := &OBX{
		SetID: "1",
		// NM is a numeric value. A document notification swaps this for TX
		// when the message type is applied, since a note travels as text.
		ValueType:                             "NM",
		ObservationIdentifier:                 p.Order.Test.Code + "^" + p.Order.Test.Name + "^LN",
		ObservationValue:                      p.Order.ResultValue,
		Units:                                 &ServiceCode{Identifier: p.Order.Test.Units},
		ReferencesRange:                       p.Order.Test.ReferenceRange(),
		AbnormalFlags:                         p.Order.AbnormalFlag,
		ObservationResultStatus:               p.Order.ResultStatus,
		DateTimeOfTheObservation:              p.EventDate.HL7(),
		ProducerID:                            &ServiceCode{},
		ResponsibleObserver:                   &XCN{},
		ObservationMethod:                     &ServiceCode{},
		EquipmentInstanceIdentifier:           &EntityIdentifier{},
		ObservationSite:                       &ServiceCode{},
		ObservationInstanceIdentifier:         &EntityIdentifier{},
		MoodCode:                              &ServiceCode{},
		PerformingOrganizationName:            &XON{},
		PerformingOrganizationAddress:         &XAD{},
		PerformingOrganizationMedicalDirector: &XCN{},
	}

	return obx
}
