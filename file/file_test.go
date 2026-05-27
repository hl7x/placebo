package file

import (
	"fmt"
	"log"
	"os"
	"placebo/pkg/event"
	"placebo/pkg/random"
	"placebo/pkg/templates"
	"strings"
	"testing"
	"time"
)

func TestCreateCSV(t *testing.T) {
	// Create a temporary directory for the test
	testDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(testDir)

	Tempdir = testDir + "/"

	examplePatient1 := &random.Patient{
		FirstName:      "Bill",
		LastName:       "Test",
		MRN:            "123",
		VisitId:        123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{RegionInfo: &random.Region{}},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	examplePatient2 := &random.Patient{
		FirstName:      "Jill",
		LastName:       "Test",
		MRN:            "123",
		VisitId:        123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{RegionInfo: &random.Region{}},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	mockPatients := random.Collection{
		Patients: []*random.Patient{examplePatient1, examplePatient2},
	}

	file, err := CreateCSV(mockPatients)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	files, err := os.ReadDir(Tempdir)
	if err != nil {
		t.Fatal("Failed to read temp dir:", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 file in temp dir, got %d", len(files))
	}

	csvFilePath := Tempdir + files[0].Name()

	if file != csvFilePath {
		t.Fatalf("Expected %v, got %v", csvFilePath, file)
	}

	csvContentBytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		t.Fatal("Failed to read CSV file:", err)
	}

	csvContent := string(csvContentBytes)

	expectedHeaders := "FirstName,LastName,MRN,VisitId,Phone,DOB,Street,StructureNumber,State,City,PostalCode,ArrivalDate,DischargeDate,Appointment,HL7Arrival,HL7Discharge,HL7Event,HL7DOB,HL7Appointment"
	expectedExamplePatient1Data := fmt.Sprintf("%v,%v,%v", examplePatient1.FirstName, examplePatient1.LastName, examplePatient1.MRN)
	expectedExamplePatient2Data := fmt.Sprintf("%v,%v,%v", examplePatient2.FirstName, examplePatient2.LastName, examplePatient2.MRN)

	if !strings.Contains(csvContent, expectedHeaders) {
		t.Errorf("%v content does not contain the expected headers: %s", csvContent, expectedHeaders)
	}

	if !strings.Contains(csvContent, expectedExamplePatient1Data) {
		t.Errorf("%v content does not contain the expected patient data: %s", csvContent, expectedExamplePatient1Data)
	}

	if !strings.Contains(csvContent, expectedExamplePatient2Data) {
		t.Errorf("%v content does not contain the expected patient data: %s", csvContent, expectedExamplePatient2Data)
	}

}

func TestCreateHl7(t *testing.T) {
	// Create a temporary directory for the test
	testDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(testDir)

	Tempdir = testDir + "/"

	examplePatient := &random.Patient{
		FirstName:      "Bill",
		LastName:       "Test",
		MRN:            "123",
		VisitId:        123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{RegionInfo: &random.Region{}},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
		Hl7Info:        &random.Hl7Dates{},
	}

	file, err := CreateHl7(examplePatient)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	files, err := os.ReadDir(Tempdir)
	if err != nil {
		t.Fatal("Failed to read temp dir:", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 file in temp dir, got %d", len(files))
	}

	hl7FilePath := Tempdir + files[0].Name()

	if file != hl7FilePath {
		t.Fatalf("Expected File: %v, got %v", hl7FilePath, file)
	}

	hl7ContentBytes, err := os.ReadFile(hl7FilePath)
	if err != nil {
		t.Fatal("Failed to read HL7 file:", err)
	}

	hl7Content := string(hl7ContentBytes)

	expectedExamplePatientData := fmt.Sprintf("%v", examplePatient.MRN)

	if !strings.Contains(hl7Content, expectedExamplePatientData) {
		t.Errorf("HL7 content does not contain the expected patient data: %s", expectedExamplePatientData)
	}

}

func TestCSVFileName(t *testing.T) {
	currentTime := time.Now()
	date := fmt.Sprintf("%v%v%v", currentTime.Year(), int(currentTime.Month()), currentTime.Day())

	var tests = []struct {
		description string
		expected    string
	}{
		{"Default Case", date + "_patient_import.csv"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := CSVFileName()
			if got != tc.expected {
				t.Fatalf("FileName()=%v expected %v", got, tc.expected)
			}
		})
	}
}

func TestHl7FileName(t *testing.T) {
	currentTime := time.Now()
	date := fmt.Sprintf("%v%v%v", currentTime.Year(), int(currentTime.Month()), currentTime.Day())

	var tests = []struct {
		description string
		expected    string
	}{
		{"Default Case", date + "hl7__patient_import.txt"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := Hl7FileName()
			if got != tc.expected {
				t.Fatalf("FileName()=%v expected %v", got, tc.expected)
			}
		})
	}
}

func TestCreateInteractiveHl7(t *testing.T) {
	// Create a temporary directory for the test
	testDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(testDir)

	Tempdir = testDir + "/"

	patient := &random.Patient{
		FirstName:      "Bill",
		LastName:       "Test",
		MRN:            "123",
		VisitId:        123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{RegionInfo: &random.Region{}},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	template := templates.SimpleHl7Info()

	patientAndTemplate := event.TemplateMapper(patient, template)

	var tests = []struct {
		description string
		input       string
		expected    string
	}{
		{"Should Include Return File Path When Patient And Template Are Constructed", patientAndTemplate, Tempdir},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got := CreateInteractiveHl7(tc.input)

			if !strings.Contains(got, tc.expected) {
				t.Fatalf("CreateIneractiveHl7(%v)=%v expected to include %v", tc.input, got, tc.expected)
			}
		})
	}

}

func TestReadInteractiveHl7(t *testing.T) {
	// Create a temporary directory for the test
	testDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(testDir)

	Tempdir = testDir + "/"

	f, err := os.CreateTemp(Tempdir, "test")
	if err != nil {
		log.Fatal(err)
	}

	writeTest := []byte("Testing")

	_, err = f.Write(writeTest)
	if err != nil {
		log.Fatal(err)
	}

	var tests = []struct {
		description string
		input       string
		expected    string
	}{
		{"Should Return the Read File String", f.Name(), "Testing"},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			got, _ := ReadFile(tc.input)

			if got != tc.expected {
				t.Fatalf("ReadInteractionHl7(%v)=%v expected %v", tc.input, got, tc.expected)
			}
		})
	}

}
