package templates

func ConstructCSVFileHeaders() []byte {

	csvHeaders := []byte(`PatientMRN,PatientEncounterId,PatientFirstName,PatientLastName,PatientDOB,PatientGender,PatientAddress,PatientCity,PatientState,PatientPostCode,PatientCountry,PatientArrivalDate,PatientArrivalTime,DepartmentReferenceId,PatientPhonePrimary,PrimaryLanguage,VisitProvider,AppointmentBeginDate,AppointmentStatus,AltID`)

	return csvHeaders

}

func CSVPatientInfo() []byte {

	patientTemplate := []byte(`{{ range .Patients}}
{{.MRN}},{{.EncounterId}},{{.FirstName}},{{.LastName}},{{.DOB}},M,{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}},{{.PatientAddress.RegionInfo.City}},{{.PatientAddress.RegionInfo.State}},10119,USA,{{.ArrivalDate}} 06:00 PM,17:30 EDT,DEP1,{{.Phone}},English,JACK RYAN,{{.Appointment}} 03:00 PM,Scheduled,TESTALTID283860873785909398765{{ end  }}
`)
	return patientTemplate

}

// ADT^A01 is an event type for Admiting the Patient. This is the deafult message template.
func SimpleHl7Info() []byte {

	hl7Example := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|202310060800||ADT^A01|12345|P|2.3|
EVN|A01|202310060800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12345|
PV1|1|IP|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}|
`)

	return hl7Example

}
