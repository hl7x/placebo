package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type DG1 struct {
	SetID                 string            `json:"SetID"`                   // DG1-1
	DiagnosisMethod       string            `json:""DiagnosisMethod`         // DG1-2
	DiagnosisCode         *ServiceCode      `json:"DiagnosisCode"`           // DG1-3
	Description           string            `json:"Description"`             // DG1-4
	DateTime              string            `json:"DateTime"`                // DG1-5
	Type                  string            `json:"Type"`                    // DG1-6
	MajorCategory         string            `json:"MajorCategory"`           // DG1-7
	RelatedGroup          string            `json:"RelatedGroup"`            // DG1-8
	ApprovalIndicator     string            `json:"ApprovalIndicator"`       // DG1-9
	GrouperReviewCode     string            `json:"GrouperReviewCode"`       // DG1-10
	OutlierType           string            `json:"OutlierType"`             // DG1-11
	OutlierDays           string            `json:"OutlierDays"`             // DG1-12
	OutlierCost           string            `json:"OutlierCost"`             // DG1-13
	GrouperVersionType    string            `json:"GrouperVersionType"`      // DG1-14
	Priority              string            `json:"Priority"`                // DG1-15
	Clinician             *XCN              `json:"Clinician"`               // DG1-16
	Classification        string            `json:"Classification"`          // DG1-17
	ConfidentialIndicator string            `json:"ConfidentialIdenticator"` // DG1-18
	AttestationDateTime   string            `json:"AttestationDateTime"`     // DG1-19
	Identifier            *EntityIdentifier `json:"Identifier"`              // DG1-20
	ActionCode            string            `json:"ActionCode"`              // DG1-21
	ParentDiagnosis       *EntityIdentifier `json:"ParentDiagnosis"`         // DG1-22
	DRGCCLValue           *ServiceCode      `json:"DRGCCLValue"`             // DG1-23
	DRGGroupingUsage      string            `json:"DRGGroupingUsage"`        // DG1-24
	DRGDiagnosisStatus    string            `json:"DRGDiagnosisStatus"`      // DG1-25
	POAIndicator          string            `json:"POAIndicator"`            // DG1-26
}

func NewDG1Segment(p *random.Patient) *DG1 {

	dg1 := &DG1{
		DiagnosisCode:   &ServiceCode{},
		Clinician:       &XCN{},
		Identifier:      &EntityIdentifier{},
		ParentDiagnosis: &EntityIdentifier{},
		DRGCCLValue:     &ServiceCode{},
	}

	return dg1
}
