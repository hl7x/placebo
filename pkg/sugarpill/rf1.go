package sugarpill

import (
	"fmt"

	"github.com/hl7x/placebo/pkg/random"
)

// RF1 carries the referral a referral message is about. It is the segment
// that makes a REF message a REF message.
type RF1 struct {
	ReferralStatus             *ServiceCode      `json:"ReferralStatus"`             // RF1-1
	ReferralPriority           *ServiceCode      `json:"ReferralPriority"`           // RF1-2
	ReferralType               *ServiceCode      `json:"ReferralType"`               // RF1-3
	ReferralDisposition        *ServiceCode      `json:"ReferralDisposition"`        // RF1-4
	ReferralCategory           *ServiceCode      `json:"ReferralCategory"`           // RF1-5
	OriginatingReferralID      *EntityIdentifier `json:"OriginatingReferralID"`      // RF1-6
	EffectiveDate              string            `json:"EffectiveDate"`              // RF1-7
	ExpirationDate             string            `json:"ExpirationDate"`             // RF1-8
	ProcessDate                string            `json:"ProcessDate"`                // RF1-9
	ReferralReason             *ServiceCode      `json:"ReferralReason"`             // RF1-10
	ExternalReferralIdentifier *EntityIdentifier `json:"ExternalReferralIdentifier"` // RF1-11
}

func NewRF1Segment(p *random.Patient) *RF1 {

	referral := &EntityIdentifier{
		EntityIdentifier: fmt.Sprint(p.VisitId),
		NamespaceID:      "PLACEBO",
	}

	rf1 := &RF1{
		ReferralStatus:        &ServiceCode{Identifier: "P", Text: "PENDING"},
		ReferralPriority:      &ServiceCode{Identifier: "R", Text: "ROUTINE"},
		ReferralType:          &ServiceCode{Identifier: "MED", Text: "MEDICAL"},
		ReferralDisposition:   &ServiceCode{},
		ReferralCategory:      &ServiceCode{Identifier: "O", Text: "OUTPATIENT"},
		OriginatingReferralID: referral,
		EffectiveDate:         p.EventDate.HL7(),
		// A referral an interface can act on has to still be open, so the
		// window runs forward from the event rather than around it.
		ExpirationDate:             p.Appointment.HL7(),
		ProcessDate:                p.EventDate.HL7(),
		ReferralReason:             &ServiceCode{Identifier: "S", Text: "SECOND OPINION"},
		ExternalReferralIdentifier: &EntityIdentifier{},
	}

	return rf1
}
