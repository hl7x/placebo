package cmd

import (
	"errors"
	"os"
	"strconv"

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
			file.Create(patients)
			return nil
		}

		intParse, err := strconv.ParseInt(os.Args[3], 10, 0)
		if err != nil {
			return err
		}

		amount := int(intParse)

		if amount > 0 {
			patients := random.NewPatients(amount)
			file.Create(patients)
			return nil
		} else {

			return nil
		}
	default:
		return errors.New("Not a Valid Command!")
	}
	return nil
}
