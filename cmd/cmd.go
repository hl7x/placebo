package cmd

import (
	"fmt"

	"placebo/file"
	"placebo/pkg/random"
)

func PatientFile(f string) error {

	switch f {
	case "":
		return nil
	case "init":
		patient := random.NewPatient()
		file.Create(patient)
		return nil
	default:
		fmt.Println("Not A Valid Command")
		return nil
	}
	return nil
}
