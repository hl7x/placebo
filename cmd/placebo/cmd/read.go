package cmd

import (
	"errors"
	"fmt"

	"github.com/hl7x/placebo/file"
	"github.com/hl7x/placebo/pkg/sugarpill"
)

func ReadHl7Message(args []string) error {

	if wantsHelp(args) {
		fmt.Println(commandHelp["read"])
		return nil
	}

	if len(args) == 0 {
		return missingSubcommand("read")
	}

	if args[0] != "sugarpill" {
		return unknownSubcommand("read", args[0])
	}

	if len(args) < 2 {
		return errors.New("'placebo read sugarpill' needs a file\n\n\tplacebo read sugarpill <path/to/hl7_file.txt>")
	}

	hl7Content, err := file.ReadFile(args[1])
	if err != nil {
		return err
	}

	message := sugarpill.ReadHL7(hl7Content)

	fmt.Println(message)

	return nil
}
