package sugarpill

import "placebo/pkg/random"

type EVN struct {
	EventTypeCode        string `json:"EventTypeCode"`        // EVN-1
	RecordedDateTime     string `json:"RecordedDateTime"`     // EVN-2
	DateTimePlannedEvent string `json:"DateTimePlannedEvent"` // EVN-3
	EventReasonCode      string `json:"EventReasonCode"`      // EVN-4
	OperatorID           string `json:"OperatorID"`           // EVN-5
	EventOccurred        string `json:"EventOccurred"`        // EVN-6
	EventFacility        string `json:"EventFacility"`        // EVN-7
}

func NewEVNSegment(p *random.Patient) *EVN {

	evn := &EVN{
		EventTypeCode:    "A01", //placeholder
		RecordedDateTime: p.Hl7Info.HL7Event,
	}

	return evn
}
