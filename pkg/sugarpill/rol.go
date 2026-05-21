package sugarpill

import "placebo/pkg/random"

type ROL struct {
	RoleInstanceID             *EntityIdentifier `json:"RoleInstanceID"`             // ROL-1
	ActionCode                 string            `json:"ActionCode"`                 // ROL-2
	Role                       *ServiceCode      `json:"Role"`                       // ROL-3
	RolePerson                 *XCN              `json:"RolePerson"`                 // ROL-4
	RoleBeginDateTime          string            `json:"RoleBeginDateTime"`          // ROL-5
	RoleEndDateTime            string            `json:"RoleEndDateTime"`            // ROL-6
	RoleDuration               *ServiceCode      `json:"RoleDuration"`               // ROL-7
	RoleActionReason           *ServiceCode      `json:"RoleActionReason"`           // ROL-8
	ProviderType               *ServiceCode      `json:"ProviderType"`               // ROL-9
	OrganizationUnitType       *ServiceCode      `json:"OrganizationUnitType"`       // ROL-10
	OfficeHomeAddressBirthplace *XAD             `json:"OfficeHomeAddressBirthplace"` // ROL-11
	Phone                      *XTN              `json:"Phone"`                      // ROL-12
	PersonIdentifier           *PatientLocation  `json:"PersonIdentifier"`           // ROL-13
}

func NewROLSegment(p *random.Patient) *ROL {

	rol := &ROL{
		RoleInstanceID:              &EntityIdentifier{},
		Role:                        &ServiceCode{},
		RolePerson:                  &XCN{},
		RoleDuration:                &ServiceCode{},
		RoleActionReason:            &ServiceCode{},
		ProviderType:                &ServiceCode{},
		OrganizationUnitType:        &ServiceCode{},
		OfficeHomeAddressBirthplace: &XAD{},
		Phone:                       &XTN{},
		PersonIdentifier:            &PatientLocation{},
	}

	return rol
}
