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

func Create(patients random.Collection) error {

	csvFile := dir + FileName()

	csvHeaders := templates.ConstructFileHeaders()
	patientTemplateActions := templates.PatientInfo()

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

func FileName() string {

	currentTime := time.Now()

	month := currentTime.Month()
	day := currentTime.Day()
	year := currentTime.Year()

	date := fmt.Sprintf("%v%v%v", year, int(month), day)

	fileName := date + "_patient_import.csv"

	return fileName
}
