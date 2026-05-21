package sugarpill

// Generic Extended Address Used In Many cases
type XAD struct {
	StreetAddress              string `json:"StreetAddress"`              // X.x.1
	OtherDesignation           string `json:"OtherDesignation"`           // X.x.2
	City                       string `json:"City"`                       // X.x.3
	State                      string `json:"State"`                      // X.x.4
	ZipCode                    string `json:"ZipCode"`                    // X.x.5
	Country                    string `json:"Country"`                    // X.x.6
	AddressType                string `json:"AddressType"`                // X.x.7
	OtherGeographicDesignation string `json:"OtherGeographicDesignation"` // X.x.8
	CountyParishCode           string `json:"CountryParishCode"`          // X.x.9
	CensusTract                string `json:"CensusTract"`                // X.x.10
}

// Extended Persons Name Generic Used in Other Cases
type XPN struct {
	LastName               string `json:"LastName"`               // PID-5.1
	FirstName              string `json:"FirstName"`              // PID-5.2
	MiddleInitial          string `json:"MiddleInitial"`          // PID-5.3
	Suffix                 string `json:"Suffix"`                 // PID-5.4
	Prefix                 string `json:"Prefix"`                 // PID-5.5
	Degree                 string `json:"Degree"`                 // PID-5.6
	NameTypeCode           string `json:"NameTypeCode"`           // PID-5.7
	NameRepresentationCode string `json:"NameRepresentationCode"` // PID-5.8
	NameContext            string `json:"NameContext"`            // PID-5.9
	NameValidityRange      string `json:"NameValidityRange"`      // PID-5.10
	NameAssemblyOrder      string `json:"NameAssemblyOrder"`      // PID-5.11
	EffectiveDate          string `json:"EffectiveDate"`          // PID-5.12
	ExpirationDate         string `json:"ExpirationDate"`         // PID-5.13
	ProfessionalSuffix     string `json:"ProfessionalSuffix"`     // PID-5.14
}

type XTN struct {
	TelephoneNumber                   string `json:"TelephoneNumber"`                   // XTN.1
	TelecommunicationUseCode          string `json:"TelecommunicationUseCode"`          // XTN.2
	TelecommunicationEquipmentType    string `json:"TelecommunicationEquipmentType"`    // XTN.3
	EmailAddress                      string `json:"EmailAddress"`                      // XTN.4
	CountryCode                       string `json:"CountryCode"`                       // XTN.5
	AreaCode                          string `json:"AreaCode"`                          // XTN.6
	LocalNumber                       string `json:"LocalNumber"`                       // XTN.7
	Extension                         string `json:"Extension"`                         // XTN.8
	AnyText                           string `json:"AnyText"`                           // XTN.9
	ExtensionPrefix                   string `json:"ExtensionPrefix"`                   // XTN.10
	SpeedDialCode                     string `json:"SpeedDialCode"`                     // XTN.11
	UnformattedTelephoneNumber        string `json:"UnformattedTelephoneNumber"`        // XTN.12
	EffectiveStartDate                string `json:"EffectiveStartDate"`                // XTN.13
	ExpirationDate                    string `json:"ExpirationDate"`                    // XTN.14
	ExpirationReason                  string `json:"ExpirationReason"`                  // XTN.15
	ProtectionCode                    string `json:"ProtectionCode"`                    // XTN.16
	SharedTelecommunicationIdentifier string `json:"SharedTelecommunicationIdentifier"` // XTN.17
	PreferenceOrder                   string `json:"PreferenceOrder"`                   // XTN.18
}

