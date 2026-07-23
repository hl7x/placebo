package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type Placer struct {
	UniquePlacerID    string `json:"UniquePlacerID"`    // OBR-2.1
	PlacerApplication string `json:"PlacerApplication"` // OBR-2.2
}

type Filler struct {
	UniqueFillerID    string `json:"UniqueFillerID"`    // OBR-3.1
	FillerApplication string `json:"FillerApplication"` // OBR-3.2
}

type OBRDateTime struct {
	EventTime         string `json:"EventTime"`         // OBR-x.1
	DegreeOfPrecision string `json:"DegreeOfPrecision"` // OBR-x.2
}

type OBRVolume struct {
	Quantity string `json:"Quantity"` // OBR-9.1
	Units    string `json:"Units"`    // OBR-9.2
}

type OBRReceptCode struct {
	IDNumber   string `json:"IDNumber"`   // OBR-10.1
	FamilyName string `json:"FamilyName"` // OBR-10.2
	GivenName  string `json:"GivenName"`  // OBR-10.3
	MiddleName string `json:"MiddleName"` // OBR-10.4
	Suffix     string `json:"Suffix"`     // OBR-10.5
	Prefix     string `json:"Prefix"`     // OBR-10.6
	Degree     string `json:"Degree"`     // OBR-10.7
	Source     string `json:"Source"`     // OBR-10.8
}

type OBRSource struct {
	SpecimenSourceCode string `json:"SpecimenSourceCode"` // OBR-15.1
	Additives          string `json:"Additives"`          // OBR-15.2
	FreeText           string `json:"FreeText"`           // OBR-15.3
	BodySite           string `json:"BodySite"`           // OBR-15.4
	SiteModifier       string `json:"SiteModifier"`       // OBR-15.5
}

type Charge struct {
	DollarAmount string `json:"DollarAmount"` // OBR-23.1
	ChargeCode   string `json:"ChargeCode"`   // OBR-23.2
}

type ParentResultCode struct {
	ObservationID     string `json:"ObservationID"`     // OBR-26.1
	ParentResultSubID string `json:"ParentResultSubID"` // OBR-26.2
	ObservationResult string `json:"ObservationResult"` // OBR-26.3
}

type QuantityTiming struct {
	QuantityAmount string `json:"QuantityAmount"` // OBR-27.1
	Interval       string `json:"Interval"`       // OBR-27.2
	Duration       string `json:"Duration"`       // OBR-27.3
	StartDate      string `json:"StartDate"`      // OBR-27.4
	EndDate        string `json:"EndDate"`        // OBR-27.5
	Priority       string `json:"Priority"`       // OBR-27.6
	Condition      string `json:"Condition"`      // OBR-27.7
	Text           string `json:"Text"`           // OBR-27.8
	Conjunction    string `json:"Conjunction"`    // OBR-27.9
	OrderSequence  string `json:"OrderSequence"`  // OBR-27.10
}

type OBRParentNumber struct {
	PlacerOrderNumber string `json:"PlacerOrderNumber"` // OBR-29.1
	FillerOrderNumber string `json:"FillerOrderNumber"` // OBR-29.2
}

type Technician struct {
	Technician string `json:"Technician"` // OBR-x.1
	StartDate  string `json:"StartDate"`  // OBR-x.2
	EndDate    string `json:"EndDate"`    // OBR-x.3
	Location   string `json:"Location"`   // OBR-x.4
}

