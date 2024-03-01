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


/* QUICK HL7 TEMPLATES*/


// ADT^A01 is an event type for Admiting the Patient. This is the deafult message template.
func SimpleHl7Info() []byte {

	hl7Example := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A01|12345|P|2.3|
EVN|A01|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346|
PV1|1|IP|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}|||||||||||||||||||||||||{{.Hl7Info.HL7Arrival}}|||
`)

	return hl7Example

}

// ADT^03 is an event type for Discharging the Patient
func DischargeHl7Info() []byte {

	dischargeHl7 := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A03|12345|P|2.3|
EVN|A03|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12345|
PV1|1|DC|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}||||||||||||||||||||||||||{{.Hl7Info.HL7Discharge}}||

`)
	return dischargeHl7

}

func PreadmitHl7Info() []byte {

	preadmit := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A05|12345|P|2.3|
EVN|A05|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346|
PV1|1|PR|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}|||||||||||||||||||||||||{{.Hl7Info.HL7Arrival}}|||
`)

	return preadmit

}

//SIU^12 is an appointment scheduling structure for Pre-Admitting the Patient on a scheduled appointment
func ScheduledPatientInfo() []byte {

	scheduleSiu := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||SIU^S12|12345|P|2.3|
SCH|12345|12345|OFFICE^123^PUNIT|{{.Hl7Info.HL7Appointment}}|0900|||60|M^Minutes|Regular Appointment||ID123|||Scheduled|
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12345|
RGS|1|A|
`)

	return scheduleSiu

}

//Referral Message
func ReferralHl7() []byte {

	referral := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|CipherHealth||20240124151605757||REF^I12|0000|P|2.3.1|||2.3.1
RF1|NEW|ASAP|NVgastro|||394445731811|202401241516|20250124|202401231139|Referral
AUT||||20240124|
PRD|||||||CENTER FOR TESTING 00
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346|
`)

	return referral

}
