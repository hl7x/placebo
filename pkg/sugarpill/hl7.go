package sugarpill

import (
	"encoding/json"
	"strings"

	"placebo/pkg/message"
	"placebo/pkg/random"
)

type MSH struct {
	Encode               string       `json:Encode`                 // MSH-2
	SendingApplication   string       `json:"SendingApplication"`   // MSH-3
	SendingFacility      string       `json:"SendingFacility"`      // MSH-4
	ReceivingApplication string       `json:"ReceivingApplication"` // MSH-5
	ReceivingFacility    string       `json:"ReceivingFacility"`    // MSH-6
	DateTimeOfMessage    string       `json:"DateTimeOfMessage"`    // MSH-7
	Security             string       `json:"Security"`             // MSH-8
	MessageType          *MessageType `json:"MessageType"`          // MSH-9
	MessageControlID     string       `json:"MessageControlID"`     // MSH-10
	ProcessingID         string       `json:"ProcessingID"`         // MSH-11
	VersionID            string       `json:"VersionID"`            // MSH-12
}

type MessageType struct {
	MessageCode  string `json:"MessageCode"`  // MSH-9.1
	TriggerEvent string `json:"TriggerEvent"` // MSH-9.2
}

type EVN struct {
	EventTypeCode        string `json:"EventTypeCode"`        // EVN-1
	RecordedDateTime     string `json:"RecordedDateTime"`     // EVN-2
	DateTimePlannedEvent string `json:"DateTimePlannedEvent"` // EVN-3
	EventReasonCode      string `json:"EventReasonCode"`      // EVN-4
	OperatorID           string `json:"OperatorID"`           // EVN-5
	EventOccurred        string `json:"EventOccurred"`        // EVN-6
	EventFacility        string `json:"EventFacility"`        // EVN-7
}

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
	TelephoneNumber                 string `json:"TelephoneNumber"`                 // XTN.1
	TelecommunicationUseCode        string `json:"TelecommunicationUseCode"`        // XTN.2
	TelecommunicationEquipmentType  string `json:"TelecommunicationEquipmentType"`  // XTN.3
	EmailAddress                    string `json:"EmailAddress"`                    // XTN.4
	CountryCode                     string `json:"CountryCode"`                     // XTN.5
	AreaCode                        string `json:"AreaCode"`                        // XTN.6
	LocalNumber                     string `json:"LocalNumber"`                     // XTN.7
	Extension                       string `json:"Extension"`                       // XTN.8
	AnyText                         string `json:"AnyText"`                         // XTN.9
	ExtensionPrefix                 string `json:"ExtensionPrefix"`                 // XTN.10
	SpeedDialCode                   string `json:"SpeedDialCode"`                   // XTN.11
	UnformattedTelephoneNumber      string `json:"UnformattedTelephoneNumber"`      // XTN.12
	EffectiveStartDate              string `json:"EffectiveStartDate"`              // XTN.13
	ExpirationDate                  string `json:"ExpirationDate"`                  // XTN.14
	ExpirationReason                string `json:"ExpirationReason"`                // XTN.15
	ProtectionCode                  string `json:"ProtectionCode"`                  // XTN.16
	SharedTelecommunicationIdentifier string `json:"SharedTelecommunicationIdentifier"` // XTN.17
	PreferenceOrder                 string `json:"PreferenceOrder"`                 // XTN.18
}