type OBR struct {
	SetID                string            `json:"SetID"`                // OBR-1
	PlacerOrderNumber    *Placer           `json:"PlacerOrderNumber"`    // OBR-2
	FillerOrderNumber    *Filler           `json:"FillerOrderNumber"`    // OBR-3
	UniversalServiceID   *ServiceCode      `json:"UniversalService"`     // OBR-4
	Priority             string            `json:"Priority"`             // OBR-5
	RequestDate          *OBRDateTime      `json:"Request"`              // OBR-6
	ObservationDate      *OBRDateTime      `json:"Observation"`          // OBR-7
	ObservationEndDate   *OBRDateTime      `json:"ObservationEndDate"`   // OBR-8
	CollectionVolume     *OBRVolume        `json:"CollectionVolume"`     // OBR-9
	CollectorID          *OBRReceptCode    `json:"Collector"`            // OBR-10
	SpecimenAction       string            `json:"SpecimenAction"`       // OBR-11
	DangerCode           *ServiceCode      `json:"ServiceCode"`          // OBR-12
	ClinicalInfo         string            `json:"ClinicalInfo"`         // OBR-13
	SpecimenRecivedDate  *OBRDateTime      `json:"SpecimenRecivedDate"`  // OBR-14
	SpecimenSource       *OBRSource        `json:"SpecimenSource"`       // OBR-15
	OrderingProvider     *OBRReceptCode    `json:"OrderingProvider"`     // OBR-16
	CallbackPhone        string            `json:"CallbackPhone"`        // OBR-17
	PlacerField1         string            `json:"PlacerField1"`         // OBR-18
	PlacerField2         string            `json:"PlacerField2"`         // OBR-19
	FillerField1         string            `json:"FillerField1"`         // OBR-20
	FillerField2         string            `json:"FillerField2"`         // OBR-21
	ResultReportDate     *OBRDateTime      `json:"ResultReportTime"`     // OBR-22
	ChargeToPractice     *Charge           `json:"ChargeToPractice"`     // OBR-23
	DiagnosticService    string            `json:"DiagnosticService"`    // OBR-24
	ResultStatus         string            `json:"ResultStatus"`         // OBR-25
	ParentResult         *ParentResultCode `json:"ParentResult"`         // OBR-26
	Quantity             *QuantityTiming   `json:"QuantityTiming"`       // OBR-27
	ResultCopyTo         *OBRReceptCode    `json:"ResultCopyTo"`         // OBR-28
	ParentNumber         *OBRParentNumber  `json:"ParentNumber"`         // OBR-29
	TransportationMode   string            `json:"TransportationMode"`   // OBR-30
	ReasonForStudy       *ServiceCode      `json:"ReasonForStudy"`       // OBR-31
	PrincipalInterpreter *Technician       `json:"PrincipalInterpreter"` // OBR-32
	AssistantInterpreter *Technician       `json:"AssistantInterpreter"` // OBR-33
	Technician           *Technician       `json:"Technician"`           // OBR-34
	Transcription        *Technician       `json:"Transcription"`        // OBR-35
	ScheduledDate        *OBRDateTime      `json:"ScheduledDate"`        // OBR-36
}

func NewOBRSegment(p *random.Patient) *OBR {

	obr := &OBR{
		PlacerOrderNumber:    &Placer{},
		FillerOrderNumber:    &Filler{},
		UniversalServiceID:   &ServiceCode{},
		RequestDate:          &OBRDateTime{},
		ObservationDate:      &OBRDateTime{},
		ObservationEndDate:   &OBRDateTime{},
		CollectionVolume:     &OBRVolume{},
		CollectorID:          &OBRReceptCode{},
		DangerCode:           &ServiceCode{},
		SpecimenRecivedDate:  &OBRDateTime{},
		SpecimenSource:       &OBRSource{},
		OrderingProvider:     &OBRReceptCode{},
		ResultReportDate:     &OBRDateTime{},
		ChargeToPractice:     &Charge{},
		ParentResult:         &ParentResultCode{},
		Quantity:             &QuantityTiming{},
		ResultCopyTo:         &OBRReceptCode{},
		ParentNumber:         &OBRParentNumber{},
		ReasonForStudy:       &ServiceCode{},
		PrincipalInterpreter: &Technician{},
		AssistantInterpreter: &Technician{},
		Technician:           &Technician{},
		Transcription:        &Technician{},
		ScheduledDate:        &OBRDateTime{},
	}

	return obr
}
