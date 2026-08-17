package sugarpill

import "github.com/hl7x/placebo/pkg/random"

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

func NewPV1Segment(p *random.Patient) *PV1 {

	doctor := &AttendingDoctor{
		IDNumber:  p.Provider.ID,
		FirstName: p.Provider.FirstName,
		LastName:  p.Provider.LastName,
	}

	location := &PatientLocation{
		Bed:      p.Location.Bed,
		Room:     p.Location.Room,
		Facility: p.Location.LocationFacility,
	}

	pv1 := &PV1{
		AssignedPatientLocation: location,
		AttendingDoctor:         doctor,
		VisitNumber:             p.VisitId,
		AdmitDateTime:           p.ArrivalDate.HL7(),
		DischargeDateTime:       p.DischargeDate.HL7(),
	}

	return pv1
}
