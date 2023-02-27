
package cmd

import (
	"fmt"
	"os"
	"strconv"
	
	"placebo/file"
	"placebo/pkg/random"
)

func PatientFile(f string) error {
	
	switch f {
	case "":
		return nil
	case "init":
		
		command := os.Args[3:]

		if len(command) == 0 {
			patients := random.NewPatients(1)
			file.Create(patients)
			return nil
		}

		intParse, err := strconv.ParseInt(os.Args[3], 10, 0)
		if err != nil {
			fmt.Println("Integer Required")
			return nil
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
		fmt.Println("Not A Valid Command")
		return nil
	}
	return nil
}