type PID struct {
	SetID                   string          `json:"SetID"`                   // PID-1
	PatientID               string          `json:"PatientID"`               // PID-2
	PatientIdentifierList   string          `json:"PatientIdentifierList"`   // PID-3
	AlternatePatientID      string          `json:"AlternatePatientID"`      // PID-4
	PatientName             *XPN    `json:"PatientName"`             // PID-5
	MotherMaidenName        string          `json:"MotherMaidenName"`        // PID-6
	DateOfBirth             string          `json:"DateOfBirth"`             // PID-7
	Sex                     string          `json:"Sex"`                     // PID-8
	PatientAlias            string          `json:"PatientAlias"`            // PID-9
	Race                    string          `json:"Race"`                    // PID-10
	PatientAddress          *XAD		`json:"PatientAddress"`          // PID-11
	CountyCode              string          `json:"CountyCode"`              // PID-12
	HomePhone               string          `json:"HomePhone"`               // PID-13
	BusinessPhone           string          `json:"BusinessPhone"`           // PID-14
	PrimaryLanguage         string          `json:"PrimaryLanguage"`         // PID-15
	MaritalStatus           string          `json:"MaritalStatus"`           // PID-16
	Religion                string          `json:"Religion"`                // PID-17
	PatientAccountNumber    string          `json:"PatientAccountNumber"`    // PID-18
	SSN                     string          `json:"SSN"`                     // PID-19
	DriversLicenseNumber    string          `json:"DriversLicenseNumber"`    // PID-20
	MotherIdentifier        string          `json:"MotherIdentifier"`        // PID-21
	EthnicGroup             string          `json:"EthnicGroup"`             // PID-22
	BirthPlace              string          `json:"BirthPlace"`              // PID-23
	MultipleBirthIndicator  string          `json:"MultipleBirthIndicator"`  // PID-24
	BirthOrder              string          `json:"BirthOrder"`              // PID-25
	Citizenship             string          `json:"Citizenship"`             // PID-26
	VeteranStatus           string          `json:"VeteranStatus"`           // PID-27
	Nationality             string          `json:"Nationality"`             // PID-28
	PatientDeathDateAndTime string          `json:"PatientDeathDateAndTime"` // PID-29
	PatientDeathIndicator   string          `json:"PatientDeathIndicator"`   // PID-30
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

type AttendingDoctor struct {
	IDNumber           string `json:"IDNumber"`           // PV1-7.1
	LastName           string `json:"LastName"`           // PV1-7.2
	FirstName          string `json:"FirstName"`          // PV1-7.3
	AssigningAuthority string `json:"AssigningAuthority"` // PV1-7.4
}

type PV1 struct {
	SetID                   string           `json:"SetID"`                   // PV1-1
	PatientClass            string           `json:"PatientClass"`            // PV1-2
	AssignedPatientLocation *PatientLocation `json:"AssignedPatientLocation"` // PV1-3
	AdmissionType           string           `json:"AdmissionType"`           // PV1-4
	PreadmitNumber          string           `json:"PreadmitNumber"`          // PV1-5
	PriorPatientLocation    string           `json:"PriorPatientLocation"`    // PV1-6
	AttendingDoctor         *AttendingDoctor `json:"AttendingDoctor"`         // PV1-7
	ReferringDoctor         string           `json:"ReferringDoctor"`         // Pv1-8
	ConsultingDoctor        string           `json:"ConsultingDoctor"`        // Pv1-9
	HospitalService         string           `json:"HospitalService"`         // PV1-10
	TemporaryLocation       string           `json:"TemporaryLocation"`       // Pv1-11
	PreadmitTestIndicator   string           `json:"PreadmitTestIndicator"`   // PV1-12
	ReadmissionIndicator    string           `json:"ReadmissionIndicator"`    // Pv1-13
	AdmitSource             string           `json:"AdmitSource"`             // PV1-14
	AmbulatoryStatus        string           `json:"AmbulatoryStatus"`        // PV1-15
	VIPIndicator            string           `json:"VIPIndicator"`            // PV1-16
	AdmittingDoctor         string           `json:"AdmittingDoctor"`         // PV1-17
	PatientType             string           `json:"PatientType"`             // PV1-18
	VisitNumber             int              `json:"VisitNumber"`             // PV1-19
	FinacialClass           string           `json:"FinacialClass"`           // PV1-20
	ChargePriceIndicator    string           `json:"ChargePriceIndicator"`    // PV1-21
	CourtesyCode            string           `json:"CourtesyCode"`            // PV1-22
	CreditRating            string           `json:"CreditRating"`            // PV1-23
	ContractCode            string           `json:"ContractCode"`            // PV1-24
	ContractEffectiveDate   string           `json:"ContractEffectiveDate"`   // PV1-25
	ContractAmount          string           `json:"ContractAmount"`          // PV1-26
	ContractPeriod          string           `json:"ContractPeriod"`          // PV1-27
	InterestCode            string           `json:"InterestCode"`            // PV1-28
	TransferToBadDebtCode   string           `json:"TransferToBadDebtCode"`   // PV1-29
	TransferToBadDebtDate   string           `json:"TransferToBadDebtDate"`   // PV1-30
	TransferToBadDebtAmount string           `json:"TransferToBadDebtAmount"` // PV1-31
	RecoveryBadDebtAmount   string           `json:"RecoveryBaddebtAmount"`   // PV1-32
	DeleteAccountIndicator  string           `json:"DeleteAccountIndicator"`  // PV1-33
	DeleteAccountDate       string           `json:"DeleteAccountDate"`       // PV1-34
	DischargeDisposition    string           `json:"DischargeDisposition"`    // PV1-35
	DischargedToLocation    string           `json:"DischargedToLocation"`    // PV1-36
	DietType                string           `json:"DietType"`                // PV1-37
	ServicingFacility       string           `json:"ServicingFacility"`       // PV1-38
	BedStatus               string           `json:"BedStatus"`               // PV1-39
	AccountStatus           string           `json:"AccountStatus"`           // PV1-40
	PendingLocation         string           `json:"PendingLocation"`         // PV1-41
	PriorTemporaryLocation  string           `json:"PiorTemporaryLocation"`   // PV1-42
	PreviousLocation        string           `json:"PreviousLocation"`        // PV1-43
	AdmitDateTime           string           `json:"AdmitDateTime"`           // PV1-44
	DischargeDateTime       string           `json:"DischargeDateTime"`       // PV1-45
	CurrentPatientBalance   string           `json:"CurrentPatientBalance"`   // Pv1-46
	TotalCharges            string           `json:"Totalcharges"`            // PV1-47
	TotalAdjustments        string           `json:"TotalAdjustments"`        // PV1-48
	TotalPayments           string           `json:"TotalPayments"`           // PV1-49
	AlternateVisitID        string           `json:"AlternateVisitID"`        // PV1-50
	VisitIndicator          string           `json:"VisitIndicator"`          // Pv1-51
}

type PV2 struct {
	PriorPendingLocation         *PatientLocation 	`json:"PriorPendingLocation"`         // PV2-1
	AccommodationCode            *ServiceCode 	`json:"AccommodationCode"`            // PV2-2
	AdmitReason                  *ServiceCode 	`json:"AdmitReason"`                  // PV2-3
	TransferReason               *ServiceCode 	`json:"TransferReason"`               // PV2-4
	PatientValuables             string 	`json:"PatientValuables"`             // PV2-5
	PatientValuablesLocation     string 	`json:"PatientValuablesLocation"`     // PV2-6
	VisitUserCode                string 	`json:"VisitUserCode"`                // PV2-7
	ExpectedAdmitDateTime        string 	`json:"ExpectedAdmitDateTime"`        // PV2-8
	ExpectedDischargeDateTime    string 	`json:"ExpectedDischargeDateTime"`    // PV2-9
	EstimatedLengthOfInpatientStay string 	`json:"EstimatedLengthOfInpatientStay"` // PV2-10
	ActualLengthOfInpatientStay  string 	`json:"ActualLengthOfInpatientStay"`  // PV2-11
	VisitDescription             string 	`json:"VisitDescription"`             // PV2-12
	ReferralSourceCode           *XCN 	`json:"ReferralSourceCode"`           // PV2-13
	PreviousServiceDate          string 	`json:"PreviousServiceDate"`          // PV2-14
	EmploymentIllnessRelated     string 	`json:"EmploymentIllnessRelated"`     // PV2-15
	PurgeStatusCode              string 	`json:"PurgeStatusCode"`              // PV2-16
	PurgeStatusDate              string 	`json:"PurgeStatusDate"`              // PV2-17
	SpecialProgramCode           string 	`json:"SpecialProgramCode"`           // PV2-18
	RetentionIndicator           string 	`json:"RetentionIndicator"`           // PV2-19
	ExpectedNumberOfInsurancePlans string 	`json:"ExpectedNumberOfInsurancePlans"` // PV2-20
	VisitPublicityCode           string 	`json:"VisitPublicityCode"`           // PV2-21
	VisitProtectionIndicator     string 	`json:"VisitProtectionIndicator"`     // PV2-22
	ClinicOrganizationName       *XON 	`json:"ClinicOrganizationName"`       // PV2-23
	PatientStatusCode            string 	`json:"PatientStatusCode"`            // PV2-24
	VisitPriorityCode            string 	`json:"VisitPriorityCode"`            // PV2-25
	PreviousTreatmentDate        string 	`json:"PreviousTreatmentDate"`        // PV2-26
	ExpectedDischargeDisposition string 	`json:"ExpectedDischargeDisposition"` // PV2-27
	SignatureOnFileDate          string 	`json:"SignatureOnFileDate"`          // PV2-28
	FirstSimilarIllnessDate      string 	`json:"FirstSimilarIllnessDate"`      // PV2-29
	PatientChargeAdjustmentCode  *ServiceCode 	`json:"PatientChargeAdjustmentCode"`  // PV2-30
	RecurringServiceCode         string 	`json:"RecurringServiceCode"`         // PV2-31
	BillingMediaCode             string 	`json:"BillingMediaCode"`             // PV2-32
	ExpectationOfBillTypeCode    string 	`json:"ExpectationOfBillTypeCode"`    // PV2-33
	MilitaryPartnershipCode      string 	`json:"VisitProtectionIndicator"`     // PV2-34
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
	OrganizationName	string		`json:"OrganizationName"`		// PV2-23.1
	TypeCode		string		`json:"TypeCode"`			// PV2-23.2
	IDNumber		string		`json:"IDNumber"`			// PV2-23.3
	CheckDigit		string		`json:"CheckDigit"`			// PV2-23.4
	CheckDigitScheme	string		`json:"CheckDigitScheme"`		// PV2-23.5
	AssigningAuthority	string		`json:"AssigningAuthority"`		// PV2-23.6
	IdentifierTypeCode	string		`json:"IdentifierTypeCode"`		// PV2-23.7
	AssigningFacility	string		`json:"AssigningFacility"`		// PV2-23.8
	RepresentationCode	string		`json:"RepresentationCode"`		// PV2-23.9
	OrganizationIdentifier	string		`json:"OrganizationIdentifier"`		// PV2-23.10
}

// Generic EI Used in Many other cases
type EntityIdentifier struct {
	EntityIdentifier	string		`json:"EntityIdentifier"`		// X-x.1
	NamespaceID		string		`json:"NamespaceID"`			// X-x.2
	UniversalID		string		`json:"UniversalID"`			// X-x.3
	UniversalType		string		`json:"UniversalType"`			// X-x.4
}

type OBX struct {
	SetID                        string `json:"SetID"`                        // OBX-1
	ValueType                    string `json:"ValueType"`                    // OBX-2
	ObservationIdentifier        string `json:"ObservationIdentifier"`        // OBX-3
	ObservationSubID             string `json:"ObservationSubID"`             // OBX-4
	ObservationValue             string `json:"ObservationValue"`             // OBX-5
	Units                        *ServiceCode `json:"Units"`                        // OBX-6
	ReferencesRange              string `json:"ReferencesRange"`              // OBX-7
	AbnormalFlags                string `json:"AbnormalFlags"`                // OBX-8
	Probability                  string `json:"Probability"`                  // OBX-9
	NatureOfAbnormalTest         string `json:"NatureOfAbnormalTest"`         // OBX-10
	ObservationResultStatus      string `json:"ObservationResultStatus"`      // OBX-11
	EffectiveDateOfReferenceRange string `json:"EffectiveDateOfReferenceRange"` // OBX-12
	UserDefinedAccessChecks      string `json:"UserDefinedAccessChecks"`      // OBX-13
	DateTimeOfTheObservation     string `json:"DateTimeOfTheObservation"`     // OBX-14
	ProducerID                   *ServiceCode `json:"ProducerID"`                   // OBX-15
	ResponsibleObserver          *XCN 	`json:"ResponsibleObserver"`          // OBX-16
	ObservationMethod            *ServiceCode `json:"ObservationMethod"`            // OBX-17
	EquipmentInstanceIdentifier  *EntityIdentifier `json:"EquipmentInstanceIdentifier"`  // OBX-18
	DateTimeOfTheAnalysis        string `json:"DateTimeOfTheAnalysis"`        // OBX-19
	ObservationSite              *ServiceCode `json:"ObservationSite"`              // OBX-20
	ObservationInstanceIdentifier *EntityIdentifier `json:"ObservationInstanceIdentifier"` // OBX-21
	MoodCode                     *ServiceCode `json:"MoodCode"`                     // OBX-22
	PerformingOrganizationName   *XON `json:"PerformingOrganizationName"`   // OBX-23
	PerformingOrganizationAddress *XAD `json:"PerformingOrganizationAddress"` // OBX-24
	PerformingOrganizationMedicalDirector *XCN `json:"PerformingOrganizationMedicalDirector"` // OBX-25
}

type NK1 struct {
	SetID                         string `json:"SetID"`                         // NK1-1
	Name                          *XPN `json:"Name"`                          // NK1-2
	Relationship                  *ServiceCode `json:"Relationship"`                  // NK1-3
	Address                       *XAD `json:"Address"`                       // NK1-4
	PhoneNumber                   *XTN `json:"PhoneNumber"`                   // NK1-5
	BusinessPhoneNumber           *XTN `json:"BusinessPhoneNumber"`           // NK1-6
	ContactRole                   *ServiceCode `json:"ContactRole"`                   // NK1-7
	StartDate                     string `json:"StartDate"`                     // NK1-8
	EndDate                       string `json:"EndDate"`                       // NK1-9
	NextOfKinAssociatedPartiesJobTitle string `json:"NextOfKinAssociatedPartiesJobTitle"` // NK1-10
	NextOfKinAssociatedPartiesJobCode *JCC `json:"NextOfKinAssociatedPartiesJobCode"` // NK1-11
	NextOfKinAssociatedPartiesEmployeeNumber *CX `json:"NextOfKinAssociatedPartiesEmployeeNumber"` // NK1-12
	OrganizationName              *XON `json:"OrganizationName"`              // NK1-13
	MaritalStatus                 *ServiceCode `json:"MaritalStatus"`                 // NK1-14
	AdministrativeSex             string `json:"AdministrativeSex"`             // NK1-15
	DateOfBirth                   string `json:"DateOfBirth"`                   // NK1-16
	LivingDependency              string `json:"LivingDependency"`              // NK1-17
	AmbulatoryStatus              string `json:"AmbulatoryStatus"`              // NK1-18
	Citizenship                   string `json:"Citizenship"`                   // NK1-19
	PrimaryLanguage               *ServiceCode `json:"PrimaryLanguage"`               // NK1-20
	LivingArrangement             string `json:"LivingArrangement"`             // NK1-21
	PublicityCode                 *ServiceCode `json:"PublicityCode"`                 // NK1-22
	ProtectionIndicator           string `json:"ProtectionIndicator"`           // NK1-23
	StudentIndicator              string `json:"StudentIndicator"`              // NK1-24
	Religion                      *ServiceCode `json:"Religion"`                      // NK1-25
	MothersMaidenName             *XPN `json:"MothersMaidenName"`             // NK1-26
	Nationality                   *ServiceCode `json:"Nationality"`                   // NK1-27
	EthnicGroup                   *ServiceCode `json:"EthnicGroup"`                   // NK1-28
	ContactReason                 *ServiceCode `json:"ContactReason"`                 // NK1-29
	ContactPersonsName            *XPN `json:"ContactPersonsName"`            // NK1-30
	ContactPersonsTelephoneNumber *XTN `json:"ContactPersonsTelephoneNumber"` // NK1-31
	ContactPersonsAddress         *XAD `json:"ContactPersonsAddress"`         // NK1-32
	NextOfKinAssociatedPartysIdentifiers *CX `json:"NextOfKinAssociatedPartysIdentifiers"` // NK1-33
	JobStatus                     string `json:"JobStatus"`                     // NK1-34
	Race                          *ServiceCode `json:"Race"`                          // NK1-35
	Handicap                      string `json:"Handicap"`                      // NK1-36
	ContactPersonSocialSecurityNumber string `json:"ContactPersonSocialSecurityNumber"` // NK1-37
	NextOfKinBirthPlace           string `json:"NextOfKinBirthPlace"`           // NK1-38
	VIPIndicator                  string `json:"VIPIndicator"`                  // NK1-39
}

type CX struct {
	IDNumber           string `json:"IDNumber"`           // CX.1
	CheckDigit         string `json:"CheckDigit"`         // CX.2
	CheckDigitScheme   string `json:"CheckDigitScheme"`   // CX.3
	AssigningAuthority string `json:"AssigningAuthority"` // CX.4
	IdentifierTypeCode string `json:"IdentifierTypeCode"` // CX.5
	AssigningFacility  string `json:"AssigningFacility"`  // CX.6
	EffectiveDate      string `json:"EffectiveDate"`      // CX.7
	ExpirationDate     string `json:"ExpirationDate"`     // CX.8
	AssigningJurisdiction string `json:"AssigningJurisdiction"` // CX.9
	AssigningAgencyOrDept string `json:"AssigningAgencyOrDept"` // CX.10
}

type JCC struct {
	JobCode			string		`json:"JobCode"`
	JobClass		string		`json:"JobClass"`
	JobDescription		string		`json:"JobDescription"`
}

type OBR struct {
	SetID			string		`json:"SetID"`			// OBR-1
	PlacerOrderNumber	*Placer		`json:"PlacerOrderNumber"`	// OBR-2
	FillerOrderNumber	*Filler		`json:"FillerOrderNumber"`	// OBR-3
	UniversalServiceID	*ServiceCode	`json:"UniversalService"`	// OBR-4
	Priority		string		`json:"Priority"`		// OBR-5
	RequestDate		*OBRDateTime	`json:"Request"`		// OBR-6
	ObservationDate		*OBRDateTime	`json:"Observation"`		// OBR-7
	ObservationEndDate	*OBRDateTime	`json:"ObservationEndDate"`	// OBR-8
	CollectionVolume	*OBRVolume	`json:"CollectionVolume"`	// OBR-9
	CollectorID		*OBRReceptCode	`json:"Collector"`		// OBR-10
	SpecimenAction		string		`json:"SpecimenAction"`		// OBR-11
	DangerCode		*ServiceCode	`json:"ServiceCode"`		// OBR-12
	ClinicalInfo		string		`json:"ClinicalInfo"`		// OBR-13
	SpecimenRecivedDate	*OBRDateTime	`json:"SpecimenRecivedDate"`	// OBR-14
	SpecimenSource		*OBRSource	`json:"SpecimenSource"`		// OBR-15
	OrderingProvider	*OBRReceptCode	`json:"OrderingProvider"`	// OBR-16
	CallbackPhone		string		`json:"CallbackPhone"`		// OBR-17
	PlacerField1		string		`json:"PlacerField1"`		// OBR-18
	PlacerField2		string		`json:"PlacerField2"`		// OBR-19
	FillerField1		string		`json:"FillerField1"`		// OBR-20
	FillerField2		string		`json:"FillerField2"`		// OBR-21
	ResultReportDate	*OBRDateTime	`json:"ResultReportTime"`	// OBR-22
	ChargeToPractice	*Charge		`json:"ChargeToPractice"`	// OBR-23
	DiagnosticService	string		`json:"DiagnosticService"`	// OBR-24
	ResultStatus		string		`json:"ResultStatus"`		// OBR-25
	ParentResult		*ParentResultCode		`json:"ParentResult"`		// OBR-26
	Quantity		*QuantityTiming	`json:"QuantityTiming"`		// OBR-27
	ResultCopyTo		*OBRReceptCode	`json:"ResultCopyTo"`		// OBR-28
	ParentNumber		*OBRParentNumber	`json:"ParentNumber"`	// OBR-29
	TransportationMode	string		`json:"TransportationMode"`	// OBR-30
	ReasonForStudy		*ServiceCode	`json:"ReasonForStudy"`		// OBR-31
	PrincipalInterpreter	*Technician	`json:"PrincipalInterpreter"`	// OBR-32
	AssistantInterpreter	*Technician	`json:"AssistantInterpreter"`	// OBR-33
	Technician		*Technician	`json:"Technician"`		// OBR-34
	Transcription		*Technician	`json:"Transcription"`		// OBR-35
	ScheduledDate		*OBRDateTime	`json:"ScheduledDate"`		// OBR-36
}

type Technician struct {
	Technician		string		`json:"Technician"`		// OBR-x.1
	StartDate		string		`json:"StartDate"`		// OBR-x.2
	EndDate			string		`json:"EndDate"`		// OBR-x.3
	Location		string		`json:"Location"`		// OBR-x.4
}

type OBRParentNumber struct {
	PlacerOrderNumber	string		`json:"PlacerOrderNumber"`	// OBR-29.1
	FillerOrderNumber	string		`json:"FillerOrderNumber"`	// OBR-29.2
}

type QuantityTiming struct {
	QuantityAmount		string		`json:"QuantityAmount"`		// OBR-27.1
	Interval		string		`json:"Interval"`		// OBR-27.2
	Duration		string		`json:"Duration"`		// OBR-27.3
	StartDate		string		`json:"StartDate"`		// OBR-27.4
	EndDate			string		`json:"EndDate"`		// OBR-27.5
	Priority		string		`json:"Priority"`		// OBR-27.6
	Condition		string		`json:"Condition"`		// OBR-27.7
	Text			string		`json:"Text"`			// OBR-27.8
	Conjunction		string		`json:"Conjunction"`		// OBR-27.9
	OrderSequence		string		`json:"OrderSequence"`		// OBR-27.10
}

type ParentResultCode struct {
	ObservationID		string		`json:"ObservationID"`		// OBR-26.1
	ParentResultSubID	string		`json:"ParentResultSubID"`	// OBR-26.2
	ObservationResult	string		`json:"ObservationResult"`	// OBR-26.3
}

type Charge struct {
	DollarAmount		string		`json:"DollarAmount"`		// OBR-23.1
	ChargeCode		string		`json:"ChargeCode"`		// OBR-23.2
}

type OBRSource struct {
	SpecimenSourceCode	string		`json:"SpecimenSourceCode"`	// OBR-15.1
	Additives		string		`json:"Additives"`		// OBR-15.2
	FreeText		string		`json:"FreeText"`		// OBR-15.3
	BodySite		string		`json:"BodySite"`		// OBR-15.4
	SiteModifier		string		`json:"SiteModifier"`		// OBR-15.5
}

type OBRReceptCode struct {
	IDNumber		string		`json:"IDNumber"`		// OBR-10.1
	FamilyName		string		`json:"FamilyName"`		// OBR-10.2
	GivenName		string		`json:"GivenName"`		// OBR-10.3
	MiddleName		string		`json:"MiddleName"`		// OBR-10.4
	Suffix			string		`json:"Suffix"`			// OBR-10.5
	Prefix			string		`json:"Prefix"`			// OBR-10.6
	Degree			string		`json:"Degree"`			// OBR-10.7
	Source			string		`json:"Source"`			// OBR-10.8
}

type OBRVolume struct {
	Quantity		string		`json:"Quantity"`		// OBR-9.1
	Units			string		`json:"Units"`			// OBR-9.2
}

type OBRDateTime struct {
	EventTime		string		`json:"EventTime"`		// OBR-x.1
	DegreeOfPrecision	string		`json:"DegreeOfPrecision"`	// OBR-x.2
}

// Generic Nested Section Used in Many Cases
type ServiceCode struct {
	Identifier		string		`json:"Identifier"`		// OBR-x.1
	Text			string		`json:"Text"`			// OBR-x.2
	CodingSystem		string		`json:"CodingSystem"`		// OBR-x.3
	AlternateIdentifier	string		`json:"AlternateIdentifier"`	// OBR-x.4
	AlternateText		string		`json:"AlternateText"`		// OBR-x.5
	AlternateCodingSystem	string		`json:"AlternateCodingSystem"`	// OBR-x.6
}

type Filler struct {
	UniqueFillerID		string		`json:"UniqueFillerID"`		// OBR-3.1
	FillerApplication	string		`json:"FillerApplication"`	// OBR-3.2
}

type Placer struct {
	UniquePlacerID		string		`json:"UniquePlacerID"`		// OBR-2.1
	PlacerApplication	string		`json:"PlacerApplication"`	// OBR-2.2
}

type HL7Message struct {
	MSH *MSH `json:"MSH"`
	EVN *EVN `json:"EVN"`
	PID *PID `json:"PID"`
	PV1 *PV1 `json:"PV1"`
	PV2 *PV2 `json:"PV2"`
	OBR *OBR `json:"OBR"`
}

func NewOBRSegment(p *random.Patient) *OBR {

	obr := &OBR{
		SetID:               "000002",
		PlacerOrderNumber:   &Placer{},
		FillerOrderNumber:   &Filler{},
		UniversalServiceID:  &ServiceCode{},
		Priority:            "",
		RequestDate:         &OBRDateTime{},
		ObservationDate:     &OBRDateTime{},
		ObservationEndDate:  &OBRDateTime{},
		CollectionVolume:    &OBRVolume{},
		CollectorID:         &OBRReceptCode{},
		SpecimenAction:      "",
		DangerCode:          &ServiceCode{},
		ClinicalInfo:        "",
		SpecimenRecivedDate: &OBRDateTime{},
		SpecimenSource:      &OBRSource{},
		OrderingProvider:    &OBRReceptCode{},
		CallbackPhone:       "",
		PlacerField1:        "",
		PlacerField2:        "",
		FillerField1:        "",
		FillerField2:        "",
		ResultReportDate:    &OBRDateTime{},
		ChargeToPractice:    &Charge{},
		DiagnosticService:   "",
		ResultStatus:        "",
		ParentResult:        &ParentResultCode{},
		Quantity:            &QuantityTiming{},
		ResultCopyTo:        &OBRReceptCode{},
		ParentNumber:        &OBRParentNumber{},
		TransportationMode:  "",
		ReasonForStudy:      &ServiceCode{},
		PrincipalInterpreter: &Technician{},
		AssistantInterpreter: &Technician{},
		Technician:           &Technician{},
		Transcription:        &Technician{},
		ScheduledDate:        &OBRDateTime{},
	}	

	return obr

}

func NewPV2Segment(p *random.Patient) *PV2 {

	pv2 := &PV2{
		PriorPendingLocation: &PatientLocation{},
		AccommodationCode: 	&ServiceCode{},
		AdmitReason: 		&ServiceCode{},
		TransferReason:		&ServiceCode{},
		ReferralSourceCode: 	&XCN{},
		ClinicOrganizationName:	&XON{},
		PatientChargeAdjustmentCode: &ServiceCode{},
	}

	return pv2
}

func NewPV1Segment(p *random.Patient) *PV1 {

	doctor := &AttendingDoctor{
		IDNumber:  "00002224",
		FirstName: "Greg",
		LastName:  "House",
	}

	location := &PatientLocation{
		Bed:      "105A",
		Room:     "225",
		Facility: "Test Facility",
	}

	pv1 := &PV1{
		AssignedPatientLocation: location,
		AttendingDoctor:         doctor,
		VisitNumber:             p.EncounterId,
		AdmitDateTime:           p.Hl7Info.HL7Arrival,
		DischargeDateTime:       p.Hl7Info.HL7Discharge,
	}

	return pv1

}

func NewPIDSegment(p *random.Patient) *PID {

	street := "123 " + p.PatientAddress.Street

	name := &XPN{
		LastName:      p.LastName,
		FirstName:     p.FirstName,
		MiddleInitial: "A",
	}

	address := &XAD{
		StreetAddress: street,
		City:          p.PatientAddress.RegionInfo.City,
		State:         p.PatientAddress.RegionInfo.State,
	}

	pid := &PID{
		SetID:                 "1",
		PatientID:             "123456789",
		PatientIdentifierList: p.MRN,
		PatientName:           name,
		DateOfBirth:           p.Hl7Info.HL7DOB,
		Sex:                   "M",
		PatientAddress:        address,
		HomePhone:             p.Phone,
	}

	return pid

}

func NewEVNSegment(p *random.Patient) *EVN {

	evn := &EVN{
		EventTypeCode:    "AO1", //placeholder
		RecordedDateTime: p.Hl7Info.HL7Event,
	}

	return evn
}

func NewMSHSegment(p *random.Patient) *MSH {

	messageType := &MessageType{
		MessageCode:  "ADT",
		TriggerEvent: "AO1",
	}

	msh := &MSH{
		Encode:               "^~\\&",
		SendingApplication:   "PLACEBO",
		SendingFacility:      "PLACEBO",
		ReceivingApplication: "DEMO",
		ReceivingFacility:    "DEMO",
		DateTimeOfMessage:    p.Hl7Info.HL7Event,
		Security:             "",
		MessageType:          messageType,
		MessageControlID:     "123456",
		ProcessingID:         "P",
		VersionID:            "2.3",
	}

	return msh
}

func NewHL7EventMessage(p *random.Patient, t string, e string) string {
	
	message := NewHL7Message(p)
	message.MSH.MessageType.MessageCode = t
	message.MSH.MessageType.TriggerEvent = e

	message.EVN.EventTypeCode = t

	product := MessageBuilder(message)

	return product

}



func NewHL7Message(p *random.Patient) *HL7Message {

	message := &HL7Message{
		MSH: NewMSHSegment(p),
		EVN: NewEVNSegment(p),
		PID: NewPIDSegment(p),
		PV1: NewPV1Segment(p),
		PV2: NewPV2Segment(p),
		OBR: NewOBRSegment(p),
	}

	return message

}

func (m *HL7Message) MessageToJson() string {

	message, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}

	return string(message)

}

