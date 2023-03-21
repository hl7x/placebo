package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"placebo/pkg/random"
	"strings"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := ioutil.TempDir("", "test_create")
	if err != nil {
		t.Fatal("Failed to create temp dir:", err)
	}
	defer os.RemoveAll(tempDir)

	dir = tempDir

	examplePatient := &random.Patient{
		FirstName:      "Bill",
		LastName:       "Test",
		MRN:            000,
		EncounterId:    123,
		Phone:          "0000000",
		DOB:            "00/00/0000",
		PatientAddress: &random.Address{},
		ArrivalDate:    "00/00/0000",
		DischargeDate:  "00/00/0000",
	}

	mockPatients := random.Collection{
		Patients: []*random.Patient{examplePatient},
	}

	Create(mockPatients)

	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		t.Fatal("Failed to read temp dir:", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 file in temp dir, got %d", len(files))
	}

	csvFilePath := tempDir + "/" + files[0].Name()
	csvContentBytes, err := ioutil.ReadFile(csvFilePath)
	if err != nil {
		t.Fatal("Failed to read CSV file:", err)
	}

	csvContent := string(csvContentBytes)

	expectedHeaders := "PatientMRN,PatientEncounterId,PatientFirstName,PatientLastName,PatientDOB,PatientGender,PatientAddress,PatientCity,PatientState,PatientPostCode,PatientCountry,PatientArrivalDate,PatientArrivalTime,DepartmentReferenceId,PatientPhonePrimary,PrimaryLanguage,VisitProvider,AppointmentBeginDate,AppointmentStatus,AltID"
	expectedPatientData := ""

	if !strings.Contains(csvContent, expectedHeaders) {
		t.Errorf("CSV content does not contain the expected headers: %s", expectedHeaders)
	}

	if !strings.Contains(csvContent, expectedPatientData) {
		t.Errorf("CSV content does not contain the expected patient data: %s", expectedPatientData)
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
