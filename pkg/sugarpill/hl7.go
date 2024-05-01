package sugarpill

import (
	"encoding/json"
	"strings"

	"placebo/pkg/random"
	"placebo/pkg/message"
)

type MSH struct {
	Encode			string `json:Encode`			// MSH-1
	SendingApplication    string    `json:"SendingApplication"`    // MSH-3
	SendingFacility       string    `json:"SendingFacility"`       // MSH-4
	ReceivingApplication  string    `json:"ReceivingApplication"`  // MSH-5
	ReceivingFacility     string    `json:"ReceivingFacility"`     // MSH-6
	DateTimeOfMessage     string    `json:"DateTimeOfMessage"`     // MSH-7
	Security              string    `json:"Security"`              // MSH-8
	MessageType           string    `json:"MessageType"`           // MSH-9
	MessageControlID      string    `json:"MessageControlID"`      // MSH-10
	ProcessingID          string    `json:"ProcessingID"`          // MSH-11
	VersionID             string    `json:"VersionID"`             // MSH-12
}

type EVN struct {
	EventTypeCode   string `json:"EventTypeCode"`				// EVN-1
	RecordedDateTime string `json:"RecordedDateTime"`		// EVN-2
	DateTimePlannedEvent 	string `json:"DateTimePlannedEvent"`	// EVN-3
	EventReasonCode		string	`json:"EventReasonCode"`	// EVN-4
	OperatorID		string	`json:"OperatorID"`		// EVN-5
	EventOccurred		string	`json:"EventOccurred"`		// EVN-6
	EventFacility		string	`json:"EventFacility"`		// EVN-7
}

type PatientAddress struct {
	StreetAddress string `json:"StreetAddress"` // PID-11.1
	OtherDesignation	string 	`json:"OtherDesignation"`	// PID.11.2
	City          string `json:"City"`          // PID-11.3
	State         string `json:"State"`         // PID-11.4
	ZipCode       string `json:"ZipCode"`       // PID-11.5
	Country       string `json:"Country"`       // PID-11.6
	AddressType	string	`json:"AddressType"`	// PID-11.7
	OtherGeographicDesignation	string	`json:"OtherGeographicDesignation"`	// PID-11.8
	CountyParishCode	string	`json:"CountryParishCode"`	//PID-11.9
	CensusTract		string	`json:"CensusTract"`		//PID-11.10
}

type PatientName struct {
	LastName               string         `json:"LastName"`               // PID-5.1
	FirstName              string         `json:"FirstName"`              // PID-5.2
	MiddleInitial          string         `json:"MiddleInitial"`          // PID-5.3
	Suffix			string		`json:"Suffix"`			// PID-5.4
	Prefix			string		`json:"Prefix"`			// PID-5.5
	Degree			string		`json:"Degree"`			// PID-5.6
	NameTypeCode		string		`json:"NameTypeCode"`		// PID-5.7
	NameRepresentationCode	string		`json:"NameRepresentationCode"`	// PID-5.8
	NameContext		string		`json:"NameContext"`		// PID-5.9
	NameValidityRange	string		`json:"NameValidityRange"`	// PID-5.10
	NameAssemblyOrder	string		`json:"NameAssemblyOrder"`	// PID-5.11
	EffectiveDate		string		`json:"EffectiveDate"`		// PID-5.12
	ExpirationDate		string		`json:"ExpirationDate"`		// PID-5.13
	ProfessionalSuffix	string		`json:"ProfessionalSuffix"`	// PID-5.14
}

type PID struct {
	SetID                  string         `json:"SetID"`                  // PID-1
	PatientID              string         `json:"PatientID"`              // PID-2
	PatientIdentifierList  int         `json:"PatientIdentifierList"`  // PID-3
	AlternatePatientID	string		`json:"AlternatePatientID"`	// PID-4
	PatientName	*PatientName		`json:"PatientName"`		// PID-5
	MotherMaidenName	string		`json:"MotherMaidenName"`	// PID-6
	DateOfBirth            string         `json:"DateOfBirth"`            // PID-7
	Sex                    string         `json:"Sex"`                    // PID-8
	PatientAlias		string		`json:"PatientAlias"`		// PID-9
	Race			string		`json:"Race"`			// PID-10
	PatientAddress         *PatientAddress `json:"PatientAddress"`		// PID-11
	CountyCode		string		`json:"CountyCode"`		// PID-12
	HomePhone		string		`json:"HomePhone"`		// PID-13
	BusinessPhone		string		`json:"BusinessPhone"`		// PID-14
	PrimaryLanguage		string		`json:"PrimaryLanguage"`	// PID-15
	MaritalStatus		string		`json:"MaritalStatus"`		// PID-16
	Religion		string		`json:"Religion"`		// PID-17
	PatientAccountNumber	string		`json:"PatientAccountNumber"`	// PID-18
	SSN                    string         `json:"SSN"`                    // PID-19
	DriversLicenseNumber	string		`json:"DriversLicenseNumber"`	// PID-20
	MotherIdentifier	string		`json:"MotherIdentifier"`	// PID-21
	EthnicGroup		string		`json:"EthnicGroup"`		// PID-22
	BirthPlace		string		`json:"BirthPlace"`		// PID-23
	MultipleBirthIndicator	string		`json:"MultipleBirthIndicator"`	// PID-24
	BirthOrder		string		`json:"BirthOrder"`		// PID-25
	Citizenship		string		`json:"Citizenship"`		// PID-26
	VeteranStatus		string		`json:"VeteranStatus"`		// PID-27
	Nationality		string		`json:"Nationality"`		// PID-28
	PatientDeathDateAndTime	string		`json:"PatientDeathDateAndTime"`// PID-29
	PatientDeathIndicator	string		`json:"PatientDeathIndicator"`	// PID-30
}