// Generic Extended Composite ID and Name For Persons Used in Many Cases
type XCN struct {
	ID                     string `json:"ID"`                     // PV2-13.1
	FamilyName             string `json:"FamilyName"`             // PV2-13.2
	GivenName              string `json:"GivenName"`              // PV2-13.3
	MiddleInitialOrName    string `json:"MiddleInitialOrName"`    // PV2-13.4
	Suffix                 string `json:"Suffix"`                 // PV2-13.5
	Prefix                 string `json:"Prefix"`                 // PV2-13.6
	Degree                 string `json:"Degree"`                 // PV2-13.7
	SourceTable            string `json:"SourceTable"`            // PV2-13.8
	AssigningAuthority     string `json:"AssigningAuthority"`     // PV2-13.9
	NameTypeCode           string `json:"NameTypeCode"`           // PV2-13.10
	IdentifierCheckDigit   string `json:"IdentifierCheckDigit"`   // PV2-13.11
	CheckDigitScheme       string `json:"CheckDigitScheme"`       // PV2-13.12
	IdentifierTypeCode     string `json:"IdentifierTypeCode"`     // PV2-13.13
	AssigningFacility      string `json:"AssigningFacility"`      // PV2-13.14
	NameRepresentationCode string `json:"NameRepresentationCode"` // PV2-13.15
	NameContext            string `json:"NameContext"`            // PV2-13.16
	NameValidityRange      string `json:"NameValidityRange"`      // PV2-13.17
	NameAssemblyOrder      string `json:"NameAssemblyOrder"`      // PV2-13.18
	EffectiveDate          string `json:"EffectiveDate"`          // PV2-13.19
	ExpirationDate         string `json:"ExpirationDate"`         // PV2-13.20
	ProfessionalSuffix     string `json:"ProfessionalSuffix"`     // PV2-13.21
	AssigningJurisdiction  string `json:"AssigningJurisdiction"`  // PV2-13.22
	AssigningAgencyOrDept  string `json:"AssigningAgencyOrDept"`  // PV2-13.23
}

// Extended Composite Name and ID for Organization Used in many cases
type XON struct {
	OrganizationName       string `json:"OrganizationName"`       // PV2-23.1
	TypeCode               string `json:"TypeCode"`               // PV2-23.2
	IDNumber               string `json:"IDNumber"`               // PV2-23.3
	CheckDigit             string `json:"CheckDigit"`             // PV2-23.4
	CheckDigitScheme       string `json:"CheckDigitScheme"`       // PV2-23.5
	AssigningAuthority     string `json:"AssigningAuthority"`     // PV2-23.6
	IdentifierTypeCode     string `json:"IdentifierTypeCode"`     // PV2-23.7
	AssigningFacility      string `json:"AssigningFacility"`      // PV2-23.8
	RepresentationCode     string `json:"RepresentationCode"`     // PV2-23.9
	OrganizationIdentifier string `json:"OrganizationIdentifier"` // PV2-23.10
}

// Generic EI Used in Many other cases
type EntityIdentifier struct {
	EntityIdentifier string `json:"EntityIdentifier"` // X-x.1
	NamespaceID      string `json:"NamespaceID"`      // X-x.2
	UniversalID      string `json:"UniversalID"`      // X-x.3
	UniversalType    string `json:"UniversalType"`    // X-x.4
}

// Generic Nested Section Used in Many Cases
type ServiceCode struct {
	Identifier            string `json:"Identifier"`            // OBR-x.1
	Text                  string `json:"Text"`                  // OBR-x.2
	CodingSystem          string `json:"CodingSystem"`          // OBR-x.3
	AlternateIdentifier   string `json:"AlternateIdentifier"`   // OBR-x.4
	AlternateText         string `json:"AlternateText"`         // OBR-x.5
	AlternateCodingSystem string `json:"AlternateCodingSystem"` // OBR-x.6
}

type CX struct {
	IDNumber              string `json:"IDNumber"`              // CX.1
	CheckDigit            string `json:"CheckDigit"`            // CX.2
	CheckDigitScheme      string `json:"CheckDigitScheme"`      // CX.3
	AssigningAuthority    string `json:"AssigningAuthority"`    // CX.4
	IdentifierTypeCode    string `json:"IdentifierTypeCode"`    // CX.5
	AssigningFacility     string `json:"AssigningFacility"`     // CX.6
	EffectiveDate         string `json:"EffectiveDate"`         // CX.7
	ExpirationDate        string `json:"ExpirationDate"`        // CX.8
	AssigningJurisdiction string `json:"AssigningJurisdiction"` // CX.9
	AssigningAgencyOrDept string `json:"AssigningAgencyOrDept"` // CX.10
}

type JCC struct {
	JobCode        string `json:"JobCode"`
	JobClass       string `json:"JobClass"`
	JobDescription string `json:"JobDescription"`
}

type PatientLocation struct {
	PointOfCare         string `json:"PointOfCare"`         // PV1-3.1
	Room                string `json:"Room"`                // PV1-3.2
	Bed                 string `json:"Bed"`                 // PV1-3.3
	Facility            string `json:"Facility"`            // PV1-3.4
	LocationStatus      string `json:"LocationStatus"`      // PV1-3.5
	PatientLocationType string `json:"PatientLocationType"` // PV1-3.6
	Building            string `json:"Building"`            // PV1-3.7
	Floor               string `json:"Floor"`               // PV1-3.8
}
