package file

import (
	"fmt"
	"os"
	"text/template"
	"time"
	"io/ioutil"

	"placebo/pkg/random"
	"placebo/pkg/templates"
)

var Tempdir = "/tmp/"
var IntFile = "ctmephl7.tmp"

func CreateCSV(patients random.Collection) error {

	csvFile := Tempdir + CSVFileName()

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

	hl7File := Tempdir + Hl7FileName()

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

// TODO: Panics used here should be fixed
func CreateInteractiveHl7(tempAndPatient string) string {

	hl7File := Tempdir + IntFile
	
	file, err := os.Create(hl7File)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(tempAndPatient)
	if err != nil {
		panic(err)
	}

	file.Sync()


	return hl7File

}
// TODO: Panics used here should be fixed
func ReadInteractiveHl7(path string) string {
	
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)

}

func SugarPillInteractive(s string) string {

	tempFile := "/tmp/sptemp.hl7"

	file, err := os.Create(tempFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}

	file.Sync()

	return tempFile

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
