package templates

import (
// "fmt"
)

func ConstructFileHeaders() []byte {

	csvHeaders := []byte(`PatientMRN,PatientEncounterId,PatientFirstName,PatientLastName,PatientDOB,PatientGender,PatientAddress1,PatientAddress2,PatientCity,PatientState,PatientPostCode,PatientCountry,PatientArrivalDate,PatientArrivalTime,DepartmentReferenceId,PatientPhonePrimary,PrimaryLanguage,VisitProvider,AppointmentBeginDate,AppointmentStatus,AltID
`)

	return csvHeaders

}

func PatientInfo() []byte {

	patientTemplate := []byte(`{{.MRN}},{{.EncounterId}},{{.FirstName}},{{.LastName}},{{.DOB}},M,1 PENNSYLVANIA PLAZA,PENN1,NEW YORK,NY,10119,USA,11/15/2022 06:00 PM,17:30 EDT,DEP1,19737719600,English,JACK RYAN,11/15/2022 06:00 PM,Scheduled,TESTALTID283860873785909398765
`)
	return patientTemplate

}
