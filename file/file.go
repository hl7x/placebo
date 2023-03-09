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

// stub for a quick function to create the csv file
func Create(patients random.Collection) {

	csvFile := dir + FileName()

	f, err := os.Create(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	csvHeaders := templates.ConstructFileHeaders()
	patientTemplateActions := templates.PatientInfo()

	fileText := append(csvHeaders, patientTemplateActions...)

	// write template with placeholders
	f.Write(fileText)

	t, err := template.ParseFiles(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(file, patients)
	if err != nil {
		fmt.Println(err)
	}

	file.Close()

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
