package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type DateRange struct {
	StartDate string `json:"StartDate"` // X.x.1
	EndDate   string `json:"EndDate"`   // X.x.2
}

type ARV struct {
	SetID                        string       `json:"SetID"`                        // ARV-1
	AccessRestrictionCode        *ServiceCode `json:"AccessRestrictionCode"`        // ARV-2
	AccessRestrictionValue       *ServiceCode `json:"AccessRestrictionValue"`       // ARV-3
	AccessRestrictionReason      *ServiceCode `json:"AccessRestrictionReason"`      // ARV-4
	AccessRestrictionInstruction string       `json:"AccessRestrictionInstruction"` // ARV-5
	AccessRestrictionDate        *DateRange   `json:"AccessRestrictionDate"`        // ARV-6
}

func NewARVSegment(p *random.Patient) *ARV {

	arv := &ARV{
		AccessRestrictionCode:   &ServiceCode{},
		AccessRestrictionValue:  &ServiceCode{},
		AccessRestrictionReason: &ServiceCode{},
		AccessRestrictionDate:   &DateRange{},
	}

	return arv
}
