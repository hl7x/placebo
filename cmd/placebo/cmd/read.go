package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hl7x/placebo/file"
	"github.com/hl7x/placebo/pkg/sugarpill"
)

func ReadHl7Message(f string) error {

	switch f {
	case "":
		return nil
	case "sugarpill":

		filePath := os.Args[3:]

		if len(filePath) == 0 {
			return errors.New("Expected File\n Usage: placebo --read sugarpill <path/to/hl7_file.txt>")
		}

		if len(filePath) == 0 {
			return nil
		}

		hl7Content, err := file.ReadFile(filePath[0])
		if err != nil {
			return err
		}

		message := sugarpill.ReadHL7(hl7Content)

		fmt.Println(message)

		return nil
	default:
		return errors.New("Required Subcommand \n Usage: placebo --read sugarpill <path/to/hl7_file.txt>")
	}

	return nil
}