type AssignedPatientLocation struct {
	Facility     string `json:"Facility"`     // PV1-3.4
	PointOfCare  string `json:"PointOfCare"`  // PV1-3.1
}

type AttendingDoctor struct {
	IDNumber           string `json:"IDNumber"`           // PV1-7.1
	LastName           string `json:"LastName"`           // PV1-7.2
	FirstName          string `json:"FirstName"`          // PV1-7.3
	AssigningAuthority string `json:"AssigningAuthority"` // PV1-7.4
}

type PV1 struct {
	SetID                  string                `json:"SetID"`                  // PV1-1
	PatientClass           string                `json:"PatientClass"`           // PV1-2
	AssignedPatientLocation AssignedPatientLocation `json:"AssignedPatientLocation"`
	AdmissionType          string                `json:"AdmissionType"`          // PV1-4
	PreadmitNumber         string                `json:"PreadmitNumber"`         // PV1-5
	PriorPatientLocation   string                `json:"PriorPatientLocation"`   // PV1-6
	AttendingDoctor        AttendingDoctor       `json:"AttendingDoctor"`
	HospitalService        string                `json:"HospitalService"`        // PV1-10
	VisitNumber            int                `json:"VisitNumber"`            // PV1-19
	ArrivalDate            string             `json:"ArrivalDate"`
	DischargeDate          string             `json:"DischargeDate"`
}

type HL7Message struct {
	MSH *MSH `json:"MSH"`
	EVN *EVN `json:"EVN"`
	PID *PID `json:"PID"`
	PV1 *PV1 `json:"PV1"`
}

func NewPV1Segment(p *random.Patient) *PV1 {
	
	pv1 := &PV1{
		VisitNumber:  p.EncounterId,
		ArrivalDate: p.Hl7Info.HL7Arrival,
		DischargeDate:  p.Hl7Info.HL7Discharge,
	}

	return pv1

}

func NewPIDSegment(p *random.Patient) *PID {

	street := "123 "+p.PatientAddress.Street

	name := &PatientName{
		LastName: p.LastName, 
		FirstName: p.FirstName,
		MiddleInitial: "A",
	}

	address := &PatientAddress{
		StreetAddress: street,
		City: p.PatientAddress.RegionInfo.City,
		State: p.PatientAddress.RegionInfo.State,
	}

	pid := &PID{
		SetID: "1",
		PatientID: "123456789",
		PatientIdentifierList: p.MRN,
		PatientName: name,
		DateOfBirth: p.Hl7Info.HL7DOB,
		Sex: "M",
		PatientAddress: address,
		HomePhone:  p.Phone,

	}

	return pid

}

func NewEVNSegment(p *random.Patient) *EVN {
	
	evn := &EVN{
		EventTypeCode: "AO1", //placeholder
		RecordedDateTime: p.Hl7Info.HL7Event,
	}

	return evn
}

func NewMSHSegment(p *random.Patient) *MSH {
	
	msh := &MSH{
		Encode: "--",
		SendingApplication: "PLACEBO",
		SendingFacility: "PLACEBO",
		ReceivingApplication: "DEMO",
		ReceivingFacility: "DEMO",
		DateTimeOfMessage: p.Hl7Info.HL7Event,
		Security: "",
		MessageType: "ADT^A01",
		MessageControlID: "123456",
		ProcessingID: "P",
		VersionID: "2.3",
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
		message.HL7(msg.MSH, "MSH"),
        	message.HL7(msg.EVN, "EVN"),
		message.HL7(msg.PID, "PID"),
		message.HL7(msg.PV1, "PV1"),
	}

	return strings.Join(segments, "\n")

}
