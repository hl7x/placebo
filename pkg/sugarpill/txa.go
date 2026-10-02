package sugarpill

import (
	"github.com/hl7x/placebo/pkg/random"
)

// TXA carries the document a document notification is about. It is the
// segment that makes an MDM message an MDM message.
type TXA struct {
	SetID                         string            `json:"SetID"`                         // TXA-1
	DocumentType                  string            `json:"DocumentType"`                  // TXA-2
	DocumentContentPresentation   string            `json:"DocumentContentPresentation"`   // TXA-3
	ActivityDateTime              string            `json:"ActivityDateTime"`              // TXA-4
	PrimaryActivityProvider       *XCN              `json:"PrimaryActivityProvider"`       // TXA-5
	OriginationDateTime           string            `json:"OriginationDateTime"`           // TXA-6
	TranscriptionDateTime         string            `json:"TranscriptionDateTime"`         // TXA-7
	EditDateTime                  string            `json:"EditDateTime"`                  // TXA-8
	OriginatorCodeName            *XCN              `json:"OriginatorCodeName"`            // TXA-9
	AssignedDocumentAuthenticator *XCN              `json:"AssignedDocumentAuthenticator"` // TXA-10
	TranscriptionistCodeName      *XCN              `json:"TranscriptionistCodeName"`      // TXA-11
	UniqueDocumentNumber          *EntityIdentifier `json:"UniqueDocumentNumber"`          // TXA-12
	ParentDocumentNumber          *EntityIdentifier `json:"ParentDocumentNumber"`          // TXA-13
	PlacerOrderNumber             *EntityIdentifier `json:"PlacerOrderNumber"`             // TXA-14
	FillerOrderNumber             *EntityIdentifier `json:"FillerOrderNumber"`             // TXA-15
	UniqueDocumentFileName        string            `json:"UniqueDocumentFileName"`        // TXA-16
	DocumentCompletionStatus      string            `json:"DocumentCompletionStatus"`      // TXA-17
	DocumentConfidentialStatus    string            `json:"DocumentConfidentialStatus"`    // TXA-18
	DocumentAvailabilityStatus    string            `json:"DocumentAvailabilityStatus"`    // TXA-19
	DocumentStorageStatus         string            `json:"DocumentStorageStatus"`         // TXA-20
	DocumentChangeReason          string            `json:"DocumentChangeReason"`          // TXA-21
	AuthenticationPerson          *XCN              `json:"AuthenticationPerson"`          // TXA-22
	DistributedCopies             *XCN              `json:"DistributedCopies"`             // TXA-23
}

func NewTXASegment(p *random.Patient) *TXA {

	provider := &XCN{
		ID:         p.Provider.ID,
		FamilyName: p.Provider.LastName,
		GivenName:  p.Provider.FirstName,
	}

	document := &EntityIdentifier{
		EntityIdentifier: p.Document.Number,
		NamespaceID:      "PLACEBO",
	}

	txa := &TXA{
		SetID:        "1",
		DocumentType: p.Document.Type.Code,
		// TX is a plain text body. The OBX segments alongside this one are
		// where the document's lines actually travel.
		DocumentContentPresentation:   "TX",
		ActivityDateTime:              p.EventDate.HL7(),
		PrimaryActivityProvider:       provider,
		OriginationDateTime:           p.EventDate.HL7(),
		TranscriptionDateTime:         p.EventDate.HL7(),
		OriginatorCodeName:            provider,
		AssignedDocumentAuthenticator: &XCN{},
		TranscriptionistCodeName:      &XCN{},
		UniqueDocumentNumber:          document,
		ParentDocumentNumber:          &EntityIdentifier{},
		PlacerOrderNumber:             &EntityIdentifier{},
		FillerOrderNumber:             &EntityIdentifier{},
		UniqueDocumentFileName:        p.Document.Type.Name,
		DocumentCompletionStatus:      p.Document.Status,
		DocumentAvailabilityStatus:    "AV",
		DocumentStorageStatus:         "AC",
		AuthenticationPerson:          &XCN{},
		DistributedCopies:             &XCN{},
	}

	return txa
}
