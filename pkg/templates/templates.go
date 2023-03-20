package templates

func DefaultFileHeaders() []byte {

	csvHeaders := []byte(`PatientMRN,PatientEncounterId,PatientFirstName,PatientLastName,PatientDOB,PatientGender,PatientAddress,PatientCity,PatientState,PatientPostCode,PatientCountry,PatientArrivalDate,PatientArrivalTime,DepartmentReferenceId,PatientPhonePrimary,PrimaryLanguage,VisitProvider,AppointmentBeginDate,AppointmentStatus,AltID`)

	return csvHeaders

}

func DefaultPatientInfo() []byte {

	patientTemplate := []byte(`{{ range .Patients}}
{{.MRN}},{{.EncounterId}},{{.FirstName}},{{.LastName}},{{.DOB}},M,{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}},{{.PatientAddress.RegionInfo.City}},{{.PatientAddress.RegionInfo.State}},10119,USA,11/15/2022 06:00 PM,17:30 EDT,DEP1,{{.Phone}},English,JACK RYAN,11/15/2022 06:00 PM,Scheduled,TESTALTID283860873785909398765{{ end  }}
`)
	return patientTemplate

}