func JsonToMessage(s string) *HL7Message {
	jsonMessage := []byte(s)

	var message *HL7Message
	err := json.Unmarshal(jsonMessage, &message)
	if err != nil {
		panic(err)
	}

	return message

}

func (msg *HL7Message) ReadFromFile(s string) *HL7Message {
	lines := LinesFromFile(s)

	for _, line := range lines {
		segs := strings.Split(line, "|")
		segType := SegmentFinder(segs[0])
		onlyMessageSegs := segs[1:]
		bind := message.MarshallHL7(onlyMessageSegs, segType)
		msg.SegmentAssigner(bind)
	}

	return msg
}

func ReadHL7(content string) string {

	messageStruct := &HL7Message{}
	message := messageStruct.ReadFromFile(content)

	messageFmt, _ := json.MarshalIndent(message, "", " ")

	jsonMessage := SanitizeLines(string(messageFmt))

	return jsonMessage

}

func SegmentFinder(s string) interface{} {

	switch s {
	case "MSH":
		return &MSH{}
	case "EVN":
		return &EVN{}
	case "PID":
		return &PID{}
	case "PV1":
		return &PV1{}
	case "PV2":
		return &PV2{}
	case "OBR":
		return &OBR{}
	default:
		return nil
	}

}

func (msg *HL7Message) SegmentAssigner(s interface{}) *HL7Message {

	switch v := s.(type) {
	case *MSH:
		msg.MSH = v
	case *EVN:
		msg.EVN = v
	case *PID:
		msg.PID = v
	case *PV1:
		msg.PV1 = v
	case *PV2:
		msg.PV2 = v
	case *OBR:
		msg.OBR = v
	default:
		return nil
	}

	return msg
}

func MessageBuilder(msg *HL7Message) string {
	segments := []string{
		message.CreateHL7(msg.MSH, "MSH"),
		message.CreateHL7(msg.EVN, "EVN"),
		message.CreateHL7(msg.PID, "PID"),
		message.CreateHL7(msg.PV1, "PV1"),
		message.CreateHL7(msg.PV2, "PV2"),
		message.CreateHL7(msg.OBR, "OBR"),
	}

	return strings.Join(segments, "\n")

}

func LinesFromFile(s string) []string {

	splitLines := strings.Split(s, "\n")

	return splitLines
}

// Remove 'null' segments for easier JSON display
func SanitizeLines(lines string) string {
	split := strings.Split(lines, "\n")

	var newLines []string

	for _, line := range split {
		if !strings.Contains(line, ": null") {
			newLines = append(newLines, line)
		}
	}

	join := strings.Join(newLines, "\n")

	return join

}
