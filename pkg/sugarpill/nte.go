package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type NTE struct {
	SetID           string       `json:"SetID"`           // NTE-1
	SourceOfComment string       `json:"SourceOfComment"` // NTE-2
	Comment         string       `json:"Comment"`         // NTE-3
	CommentType     *ServiceCode `json:"CommentType"`     // NTE-4
}

func NewNTESegment(p *random.Patient) *NTE {

	nte := &NTE{
		CommentType: &ServiceCode{},
	}

	return nte
}
