package file

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"placebo/pkg/random"
	"placebo/pkg/templates"
)

var dir = "/tmp/"

func CreateCSV(patients random.Collection) error {

	csvFile := dir + CSVFileName()

	csvHeaders := templates.ConstructCSVFileHeaders()
	patientTemplateActions := templates.CSVPatientInfo()

	fileText := append(csvHeaders, patientTemplateActions...)

	t, err := template.New("csv").Parse(string(fileText))
	if err != nil {
		return err
	}

	file, err := os.Create(csvFile)
	if err != nil {
		return err
	}

	err = t.Execute(file, patients)
	if err != nil {
		return err
	}

	return nil

}

func CreateHl7(patient *random.Patient) error {

	hl7File := dir + Hl7FileName()

	fileText := templates.SimpleHl7Info()
	
	t, err := template.New("txt").Parse(string(fileText))
	if err != nil {
		return err
	}

	file, err := os.Create(hl7File)
	if err != nil {
		return err
	}

	err = t.Execute(file, patient)
	if err != nil {
		return err
	}

	return nil

}

func CSVFileName() string {

	currentTime := time.Now()

	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v%v%v", year, int(month), day)

	fileName := date + "_patient_import.csv"

	return fileName
}

func Hl7FileName() string {

	currentTime := time.Now()

	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v%v%v", year, int(month), day)

	fileName := date + "hl7__patient_import.txt"

	return fileName
}
