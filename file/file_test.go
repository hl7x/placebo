package file

import (
	"fmt"
	"log"
	"os"
	"placebo/pkg/random"
	"strings"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	dir = tempDir + "/"

	examplePatient1 := &random.Patient{
		FirstName:      "Bill",
		LastName:       "Test",
		MRN:            123,
		EncounterId:    123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	examplePatient2 := &random.Patient{
		FirstName:      "Jill",
		LastName:       "Test",
		MRN:            123,
		EncounterId:    123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	mockPatients := random.Collection{
		Patients: []*random.Patient{examplePatient1, examplePatient2},
	}

	err = Create(mockPatients)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal("Failed to read temp dir:", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 file in temp dir, got %d", len(files))
	}

	csvFilePath := dir + files[0].Name()
	csvContentBytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		t.Fatal("Failed to read CSV file:", err)
	}

	csvContent := string(csvContentBytes)

	expectedHeaders := "PatientMRN,PatientEncounterId,PatientFirstName,PatientLastName,PatientDOB,PatientGender,PatientAddress,PatientCity,PatientState,PatientPostCode,PatientCountry,PatientArrivalDate,PatientArrivalTime,DepartmentReferenceId,PatientPhonePrimary,PrimaryLanguage,VisitProvider,AppointmentBeginDate,AppointmentStatus,AltID"
	expectedExample1PatientData := fmt.Sprintf("%v,%v,%v", examplePatient1.MRN, examplePatient1.EncounterId, examplePatient1.FirstName)
	expectedExample2PatientData := fmt.Sprintf("%v,%v,%v", examplePatient2.MRN, examplePatient2.EncounterId, examplePatient2.FirstName)

	if !strings.Contains(csvContent, expectedHeaders) {
		t.Errorf("CSV content does not contain the expected headers: %s", expectedHeaders)
	}

	if !strings.Contains(csvContent, expectedExample1PatientData) {
		t.Errorf("CSV content does not contain the expected patient data: %s", expectedExample1PatientData)
	}

	if !strings.Contains(csvContent, expectedExample2PatientData) {
		t.Errorf("CSV content does not contain the expected patient data: %s", expectedExample2PatientData)
	}

}

func TestFileName(t *testing.T) {
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
			got := FileName()
			if got != tc.expected {
				t.Fatalf("FileName()=%v expected %v", got, tc.expected)
			}
		})
	}
}
