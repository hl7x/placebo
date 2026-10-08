package cmd

import (
	"fmt"
	"strconv"

	"github.com/hl7x/placebo/file"
	"github.com/hl7x/placebo/pkg/event"
	"github.com/hl7x/placebo/pkg/random"
)

func File(args []string) error {

	if wantsHelp(args) {
		fmt.Println(commandHelp["file"])
		return nil
	}

	if len(args) == 0 {
		return missingSubcommand("file")
	}

	switch args[0] {
	case "csv":

		amount := 1

		if len(args) > 1 {
			parsed, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid patient count %q: expected a number", args[1])
			}

			amount = parsed
		}

		if amount < 1 {
			return fmt.Errorf("invalid patient count %q: expected a number greater than 0", args[1])
		}

		patients := random.NewPatients(amount)
		created, err := file.CreateCSV(patients)
		if err != nil {
			return err
		}

		fmt.Printf("File Created: %v\n", created)

		return nil

	case "hl7":

		if len(args) > 1 && args[1] == "types" {
			fmt.Println(event.Catalog())
			return nil
		}

		messageType, triggerType, err := event.Resolve(args[1:])
		if err != nil {
			return err
		}

		patient := random.NewPatient()
		created, err := file.CreateHl7(patient, messageType, triggerType)
		if err != nil {
			return err
		}

		fmt.Printf("File Created: %v\n", created)

		return nil

	default:
		return unknownSubcommand("file", args[0])
	}
}
