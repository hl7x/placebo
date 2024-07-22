package cmd

import (
	"errors"
	"os"
	"strconv"
	"fmt"

	"placebo/file"
	"placebo/pkg/random"
)

func File(f string) error {

	switch f {
	case "":
		return nil
	case "csv":

		command := os.Args[3:]

		if len(command) == 0 {
			patients := random.NewPatients(1)
			file, err := file.CreateCSV(patients)
			if err != nil {
				return err
			}

			fmt.Printf("File Created: %v\n", file)
			return nil
		}

		intParse, err := strconv.ParseInt(os.Args[3], 10, 0)
		if err != nil {
			return err
		}

		amount := int(intParse)

		if amount > 0 {
			patients := random.NewPatients(amount)
			file, err := file.CreateCSV(patients)
			if err != nil {
				return err
			}

			fmt.Printf("File Created: %v\n", file)
			return nil
		} else {

			return nil
		}
	case "hl7":
		patient := random.NewPatient()
		file, err := file.CreateHl7(patient)
		if err != nil {
			return err
		}

		fmt.Printf("File Created: %v\n", file)

		return nil
	default:
		return errors.New("Command Not Found.")
	}
}
