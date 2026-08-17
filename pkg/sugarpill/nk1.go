package sugarpill

import "github.com/hl7x/placebo/pkg/random"

type NK1 struct {
	SetID                                    string       `json:"SetID"`                                    // NK1-1
	Name                                     *XPN         `json:"Name"`                                     // NK1-2
	Relationship                             *ServiceCode `json:"Relationship"`                             // NK1-3
	Address                                  *XAD         `json:"Address"`                                  // NK1-4
	PhoneNumber                              *XTN         `json:"PhoneNumber"`                              // NK1-5
	BusinessPhoneNumber                      *XTN         `json:"BusinessPhoneNumber"`                      // NK1-6
	ContactRole                              *ServiceCode `json:"ContactRole"`                              // NK1-7
	StartDate                                string       `json:"StartDate"`                                // NK1-8
	EndDate                                  string       `json:"EndDate"`                                  // NK1-9
	NextOfKinAssociatedPartiesJobTitle       string       `json:"NextOfKinAssociatedPartiesJobTitle"`       // NK1-10
	NextOfKinAssociatedPartiesJobCode        *JCC         `json:"NextOfKinAssociatedPartiesJobCode"`        // NK1-11
	NextOfKinAssociatedPartiesEmployeeNumber *CX          `json:"NextOfKinAssociatedPartiesEmployeeNumber"` // NK1-12
	OrganizationName                         *XON         `json:"OrganizationName"`                         // NK1-13
	MaritalStatus                            *ServiceCode `json:"MaritalStatus"`                            // NK1-14
	AdministrativeSex                        string       `json:"AdministrativeSex"`                        // NK1-15
	DateOfBirth                              string       `json:"DateOfBirth"`                              // NK1-16
	LivingDependency                         string       `json:"LivingDependency"`                         // NK1-17
	AmbulatoryStatus                         string       `json:"AmbulatoryStatus"`                         // NK1-18
	Citizenship                              string       `json:"Citizenship"`                              // NK1-19
	PrimaryLanguage                          *ServiceCode `json:"PrimaryLanguage"`                          // NK1-20
	LivingArrangement                        string       `json:"LivingArrangement"`                        // NK1-21
	PublicityCode                            *ServiceCode `json:"PublicityCode"`                            // NK1-22
	ProtectionIndicator                      string       `json:"ProtectionIndicator"`                      // NK1-23
	StudentIndicator                         string       `json:"StudentIndicator"`                         // NK1-24
	Religion                                 *ServiceCode `json:"Religion"`                                 // NK1-25
	MothersMaidenName                        *XPN         `json:"MothersMaidenName"`                        // NK1-26
	Nationality                              *ServiceCode `json:"Nationality"`                              // NK1-27
	EthnicGroup                              *ServiceCode `json:"EthnicGroup"`                              // NK1-28
	ContactReason                            *ServiceCode `json:"ContactReason"`                            // NK1-29
	ContactPersonsName                       *XPN         `json:"ContactPersonsName"`                       // NK1-30
	ContactPersonsTelephoneNumber            *XTN         `json:"ContactPersonsTelephoneNumber"`            // NK1-31
	ContactPersonsAddress                    *XAD         `json:"ContactPersonsAddress"`                    // NK1-32
	NextOfKinAssociatedPartysIdentifiers     *CX          `json:"NextOfKinAssociatedPartysIdentifiers"`     // NK1-33
	JobStatus                                string       `json:"JobStatus"`                                // NK1-34
	Race                                     *ServiceCode `json:"Race"`                                     // NK1-35
	Handicap                                 string       `json:"Handicap"`                                 // NK1-36
	ContactPersonSocialSecurityNumber        string       `json:"ContactPersonSocialSecurityNumber"`        // NK1-37
	NextOfKinBirthPlace                      string       `json:"NextOfKinBirthPlace"`                      // NK1-38
	VIPIndicator                             string       `json:"VIPIndicator"`                             // NK1-39
}

func NewNK1Segment(P *random.Patient) *NK1 {

	nk1 := &NK1{
		Name:                                     &XPN{},
		Relationship:                             &ServiceCode{},
		Address:                                  &XAD{},
		PhoneNumber:                              &XTN{},
		BusinessPhoneNumber:                      &XTN{},
		ContactRole:                              &ServiceCode{},
		NextOfKinAssociatedPartiesJobCode:        &JCC{},
		NextOfKinAssociatedPartiesEmployeeNumber: &CX{},
		OrganizationName:                         &XON{},
		MaritalStatus:                            &ServiceCode{},
		PrimaryLanguage:                          &ServiceCode{},
		PublicityCode:                            &ServiceCode{},
		Religion:                                 &ServiceCode{},
		MothersMaidenName:                        &XPN{},
		Nationality:                              &ServiceCode{},
		EthnicGroup:                              &ServiceCode{},
		ContactReason:                            &ServiceCode{},
		ContactPersonsName:                       &XPN{},
		ContactPersonsTelephoneNumber:            &XTN{},
		ContactPersonsAddress:                    &XAD{},
		NextOfKinAssociatedPartysIdentifiers:     &CX{},
		Race:                                     &ServiceCode{},
	}

	return nk1
}
