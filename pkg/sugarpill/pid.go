package sugarpill

import "placebo/pkg/random"

type PID struct {
	SetID                   string `json:"SetID"`                   // PID-1
	PatientID               string `json:"PatientID"`               // PID-2
	PatientIdentifierList   string `json:"PatientIdentifierList"`   // PID-3
	AlternatePatientID      string `json:"AlternatePatientID"`      // PID-4
	PatientName             *XPN   `json:"PatientName"`             // PID-5
	MotherMaidenName        string `json:"MotherMaidenName"`        // PID-6
	DateOfBirth             string `json:"DateOfBirth"`             // PID-7
	Sex                     string `json:"Sex"`                     // PID-8
	PatientAlias            string `json:"PatientAlias"`            // PID-9
	Race                    string `json:"Race"`                    // PID-10
	PatientAddress          *XAD   `json:"PatientAddress"`          // PID-11
	CountyCode              string `json:"CountyCode"`              // PID-12
	HomePhone               string `json:"HomePhone"`               // PID-13
	BusinessPhone           string `json:"BusinessPhone"`           // PID-14
	PrimaryLanguage         string `json:"PrimaryLanguage"`         // PID-15
	MaritalStatus           string `json:"MaritalStatus"`           // PID-16
	Religion                string `json:"Religion"`                // PID-17
	PatientAccountNumber    string `json:"PatientAccountNumber"`    // PID-18
	SSN                     string `json:"SSN"`                     // PID-19
	DriversLicenseNumber    string `json:"DriversLicenseNumber"`    // PID-20
	MotherIdentifier        string `json:"MotherIdentifier"`        // PID-21
	EthnicGroup             string `json:"EthnicGroup"`             // PID-22
	BirthPlace              string `json:"BirthPlace"`              // PID-23
	MultipleBirthIndicator  string `json:"MultipleBirthIndicator"`  // PID-24
	BirthOrder              string `json:"BirthOrder"`              // PID-25
	Citizenship             string `json:"Citizenship"`             // PID-26
	VeteranStatus           string `json:"VeteranStatus"`           // PID-27
	Nationality             string `json:"Nationality"`             // PID-28
	PatientDeathDateAndTime string `json:"PatientDeathDateAndTime"` // PID-29
	PatientDeathIndicator   string `json:"PatientDeathIndicator"`   // PID-30
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
		Sex:                   p.Sex,
		PatientAddress:        address,
		HomePhone:             p.Phone,
	}

	return pid
}
