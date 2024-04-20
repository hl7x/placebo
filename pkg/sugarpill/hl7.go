package sugarpill

import (
	"time"

	"placebo/pkg/random"
)

type MSH struct {
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
	HL7Type   string `json:"HL7Type"`
	TimeStamp string `json:"TimeStamp"`
}

type PatientAddress struct {
	StreetAddress string `json:"StreetAddress"` // PID-11.1
	City          string `json:"City"`          // PID-11.3
	State         string `json:"State"`         // PID-11.4
	ZipCode       string `json:"ZipCode"`       // PID-11.5
	Country       string `json:"Country"`       // PID-11.6
}

type PID struct {
	SetID                  string         `json:"SetID"`                  // PID-1
	PatientID              string         `json:"PatientID"`              // PID-2
	PatientIdentifierList  string         `json:"PatientIdentifierList"`  // PID-3
	LastName               string         `json:"LastName"`               // PID-5.1
	FirstName              string         `json:"FirstName"`              // PID-5.2
	MiddleInitial          string         `json:"MiddleInitial"`          // PID-5.3
	DateOfBirth            string         `json:"DateOfBirth"`            // PID-7
	Sex                    string         `json:"Sex"`                    // PID-8
	PatientAddress         PatientAddress `json:"PatientAddress"`
	PhoneNumberHome        string         `json:"PhoneNumberHome"`        // PID-13.1
	PhoneNumberBusiness    string         `json:"PhoneNumberBusiness"`    // PID-13.2
	PrimaryLanguage        string         `json:"PrimaryLanguage"`        // PID-15
	MaritalStatus          string         `json:"MaritalStatus"`          // PID-16
	SSN                    string         `json:"SSN"`                    // PID-19
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
	VisitNumber            string                `json:"VisitNumber"`            // PV1-19
	ArrivalDate            time.Time             `json:"ArrivalDate"`
	DischargeDate          time.Time             `json:"DischargeDate"`
}

type HL7Message struct {
	MSH MSH `json:"MSH"`
	EVN EVN `json:"EVN"`
	PID PID `json:"PID"`
	PV1 PV1 `json:"PV1"`
}

func (m *PV1) Pv1Line(p *random.Patient) *PV1 {

}

func (m *PID) PidLine(p *random.Patient) *PID {

}

func (m *EVN) EvnLine(p *random.Patient) *EVN {

}

func (m *MSH) MshLine(p *random.Patient) *MSH {

}

func (m *HL7Message) MessageLineFusesr(p *random.Patient) *Hl7Message {

	msh = m.MSH
	msh.MshLine(p)

	evn = m.EVN
	evn.EvnLine(p)

	pid = m.PID
	pid.PidLine(p)

	pv1 = m.PV1
	pv1.Pv1Line(p)

	return m

}

func PatientData(p *random.Patient) *Hl7Message {
	
	m := Hl7Message{}
	
	message := m.MessageLineFuser(p)

	return message
}
