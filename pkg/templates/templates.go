package templates

/*
DEPRECATED: templates make it easier to format the output.
But this requires too much continuous maintenance overhead.
This will stay in place for now, but should not continue to be maintained and added to.
*/

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

	hl7Example := []byte(`MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A01|12345|P|2.3|
EVN|A01|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346||{{.Phone}}|
PV1|1|IP|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}|||||||||||||||||||||||||{{.Hl7Info.HL7Arrival}}|||
`)

	return hl7Example

}

// ADT^03 is an event type for Discharging the Patient
func DischargeHl7Info() []byte {

	dischargeHl7 := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A03|12345|P|2.3|
EVN|A03|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346||{.Phone}}|
PV1|1|DC|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}||||||||||||||||||||||||||{{.Hl7Info.HL7Discharge}}||

`)
	return dischargeHl7

}

func PreadmitHl7Info() []byte {

	preadmit := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||ADT^A05|12345|P|2.3|
EVN|A05|{{.Hl7Info.HL7Event}}0800||
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346||{{.Phone}}|
PV1|1|PR|PUNIT^RM 123^P^PUNIT|ER|||15551234567^HOUSE^GREG|1255568135^BETA^ALPHA|30104384^SMITH^JOHN|HIM||||Phys/Clinic|||123456789^BETA^ALPHA|TESTX|{{.EncounterId}}|||||||||||||||||||||||||{{.Hl7Info.HL7Arrival}}|||
`)

	return preadmit

}

// SIU^12 is an appointment scheduling structure for Pre-Admitting the Patient on a scheduled appointment
func ScheduledPatientInfo() []byte {

	scheduleSiu := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP|LAB|{{.Hl7Info.HL7Event}}0800||SIU^S12|12345|P|2.3|
SCH|12345|12345|OFFICE^123^PUNIT|{{.Hl7Info.HL7Appointment}}|0900|||60|M^Minutes|Regular Appointment||ID123|||Scheduled|
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12345||{.Phone}}|
RGS|1|A|
`)

	return scheduleSiu

}

// Referral Message
func ReferralHl7() []byte {

	referral := []byte(`
MSH|^~\&|SENDAPP|PLACEBO|RECVAPP||20240124151605757||REF^I12|0000|P|2.3.1|||2.3.1
RF1|NEW|ASAP|NVgastro|||{{.EncounterId}}|{{.Hl7Info.HL7Appointment}}|20250124|202401231139|Referral
AUT||||20240124|
PRD|||||||CENTER FOR TESTING 00
PID|1|56789|{{.MRN}}||{{.LastName}}^{{.FirstName}}||{{.Hl7Info.HL7DOB}}|M|||{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}^^{{.PatientAddress.RegionInfo.City}}^{{.PatientAddress.RegionInfo.State}}^12346||{{.Phone}}|
`)

	return referral

}

/*
func HelperHl7() []byte {

	helper := []byte(`
	{
  "MSH": {
    "SendingApplication": "YourApp", // MSH-3
    "SendingFacility": "YourFacility", // MSH-4
    "ReceivingApplication": "TheirApp", // MSH-5
    "ReceivingFacility": "TheirFacility", // MSH-6
    "DateTimeOfMessage": "{{.Hl7Info.HL7Event}}0800", // MSH-7
    "Security": "", // MSH-8
    "MessageType": "ADT^A01", // MSH-9
    "MessageControlID": "123456", // MSH-10
    "ProcessingID": "P", // MSH-11
    "VersionID": "2.3" // MSH-12
  },
  "EVN": {
	"HL7Type": "A01",
	"TimeStamp": "{{.Hl7Info.HL7Event}}0800"
  },
  "PID": {
    "SetID": "1", // PID-1
    "PatientID": "123456", // PID-2
    "PatientIdentifierList": "{{.MRN}}^^^YourSystem^MR", // PID-3
    "LastName": "{{.LastName}}", // PID-5.1
    "FirstName": "{{.FirstName}}", // PID-5.2
    "MiddleInitial": "A", // PID-5.3
    "DateOfBirth": "{{.Hl7Info.HL7DOB}}", // PID-7
    "Sex": "M", // PID-8
    "PatientAddress": {
      "StreetAddress": "{{.PatientAddress.StructureNumber}} {{.PatientAddress.Street}}", // PID-11.1
      "City": "{{.PatientAddress.RegionInfo.City}}", // PID-11.3
      "State": "{{.PatientAddress.RegionInfo.State}}", // PID-11.4
      "ZipCode": "12345", // PID-11.5
      "Country": "USA" // PID-11.6
    },
    "PhoneNumberHome": "{{.Phone}}", // PID-13.1
    "PhoneNumberBusiness": "555-5678", // PID-13.2
    "PrimaryLanguage": "English", // PID-15
    "MaritalStatus": "Single", // PID-16
    "SSN": "123-45-6789" // PID-19
  },
  "PV1": {
    "SetID": "1", // PV1-1
    "PatientClass": "O", // PV1-2
    "AssignedPatientLocation": {
      "Facility": "PUNIT", // PV1-3.4
      "PointOfCare": "ER" // PV1-3.1
    },
    "AdmissionType": "ER", // PV1-4
    "PreadmitNumber": "789012", // PV1-5
    "PriorPatientLocation": "ICU", // PV1-6
    "AttendingDoctor": {
      "IDNumber": "54321", // PV1-7.1
      "LastName": "Smith", // PV1-7.2
      "FirstName": "Jane", // PV1-7.3
      "AssigningAuthority": "Medical Staff" // PV1-7.4
    },
    "HospitalService": "General", // PV1-10
    "VisitNumber": "{{.EncounterId}}", // PV1-19
    "ArrivalDate": "", //
    "DischargeDate": "" //
  }
}
`)

	return helper

}
*/
