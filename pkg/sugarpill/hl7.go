package sugarpill

import (
	"encoding/json"
	"strings"

	"placebo/pkg/message"
	"placebo/pkg/random"
)

type MSH struct {
	Encode               string `json:Encode`                 // MSH-1
	SendingApplication   string `json:"SendingApplication"`   // MSH-3
	SendingFacility      string `json:"SendingFacility"`      // MSH-4
	ReceivingApplication string `json:"ReceivingApplication"` // MSH-5
	ReceivingFacility    string `json:"ReceivingFacility"`    // MSH-6
	DateTimeOfMessage    string `json:"DateTimeOfMessage"`    // MSH-7
	Security             string `json:"Security"`             // MSH-8
	MessageType          string `json:"MessageType"`          // MSH-9
	MessageControlID     string `json:"MessageControlID"`     // MSH-10
	ProcessingID         string `json:"ProcessingID"`         // MSH-11
	VersionID            string `json:"VersionID"`            // MSH-12
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

type PatientAddress struct {
	StreetAddress              string `json:"StreetAddress"`              // PID-11.1
	OtherDesignation           string `json:"OtherDesignation"`           // PID.11.2
	City                       string `json:"City"`                       // PID-11.3
	State                      string `json:"State"`                      // PID-11.4
	ZipCode                    string `json:"ZipCode"`                    // PID-11.5
	Country                    string `json:"Country"`                    // PID-11.6
	AddressType                string `json:"AddressType"`                // PID-11.7
	OtherGeographicDesignation string `json:"OtherGeographicDesignation"` // PID-11.8
	CountyParishCode           string `json:"CountryParishCode"`          //PID-11.9
	CensusTract                string `json:"CensusTract"`                //PID-11.10
}

type PatientName struct {
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

type PID struct {
	SetID                   string          `json:"SetID"`                   // PID-1
	PatientID               string          `json:"PatientID"`               // PID-2
	PatientIdentifierList   int             `json:"PatientIdentifierList"`   // PID-3
	AlternatePatientID      string          `json:"AlternatePatientID"`      // PID-4
	PatientName             *PatientName    `json:"PatientName"`             // PID-5
	MotherMaidenName        string          `json:"MotherMaidenName"`        // PID-6
	DateOfBirth             string          `json:"DateOfBirth"`             // PID-7
	Sex                     string          `json:"Sex"`                     // PID-8
	PatientAlias            string          `json:"PatientAlias"`            // PID-9
	Race                    string          `json:"Race"`                    // PID-10
	PatientAddress          *PatientAddress `json:"PatientAddress"`          // PID-11
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

type HL7Message struct {
	MSH *MSH `json:"MSH"`
	EVN *EVN `json:"EVN"`
	PID *PID `json:"PID"`
	PV1 *PV1 `json:"PV1"`
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

	name := &PatientName{
		LastName:      p.LastName,
		FirstName:     p.FirstName,
		MiddleInitial: "A",
	}

	address := &PatientAddress{
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

	msh := &MSH{
		Encode:               "--",
		SendingApplication:   "PLACEBO",
		SendingFacility:      "PLACEBO",
		ReceivingApplication: "DEMO",
		ReceivingFacility:    "DEMO",
		DateTimeOfMessage:    p.Hl7Info.HL7Event,
		Security:             "",
		MessageType:          "ADT^A01",
		MessageControlID:     "123456",
		ProcessingID:         "P",
		VersionID:            "2.3",
	}

	return msh
}

func NewHL7Message(p *random.Patient) *HL7Message {

	message := &HL7Message{
		MSH: NewMSHSegment(p),
		EVN: NewEVNSegment(p),
		PID: NewPIDSegment(p),
		PV1: NewPV1Segment(p),
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

func MessageBuilder(msg *HL7Message) string {
	segments := []string{
		message.CreateHL7(msg.MSH, "MSH"),
		message.CreateHL7(msg.EVN, "EVN"),
		message.CreateHL7(msg.PID, "PID"),
		message.CreateHL7(msg.PV1, "PV1"),
	}

	return strings.Join(segments, "\n")

}
